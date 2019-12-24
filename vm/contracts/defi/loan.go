package defi

import (
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/vm/contracts/common"
	"github.com/vitelabs/go-vite/vm/util"
	"github.com/vitelabs/go-vite/vm_db"
	"math/big"
)

func CheckLoanParam(param *ParamNewLoan) error {
	if param.Token != ledger.ViteTokenId || param.DayRate <= MinDayRate || param.DayRate >= MaxDayRate ||
		param.ShareAmount.Cmp(minShareAmount) < 0 || param.Shares <= 0 ||
		param.SubscribeDays < MinSubDays || param.SubscribeDays > MaxSubDays || param.ExpireDays <= 0 {
		return InvalidInputParamErr
	} else {
		return nil
	}
}

func NewLoan(address types.Address, db vm_db.VmDb, param *ParamNewLoan, interest *big.Int) *Loan {
	loan := &Loan{}
	loan.Id = NewLoanSerialNo(db)
	loan.Address = address.Bytes()
	loan.Token = param.Token.Bytes()
	loan.ShareAmount = param.ShareAmount.Bytes()
	loan.Shares = param.Shares
	loan.Interest = interest.Bytes()
	loan.DayRate = param.DayRate
	loan.SubscribeDays = param.SubscribeDays
	loan.ExpireDays = param.ExpireDays
	loan.Status = LoanOpen
	loan.Created = GetDeFiTimestamp(db)
	return loan
}

func OnLoanInvest(db vm_db.VmDb, loan *Loan, amount *big.Int) {
	loan.Invested = common.AddBigInt(loan.Invested, amount.Bytes())
	SaveLoan(db, loan)
}

func OnLoanCancelInvest(db vm_db.VmDb, loan *Loan, amount []byte) {
	if common.CmpForBigInt(loan.Invested, amount) < 0 {
		panic(ExceedFundAvailableErr)
	} else {
		loan.Invested = common.SubBigIntAbs(loan.Invested, amount)
		SaveLoan(db, loan)
	}
}

func DoRefundLoan(db vm_db.VmDb, loan *Loan) {
	address, _ := types.BytesToAddress(loan.Address)
	switch loan.Status {
	case LoanFailed:
		OnAccLoanFailed(db, address, loan.Interest)
		AddBaseAccountEvent(db, loan.Address, BaseLoanInterestRelease, 0, loan.Id, loan.Interest)
	case LoanExpiredRefunded:
		amount := CalculateAmount(loan.Shares, loan.ShareAmount)
		OnAccLoanExpired(db, address, amount)
		AddLoanAccountEvent(db, loan.Address, LoanAccExpiredRefund, 0, loan.Id, amount.Bytes())
	}
	if loan.SubscribedShares > 0 {
		refundLoanSubscriptions(db, loan)
	}
	DeleteLoan(db, loan)
	AddLoanUpdateEvent(db, loan)
}

func DoCancelExpiredLoanInvests(db vm_db.VmDb, loan *Loan) (blocks []*ledger.AccountBlock, err error) {
	dexIdsToCancel := make([]byte, 0, 80)
	err = traverseLoanInvests(db, loan.Id, func(investId uint64) error {
		var blks []*ledger.AccountBlock
		if invest, ok := GetInvest(db, investId); ok && invest.Status == InvestSuccess {
			CancellingInvest(db, invest)
			switch invest.BizType {
			case InvestForMining, InvestForSVIP:
				dexIdsToCancel = append(dexIdsToCancel, common.Uint64ToBytes(investId)...)
			case InvestForQuota:
				blks, err = DoCancelQuotaInvest(invest.InvestHash)
			case InvestForSBP:
				blks, err = DoRevokeSBP(db, invest.InvestHash)
			}
		}
		if err != nil {
			return err
		} else {
			blocks = append(blocks, blks...)
			return nil
		}
	})
	if err != nil {
		return
	}
	if len(dexIdsToCancel) != 0 {
		if dexBlk, err1 := DoCancelDexInvest(dexIdsToCancel); err1 != nil {
			return nil, err1
		} else {
			blocks = append(blocks, dexBlk...)
		}
	}
	return
}

func refundLoanSubscriptions(db vm_db.VmDb, loan *Loan) {
	traverseLoanSubscriptions(db, loan, func(sub *Subscription) error {
		amount := CalculateAmount(sub.Shares, sub.ShareAmount)
		if loan.Status == LoanExpiredRefunded {
			OnAccRefundSuccessSubscription(db, sub.Address, amount)
			AddBaseAccountEvent(db, sub.Address, BaseSubscribeExpiredRefund, 0, loan.Id, amount.Bytes())
		} else if loan.Status == LoanFailed {
			OnAccRefundFailedSubscription(db, sub.Address, amount)
			AddBaseAccountEvent(db, sub.Address, BaseSubscribeFailedRelease, 0, loan.Id, amount.Bytes())
		}
		sub.Status = loan.Status
		DeleteSubscription(db, sub)
		AddSubscriptionUpdateEvent(db, sub)
		return nil
	})
}

func expireLoanSubscriptions(db vm_db.VmDb, loan *Loan) {
	traverseLoanSubscriptions(db, loan, func(sub *Subscription) error {
		sub.Status = loan.Status
		AddSubscriptionUpdateEvent(db, sub)
		return nil
	})
}

func NewSubscription(address types.Address, db vm_db.VmDb, param *ParamSubscribe, loan *Loan) *Subscription {
	sub := &Subscription{}
	sub.LoanId = param.LoanId
	sub.Address = address.Bytes()
	sub.Token = loan.Token
	sub.Shares = param.Shares
	sub.ShareAmount = loan.ShareAmount
	sub.Status = LoanOpen
	sub.Created = GetDeFiTimestamp(db)
	return sub
}

func DoSubscribe(db vm_db.VmDb, gs util.GlobalStatus, loan *Loan, shares int32, deFiDayHeight uint64) (err error) {
	loan.SubscribedShares = loan.SubscribedShares + shares
	loan.Updated = GetDeFiTimestamp(db)
	if loan.Shares == loan.SubscribedShares {
		loan.Status = LoanSuccess
		loan.ExpireHeight = GetExpireHeight(gs, loan.ExpireDays, deFiDayHeight)
		loan.StartHeight = gs.SnapshotBlock().Height
		loan.StartTime = loan.Updated
		OnAccLoanSuccess(db, loan.Address, loan)
		AddLoanAccountEvent(db, loan.Address, LoanAccNewSuccessLoan, 0, loan.Id, CalculateAmount(loan.Shares, loan.ShareAmount).Bytes())
	}
	SaveLoan(db, loan)
	AddLoanUpdateEvent(db, loan)
	if loan.Status == LoanSuccess {
		err = traverseLoanSubscriptions(db, loan, func(sub *Subscription) (err1 error) {
			sub.Status = LoanSuccess
			sub.Updated = loan.Updated
			amount := CalculateAmount(sub.Shares, sub.ShareAmount)
			SaveSubscription(db, sub)
			AddSubscriptionUpdateEvent(db, sub)
			if _, err1 = OnAccSubscribeSuccess(db, sub.Address, amount); err1 != nil {
				return err1
			}
			AddBaseAccountEvent(db, sub.Address, BaseSubscribeSuccessReduce, 0, loan.Id, amount.Bytes())
			return
		})
	}
	return
}

func GetLoanSubscriptions(db vm_db.VmDb, loanId uint64) (subs []*Subscription, err error) {
	loan := &Loan{}
	loan.Id = loanId
	err = traverseLoanSubscriptions(db, loan, func(sub *Subscription) error {
		subs = append(subs, sub)
		return nil
	})
	return
}

func traverseLoanSubscriptions(db vm_db.VmDb, loan *Loan, traverseFunc func(sub *Subscription) error ) (err error) {
	iterator, err := db.NewStorageIterator(append(subscriptionKeyPrefix, common.Uint64ToBytes(loan.Id)...))
	if err != nil {
		panic(err)
	}
	defer iterator.Release()
	for {
		if !iterator.Next() {
			if iterator.Error() != nil {
				panic(iterator.Error())
			}
			break
		}
		data := iterator.Value()
		sub := &Subscription{}
		if err = sub.DeSerialize(data); err != nil {
			panic(err)
		}
		if err = traverseFunc(sub); err != nil {
			return
		}
	}
	return
}

func CalculateInterest(shares int32, shareAmount *big.Int, dayRate, days int32) *big.Int {
	totalRate := dayRate * days
	totalAmount := CalculateAmount1(shares, shareAmount)
	return new(big.Int).SetBytes(common.CalculateAmountForRate(totalAmount.Bytes(), totalRate, LoanRateCardinalNum))
}