package dex

import (
	"fmt"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/interfaces"
	"github.com/vitelabs/go-vite/vm_db"
	"math/big"
)

func DoDivideFees(db vm_db.VmDb, periodId uint64) error {
	var (
		feeSumsMap map[uint64]*FeeSumByPeriod
		vxSumFunds *VxFunds
		err        error
		ok         bool
	)

	//allow divide history fees that not divided yet
	if feeSumsMap = GetNotDividedFeeSumsByPeriodId(db, periodId); len(feeSumsMap) == 0 { // no fee to divide
		return nil
	}
	if vxSumFunds, ok = GetVxSumFunds(db); !ok {
		return nil
	}
	foundVxSumFunds, vxSumAmtBytes, needUpdateVxSum, _ := MatchVxFundsByPeriod(vxSumFunds, periodId, false)
	//fmt.Printf("foundVxSumFunds %v, vxSumAmtBytes %s, needUpdateVxSum %v with periodId %d\n", foundVxSumFunds, new(big.Int).SetBytes(vxSumAmtBytes).String(), needUpdateVxSum, periodId)
	if !foundVxSumFunds { // not found vxSumFunds
		return nil
	}
	if needUpdateVxSum {
		SaveVxSumFunds(db, vxSumFunds)
	}
	vxSumAmt := new(big.Int).SetBytes(vxSumAmtBytes)
	if vxSumAmt.Sign() <= 0 {
		return nil
	}
	// sum fees from multi period not divided
	feeSumMap := make(map[types.TokenTypeId]*big.Int)
	for pId, fee := range feeSumsMap {
		for _, feeAccount := range fee.FeesForDividend {
			if tokenId, err := types.BytesToTokenTypeId(feeAccount.Token); err != nil {
				return err
			} else {
				toDividendAmt, _ := splitDividendPool(feeAccount)
				if amt, ok := feeSumMap[tokenId]; !ok {
					feeSumMap[tokenId] = toDividendAmt
				} else {
					feeSumMap[tokenId] = amt.Add(amt, toDividendAmt)
				}
			}
		}
		MarkFeeSumAsFeeDivided(db, fee, pId)
	}

	var (
		userVxFundsKey, userVxFundsBytes []byte
	)

	iterator, err := db.NewStorageIterator(VxFundKeyPrefix)
	if err != nil {
		return err
	}
	defer iterator.Release()

	feeSumLeavedMap := make(map[types.TokenTypeId]*big.Int)
	dividedVxAmtMap := make(map[types.TokenTypeId]*big.Int)
	for {
		if len(feeSumMap) == 0 {
			break
		}
		if ok = iterator.Next(); ok {
			userVxFundsKey = iterator.Key()
			userVxFundsBytes = iterator.Value()
			if len(userVxFundsBytes) == 0 {
				continue
			}
		} else {
			break
		}

		addressBytes := userVxFundsKey[len(VxFundKeyPrefix):]
		address := types.Address{}
		if err = address.SetBytes(addressBytes); err != nil {
			return err
		}
		userVxFunds := &VxFunds{}
		if err = userVxFunds.DeSerialize(userVxFundsBytes); err != nil {
			return err
		}

		var userFeeDividend = make(map[types.TokenTypeId]*big.Int)
		foundVxFunds, userVxAmtBytes, needUpdateVxFunds, needDeleteVxFunds := MatchVxFundsByPeriod(userVxFunds, periodId, true)
		if !foundVxFunds {
			continue
		}
		if needDeleteVxFunds {
			DeleteVxFunds(db, address.Bytes())
		} else if needUpdateVxFunds {
			SaveVxFunds(db, address.Bytes(), userVxFunds)
		}
		userVxAmount := new(big.Int).SetBytes(userVxAmtBytes)
		//fmt.Printf("address %s, userVxAmount %s, needDeleteVxFunds %v\n", string(address.Bytes()), userVxAmount.String(), needDeleteVxFunds)
		if !IsValidVxAmountForDividend(userVxAmount) { //skip vxAmount not valid for dividend
			continue
		}

		var finished bool
		for tokenId, feeSumAmount := range feeSumMap {
			if _, ok = feeSumLeavedMap[tokenId]; !ok {
				feeSumLeavedMap[tokenId] = new(big.Int).Set(feeSumAmount)
				dividedVxAmtMap[tokenId] = big.NewInt(0)
			}
			//fmt.Printf("tokenId %s, address %s, vxSumAmt %s, userVxAmount %s, dividedVxAmt %s, toDivideFeeAmt %s, toDivideLeaveAmt %s\n", tokenId.String(), address.String(), vxSumAmt.String(), userVxAmount.String(), dividedVxAmtMap[tokenId], toDivideFeeAmt.String(), toDivideLeaveAmt.String())
			userFeeDividend[tokenId], finished = DivideByProportion(vxSumAmt, userVxAmount, dividedVxAmtMap[tokenId], feeSumAmount, feeSumLeavedMap[tokenId])
			if finished {
				delete(feeSumMap, tokenId)
			}
			AddFeeDividendEvent(db, address, tokenId, userVxAmount, userFeeDividend[tokenId])
		}
		if err = BatchSaveUserFund(db, address, userFeeDividend); err != nil {
			return err
		}
	}
	return DoDivideBrokerFees(db, feeSumsMap)
}

func DoDivideBrokerFees(db vm_db.VmDb, periodIdToFeeSum map[uint64]*FeeSumByPeriod) error {
	var (
		iterators                          = make([]interfaces.StorageIterator, 0, len(periodIdToFeeSum))
		ok                                 bool
		brokerFeeSumKey, brokerFeeSumBytes []byte
	)

	defer func() {
		for _, it := range iterators {
			it.Release()
		}
	}()

	for pId, _ := range periodIdToFeeSum {
		iterator, err := db.NewStorageIterator(append(brokerFeeSumKeyPrefix, Uint64ToBytes(pId)...))
		if err != nil {
			return err
		}
		iterators = append(iterators, iterator)
		if ok = iterator.Next(); ok {
			brokerFeeSumKey = iterator.Key() //3+8+21
			brokerFeeSumBytes = iterator.Value()
			if len(brokerFeeSumBytes) == 0 {
				continue
			}
			if len(brokerFeeSumKey) != 32 {
				panic(fmt.Errorf("invalid broker fee type"))
			}
			DeleteBrokerFeeSumByKey(db, brokerFeeSumKey)
		} else {
			break
		}
		brokerFeeSum := &BrokerFeeSumByPeriod{}
		if err = brokerFeeSum.DeSerialize(brokerFeeSumBytes); err != nil {
			panic(err)
		}
		addr, err := types.BytesToAddress(brokerFeeSumKey[11:])
		if err != nil {
			panic(err)
		}
		userFund := make(map[types.TokenTypeId]*big.Int)
		for _, feeAcc := range brokerFeeSum.BrokerFees {
			tokenId, err := types.BytesToTokenTypeId(feeAcc.Token)
			if err != nil {
				panic(err)
			}
			for _, mkFee := range feeAcc.MarketFees {
				if fd, ok1 := userFund[tokenId]; ok1 {
					userFund[tokenId] = new(big.Int).Add(fd, new(big.Int).SetBytes(mkFee.Amount))
				} else {
					userFund[tokenId] = new(big.Int).SetBytes(mkFee.Amount)
				}
				AddBrokerFeeDividendEvent(db, addr, mkFee)
			}
		}
		BatchSaveUserFund(db, addr, userFund)
	}
	return nil
}

func DivideByProportion(totalReferAmt, partReferAmt, dividedReferAmt, toDivideTotalAmt, toDivideLeaveAmt *big.Int) (proportionAmt *big.Int, finished bool) {
	dividedReferAmt.Add(dividedReferAmt, partReferAmt)
	proportion := new(big.Float).SetPrec(bigFloatPrec).Quo(new(big.Float).SetPrec(bigFloatPrec).SetInt(partReferAmt), new(big.Float).SetPrec(bigFloatPrec).SetInt(totalReferAmt))
	proportionAmt = RoundAmount(new(big.Float).SetPrec(bigFloatPrec).Mul(new(big.Float).SetPrec(bigFloatPrec).SetInt(toDivideTotalAmt), proportion))
	toDivideLeaveNewAmt := new(big.Int).Sub(toDivideLeaveAmt, proportionAmt)
	if toDivideLeaveNewAmt.Sign() <= 0 || dividedReferAmt.Cmp(totalReferAmt) >= 0 {
		proportionAmt.Set(toDivideLeaveAmt)
		finished = true
	} else {
		toDivideLeaveAmt.Set(toDivideLeaveNewAmt)
	}
	return proportionAmt, finished
}
