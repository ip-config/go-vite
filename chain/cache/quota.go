package chain_cache

import (
	"container/list"
	"fmt"
	"github.com/pkg/errors"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/ledger"
)

type item struct {
	BlockCount uint64
	Quota      uint64
}
type quotaList struct {
	chain Chain

	backElement map[types.Address]*item

	used                 map[types.Address]*item
	usedStart            *list.Element
	usedAccumulateHeight uint64

	list          *list.List
	listMaxLength int

	status byte
}

func newQuotaList(chain Chain) *quotaList {
	ql := &quotaList{
		chain: chain,
		used:  make(map[types.Address]*item),

		backElement: make(map[types.Address]*item),

		list:                 list.New(),
		listMaxLength:        600,
		usedAccumulateHeight: 75,
	}

	return ql
}

func (ql *quotaList) init() error {
	if err := ql.build(); err != nil {
		return err
	}

	ql.status = 1
	return nil
}

func (ql *quotaList) GetSnapshotQuotaUsed(addr *types.Address) (uint64, uint64) {
	used := ql.used[*addr]
	if used == nil {
		return 0, 0
	}
	quota := used.Quota
	blockCount := used.BlockCount
	latestUsed := ql.backElement[*addr]
	if latestUsed != nil {
		return quota - latestUsed.Quota, blockCount - latestUsed.BlockCount
	}

	return quota, blockCount
}

func (ql *quotaList) GetQuotaUsed(addr *types.Address) (uint64, uint64) {
	used := ql.used[*addr]
	if used == nil {
		return 0, 0
	}
	return used.Quota, used.BlockCount
}

func (ql *quotaList) Add(addr *types.Address, quota uint64) {
	backItem := ql.backElement[*addr]
	if backItem == nil {
		backItem = &item{}
		ql.backElement[*addr] = backItem
	}
	backItem.BlockCount += 1
	backItem.Quota += quota

	usedItem := ql.used[*addr]
	if usedItem == nil {
		usedItem = &item{}
		ql.used[*addr] = usedItem
	}
	usedItem.BlockCount += 1
	usedItem.Quota += quota
}

func (ql *quotaList) Sub(addr *types.Address, quota uint64) {
	ql.subBackElement(addr, 1, quota)
	ql.subUsed(addr, 1, quota)
}

func (ql *quotaList) NewNext() {
	if ql.status < 1 {
		return
	}
	ql.backElement = make(map[types.Address]*item)
	ql.list.PushBack(ql.backElement)

	if uint64(ql.list.Len()) <= ql.usedAccumulateHeight {
		return
	}

	quotaUsedStart := ql.usedStart.Value.(map[types.Address]*item)
	for addr, usedStartItem := range quotaUsedStart {
		if usedStartItem == nil {
			continue
		}
		ql.subUsed(&addr, usedStartItem.BlockCount, usedStartItem.Quota)
	}
	ql.usedStart = ql.usedStart.Next()
}

func (ql *quotaList) Rollback(n int) error {
	if n >= ql.listMaxLength {
		ql.list.Init()
	} else {
		// TODO
		for i := 0; i < n && ql.list.Len() > 0; i++ {
			ql.list.Remove(ql.list.Back())
		}
	}

	return ql.build()
}

func (ql *quotaList) build() (returnError error) {
	defer func() {
		if returnError != nil {
			return
		}
		ql.backElement = ql.list.Back().Value.(map[types.Address]*item)

		ql.resetUsedStart()

		ql.calculateUsed()
	}()

	listLength := uint64(ql.list.Len())

	if listLength >= ql.usedAccumulateHeight {
		return nil
	}

	latestSb, err := ql.chain.QueryLatestSnapshotBlock()
	if err != nil {
		return err
	}

	latestSbHeight := latestSb.Height

	if latestSbHeight <= listLength {
		return nil
	}

	endSbHeight := latestSbHeight + 1 - listLength
	startSbHeight := uint64(1)

	lackListLen := uint64(ql.listMaxLength) - listLength
	if endSbHeight > lackListLen {
		startSbHeight = endSbHeight - lackListLen
	}

	var snapshotSegments []*ledger.SnapshotChunk

	if listLength <= 0 {
		snapshotSegments, err = ql.chain.GetSubLedgerAfterHeight(startSbHeight)
		if err != nil {
			return err
		}

		if snapshotSegments == nil {
			return errors.New(fmt.Sprintf("ql.chain.GetSubLedgerAfterHeight, snapshotSegments is nil, startSbHeight is %d", startSbHeight))
		}

		for _, seg := range snapshotSegments[1:] {

			newItem := make(map[types.Address]*item)
			for _, block := range seg.AccountBlocks {
				if _, ok := newItem[block.AccountAddress]; !ok {
					newItem[block.AccountAddress] = &item{
						Quota:      block.Quota,
						BlockCount: 1,
					}
				} else {
					newItem[block.AccountAddress].Quota += block.Quota
					newItem[block.AccountAddress].BlockCount += 1
				}

			}
			ql.list.PushBack(newItem)
		}

		if snapshotSegments[len(snapshotSegments)-1].SnapshotBlock != nil {
			ql.list.PushBack(make(map[types.Address]*item))
		}

	} else {
		snapshotSegments, err = ql.chain.GetSubLedger(startSbHeight, endSbHeight)
		if err != nil {
			return err
		}

		if snapshotSegments == nil {
			return errors.New(fmt.Sprintf("ql.chain.GetSubLedger, snapshotSegments is nil, startSbHeight is %d, endSbHeight is %d",
				startSbHeight, endSbHeight))
		}

		segLength := len(snapshotSegments)
		for i := segLength - 1; i > 0; i-- {
			seg := snapshotSegments[i]
			newItem := make(map[types.Address]*item)

			for _, block := range seg.AccountBlocks {
				if _, ok := newItem[block.AccountAddress]; !ok {
					newItem[block.AccountAddress] = &item{
						Quota:      block.Quota,
						BlockCount: 1,
					}
				} else {
					newItem[block.AccountAddress].Quota += block.Quota
					newItem[block.AccountAddress].BlockCount += 1
				}
			}
			ql.list.PushFront(newItem)
		}

	}

	return nil
}

func (ql *quotaList) subBackElement(addr *types.Address, blockCount, quota uint64) {
	backItem := ql.backElement[*addr]
	if backItem == nil {
		return
	}
	backItem.BlockCount -= blockCount
	if backItem.BlockCount <= 0 {
		delete(ql.backElement, *addr)
		return
	}
	backItem.Quota -= quota

}

func (ql *quotaList) subUsed(addr *types.Address, blockCount, quota uint64) {
	usedItem := ql.used[*addr]
	if usedItem == nil {
		return
	}
	usedItem.BlockCount -= blockCount
	if usedItem.BlockCount <= 0 {
		delete(ql.used, *addr)
		return
	}
	usedItem.Quota -= quota
}

func (ql *quotaList) calculateUsed() {
	used := make(map[types.Address]*item)

	pointer := ql.usedStart
	for pointer != nil {
		tmpUsed := pointer.Value.(map[types.Address]*item)
		for addr, tmpItem := range tmpUsed {
			if used[addr] == nil {
				used[addr] = &item{}
			}

			used[addr].BlockCount += tmpItem.BlockCount
			used[addr].Quota += tmpItem.Quota
		}

		pointer = pointer.Next()
	}
	ql.used = used
}

func (ql *quotaList) resetUsedStart() {
	ql.usedStart = ql.list.Back()
	for i := uint64(1); i < ql.usedAccumulateHeight; i++ {
		prev := ql.usedStart.Prev()
		if prev == nil {
			break
		}
		ql.usedStart = prev

	}
}