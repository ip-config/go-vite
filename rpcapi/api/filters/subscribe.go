package filters

import (
	"context"
	"errors"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/log15"
	"github.com/vitelabs/go-vite/rpc"
	"github.com/vitelabs/go-vite/rpcapi/api"
	"github.com/vitelabs/go-vite/vite"
	"sync"
	"time"
)

var (
	deadline = 5 * time.Minute // consider a filter inactive if it has not been polled for within deadline
)

type filter struct {
	typ              FilterType
	deadline         *time.Timer
	param            api.FilterParam
	s                *RpcSubscription
	blocks           []*AccountBlock
	blocksWithHeight []*AccountBlockWithHeight
	logs             []*api.Logs
	snapshotBlocks   []*SnapshotBlock
	onroadMsgs       []*OnroadMsg
}

type SubscribeApi struct {
	vite        *vite.Vite
	log         log15.Logger
	filterMap   map[rpc.ID]*filter
	filterMapMu sync.Mutex
	eventSystem *EventSystem
}

func NewSubscribeApi(vite *vite.Vite) *SubscribeApi {
	if Es == nil {
		panic("Set \"SubscribeEnabled\" to \"true\" in node_config.json")
	}
	s := &SubscribeApi{
		vite:        vite,
		log:         log15.New("module", "rpc_api/subscribe_api"),
		filterMap:   make(map[rpc.ID]*filter),
		eventSystem: Es,
	}
	go s.timeoutLoop()
	return s
}

func (s *SubscribeApi) timeoutLoop() {
	s.log.Info("start timeout loop")
	// delete timeout filters every 5 minutes
	ticker := time.NewTicker(5 * time.Minute)
	for {
		<-ticker.C
		s.filterMapMu.Lock()
		for id, f := range s.filterMap {
			select {
			case <-f.deadline.C:
				f.s.Unsubscribe()
				delete(s.filterMap, id)
			default:
				continue
			}
		}
		s.filterMapMu.Unlock()
	}
}

type RpcFilterParam struct {
	AddrRange map[string]*api.Range `json:"addrRange"`
	Topics    [][]types.Hash        `json:"topics"`
}

type AccountBlock struct {
	Hash    types.Hash `json:"hash"`
	Removed bool       `json:"removed"`
}

type OnroadMsg struct {
	Hash     types.Hash `json:"hash"`
	Closed   bool       `json:"closed"` // Deprecated: use received instead
	Received bool       `json:"received"`
	Removed  bool       `json:"removed"`
}

type AccountBlockWithHeight struct {
	Hash      types.Hash `json:"hash"`
	Height    uint64     `json:"height"` // Deprecated
	HeightStr string     `json:"heightStr"`
	Removed   bool       `json:"removed"`
}

type SnapshotBlock struct {
	Hash      types.Hash `json:"hash"`
	Height    uint64     `json:"height"` // Deprecated
	HeightStr string     `json:"heightStr"`
	Removed   bool       `json:"removed"`
}

// Deprecated: use subscribe_createSnapshotBlockFilter instead
func (s *SubscribeApi) NewSnapshotBlocksFilter() (rpc.ID, error) {
	return s.createSnapshotBlockFilter()
}
func (s *SubscribeApi) CreateSnapshotBlockFilter() (rpc.ID, error) {
	return s.createSnapshotBlockFilter()
}
func (s *SubscribeApi) createSnapshotBlockFilter() (rpc.ID, error) {
	s.log.Info("createSnapshotBlockFilter")
	var (
		sbCh  = make(chan []*SnapshotBlock)
		sbSub = s.eventSystem.SubscribeSnapshotBlocks(sbCh)
	)

	s.filterMapMu.Lock()
	s.filterMap[sbSub.ID] = &filter{typ: sbSub.sub.typ, deadline: time.NewTimer(deadline), s: sbSub}
	s.filterMapMu.Unlock()

	go func() {
		for {
			select {
			case sb := <-sbCh:
				s.filterMapMu.Lock()
				if f, found := s.filterMap[sbSub.ID]; found {
					f.snapshotBlocks = append(f.snapshotBlocks, sb...)
				}
				s.filterMapMu.Unlock()
			case <-sbSub.Err():
				s.filterMapMu.Lock()
				delete(s.filterMap, sbSub.ID)
				s.filterMapMu.Unlock()
				return
			}
		}
	}()

	return sbSub.ID, nil
}

// Deprecated: use subscribe_createAccountBlockFilter instead
func (s *SubscribeApi) NewAccountBlocksFilter() (rpc.ID, error) {
	return s.createAccountBlockFilter()
}
func (s *SubscribeApi) CreateAccountBlockFilter() (rpc.ID, error) {
	return s.createAccountBlockFilter()
}
func (s *SubscribeApi) createAccountBlockFilter() (rpc.ID, error) {
	s.log.Info("createAccountBlockFilter")
	var (
		acCh  = make(chan []*AccountBlock)
		acSub = s.eventSystem.SubscribeAccountBlocks(acCh)
	)

	s.filterMapMu.Lock()
	s.filterMap[acSub.ID] = &filter{typ: acSub.sub.typ, deadline: time.NewTimer(deadline), s: acSub}
	s.filterMapMu.Unlock()

	go func() {
		for {
			select {
			case ac := <-acCh:
				s.filterMapMu.Lock()
				if f, found := s.filterMap[acSub.ID]; found {
					f.blocks = append(f.blocks, ac...)
				}
				s.filterMapMu.Unlock()
			case <-acSub.Err():
				s.filterMapMu.Lock()
				delete(s.filterMap, acSub.ID)
				s.filterMapMu.Unlock()
				return
			}
		}
	}()

	return acSub.ID, nil
}

// Deprecated: use subscribe_createAccountBlockFilterByAddress instead
func (s *SubscribeApi) NewAccountBlocksByAddrFilter(addr types.Address) (rpc.ID, error) {
	return s.createAccountBlockFilterByAddress(addr)
}
func (s *SubscribeApi) CreateAccountBlockFilterByAddress(addr types.Address) (rpc.ID, error) {
	return s.createAccountBlockFilterByAddress(addr)
}
func (s *SubscribeApi) createAccountBlockFilterByAddress(addr types.Address) (rpc.ID, error) {
	s.log.Info("createAccountBlockFilterByAddress")
	var (
		acCh  = make(chan []*AccountBlockWithHeight)
		acSub = s.eventSystem.SubscribeAccountBlocksByAddr(addr, acCh)
	)

	s.filterMapMu.Lock()
	s.filterMap[acSub.ID] = &filter{typ: acSub.sub.typ, deadline: time.NewTimer(deadline), s: acSub}
	s.filterMapMu.Unlock()

	go func() {
		for {
			select {
			case ac := <-acCh:
				s.filterMapMu.Lock()
				if f, found := s.filterMap[acSub.ID]; found {
					f.blocksWithHeight = append(f.blocksWithHeight, ac...)
				}
				s.filterMapMu.Unlock()
			case <-acSub.Err():
				s.filterMapMu.Lock()
				delete(s.filterMap, acSub.ID)
				s.filterMapMu.Unlock()
				return
			}
		}
	}()

	return acSub.ID, nil
}

// Deprecated: use subscribe_createUnreceivedBlockFilterByAddress instead
func (s *SubscribeApi) NewOnroadBlocksByAddrFilter(addr types.Address) (rpc.ID, error) {
	return s.createUnreceivedBlockFilterByAddress(addr)
}
func (s *SubscribeApi) CreateUnreceivedBlockFilterByAddress(addr types.Address) (rpc.ID, error) {
	return s.createUnreceivedBlockFilterByAddress(addr)
}
func (s *SubscribeApi) createUnreceivedBlockFilterByAddress(addr types.Address) (rpc.ID, error) {
	s.log.Info("createUnreceivedBlockFilterByAddress")
	var (
		acCh  = make(chan []*OnroadMsg)
		acSub = s.eventSystem.SubscribeOnroadBlocksByAddr(addr, acCh)
	)

	s.filterMapMu.Lock()
	s.filterMap[acSub.ID] = &filter{typ: acSub.sub.typ, deadline: time.NewTimer(deadline), s: acSub}
	s.filterMapMu.Unlock()

	go func() {
		for {
			select {
			case ac := <-acCh:
				s.filterMapMu.Lock()
				if f, found := s.filterMap[acSub.ID]; found {
					f.onroadMsgs = append(f.onroadMsgs, ac...)
				}
				s.filterMapMu.Unlock()
			case <-acSub.Err():
				s.filterMapMu.Lock()
				delete(s.filterMap, acSub.ID)
				s.filterMapMu.Unlock()
				return
			}
		}
	}()

	return acSub.ID, nil
}

// Deprecated: use subscribe_createVmLogFilter instead
func (s *SubscribeApi) NewLogsFilter(param RpcFilterParam) (rpc.ID, error) {
	return s.createVmLogFilter(param.AddrRange, param.Topics)
}
func (s *SubscribeApi) CreateVmLogFilter(param api.VmLogFilterParam) (rpc.ID, error) {
	return s.createVmLogFilter(param.AddrRange, param.Topics)
}
func (s *SubscribeApi) createVmLogFilter(rangeMap map[string]*api.Range, topics [][]types.Hash) (rpc.ID, error) {
	s.log.Info("createVmLogFilter")
	p, err := api.ToFilterParam(rangeMap, topics)
	if err != nil {
		return "", err
	}
	var (
		logsCh  = make(chan []*api.Logs)
		logsSub = s.eventSystem.SubscribeLogs(p, logsCh)
	)

	s.filterMapMu.Lock()
	s.filterMap[logsSub.ID] = &filter{typ: logsSub.sub.typ, deadline: time.NewTimer(deadline), s: logsSub}
	s.filterMapMu.Unlock()

	go func() {
		for {
			select {
			case l := <-logsCh:
				s.filterMapMu.Lock()
				if f, found := s.filterMap[logsSub.ID]; found {
					f.logs = append(f.logs, l...)
				}
				s.filterMapMu.Unlock()
			case <-logsSub.Err():
				s.filterMapMu.Lock()
				delete(s.filterMap, logsSub.ID)
				s.filterMapMu.Unlock()
				return
			}
		}
	}()

	return logsSub.ID, nil
}

func (s *SubscribeApi) UninstallFilter(id rpc.ID) bool {
	s.log.Info("UninstallFilter")
	s.filterMapMu.Lock()
	f, found := s.filterMap[id]
	if found {
		delete(s.filterMap, id)
	}
	s.filterMapMu.Unlock()
	if found {
		f.s.Unsubscribe()
	}
	return found
}

type AccountBlocksMsg struct {
	Blocks []*AccountBlock `json:"result"`
	Id     rpc.ID          `json:"subscription"`
}

type AccountBlocksWithHeightMsg struct {
	Blocks []*AccountBlockWithHeight `json:"result"`
	Id     rpc.ID                    `json:"subscription"`
}

type LogsMsg struct {
	Logs []*api.Logs `json:"result"`
	Id   rpc.ID      `json:"subscription"`
}

type OnroadBlocksMsg struct {
	Blocks []*OnroadMsg `json:"result"`
	Id     rpc.ID       `json:"subscription"`
}

type SnapshotBlocksMsg struct {
	Blocks []*SnapshotBlock `json:"result"`
	Id     rpc.ID           `json:"subscription"`
}

// Deprecated: use subscribe_getChangesByFilterId instead
func (s *SubscribeApi) GetFilterChanges(id rpc.ID) (interface{}, error) {
	return s.getChangesByFilterId(id)
}
func (s *SubscribeApi) GetChangesByFilterId(id rpc.ID) (interface{}, error) {
	return s.getChangesByFilterId(id)
}
func (s *SubscribeApi) getChangesByFilterId(id rpc.ID) (interface{}, error) {
	s.log.Info("getChangesByFilterId", "id", id)
	s.filterMapMu.Lock()
	defer s.filterMapMu.Unlock()

	if f, found := s.filterMap[id]; found {
		if !f.deadline.Stop() {
			<-f.deadline.C
		}
		f.deadline.Reset(deadline)

		switch f.typ {
		case AccountBlocksSubscription:
			blocks := f.blocks
			f.blocks = nil
			return AccountBlocksMsg{blocks, id}, nil
		case AccountBlocksWithHeightSubscription:
			blocks := f.blocksWithHeight
			f.blocksWithHeight = nil
			return AccountBlocksWithHeightMsg{blocks, id}, nil
		case OnroadBlocksSubscription:
			onroadMsgs := f.onroadMsgs
			f.onroadMsgs = nil
			return OnroadBlocksMsg{onroadMsgs, id}, nil
		case LogsSubscription:
			logs := f.logs
			f.logs = nil
			return LogsMsg{logs, id}, nil
		case SnapshotBlocksSubscription:
			snapshotBlocks := f.snapshotBlocks
			f.snapshotBlocks = nil
			return SnapshotBlocksMsg{snapshotBlocks, id}, nil
		}
	}

	return nil, errors.New("filter not found")
}

// Deprecated: use subscribe_createSnapshotBlockSubscription instead
func (s *SubscribeApi) NewSnapshotBlocks(ctx context.Context) (*rpc.Subscription, error) {
	return s.createSnapshotBlockSubscription(ctx)
}
func (s *SubscribeApi) CreateSnapshotBlockSubscription(ctx context.Context) (*rpc.Subscription, error) {
	return s.createSnapshotBlockSubscription(ctx)
}
func (s *SubscribeApi) createSnapshotBlockSubscription(ctx context.Context) (*rpc.Subscription, error) {
	s.log.Info("createSnapshotBlockSubscription")
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}
	rpcSub := notifier.CreateSubscription()

	go func() {
		snapshotBlockHashChan := make(chan []*SnapshotBlock, 128)
		sbSub := s.eventSystem.SubscribeSnapshotBlocks(snapshotBlockHashChan)
		for {
			select {
			case h := <-snapshotBlockHashChan:
				notifier.Notify(rpcSub.ID, h)
			case <-rpcSub.Err():
				sbSub.Unsubscribe()
				return
			case <-notifier.Closed():
				sbSub.Unsubscribe()
				return
			}
		}
	}()

	return rpcSub, nil
}

// Deprecated: use subscribe_createAccountBlockSubscription instead
func (s *SubscribeApi) NewAccountBlocks(ctx context.Context) (*rpc.Subscription, error) {
	return s.createAccountBlockSubscription(ctx)
}
func (s *SubscribeApi) CreateAccountBlockSubscription(ctx context.Context) (*rpc.Subscription, error) {
	return s.createAccountBlockSubscription(ctx)
}
func (s *SubscribeApi) createAccountBlockSubscription(ctx context.Context) (*rpc.Subscription, error) {
	s.log.Info("createAccountBlockSubscription")
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}
	rpcSub := notifier.CreateSubscription()

	go func() {
		accountBlockHashCh := make(chan []*AccountBlock, 128)
		acSub := s.eventSystem.SubscribeAccountBlocks(accountBlockHashCh)
		for {
			select {
			case h := <-accountBlockHashCh:
				notifier.Notify(rpcSub.ID, h)
			case <-rpcSub.Err():
				acSub.Unsubscribe()
				return
			case <-notifier.Closed():
				acSub.Unsubscribe()
				return
			}
		}
	}()

	return rpcSub, nil
}

// Deprecated: use subscribe_createAccountBlockSubscriptionByAddress instead
func (s *SubscribeApi) NewAccountBlocksByAddr(ctx context.Context, addr types.Address) (*rpc.Subscription, error) {
	return s.createAccountBlockSubscriptionByAddress(ctx, addr)
}
func (s *SubscribeApi) CreateAccountBlockSubscriptionByAddress(ctx context.Context, addr types.Address) (*rpc.Subscription, error) {
	return s.createAccountBlockSubscriptionByAddress(ctx, addr)
}
func (s *SubscribeApi) createAccountBlockSubscriptionByAddress(ctx context.Context, addr types.Address) (*rpc.Subscription, error) {
	s.log.Info("createAccountBlockSubscriptionByAddress")
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}
	rpcSub := notifier.CreateSubscription()

	go func() {
		accountBlockCh := make(chan []*AccountBlockWithHeight, 128)
		acSub := s.eventSystem.SubscribeAccountBlocksByAddr(addr, accountBlockCh)
		for {
			select {
			case h := <-accountBlockCh:
				notifier.Notify(rpcSub.ID, h)
			case <-rpcSub.Err():
				acSub.Unsubscribe()
				return
			case <-notifier.Closed():
				acSub.Unsubscribe()
				return
			}
		}
	}()

	return rpcSub, nil
}

// Deprecated: use subscribe_createUnreceivedBlockSubscriptionByAddress instead
func (s *SubscribeApi) NewOnroadBlocksByAddr(ctx context.Context, addr types.Address) (*rpc.Subscription, error) {
	return s.createUnreceivedBlockSubscriptionByAddress(ctx, addr)
}
func (s *SubscribeApi) CreateUnreceivedBlockSubscriptionByAddress(ctx context.Context, addr types.Address) (*rpc.Subscription, error) {
	return s.createUnreceivedBlockSubscriptionByAddress(ctx, addr)
}
func (s *SubscribeApi) createUnreceivedBlockSubscriptionByAddress(ctx context.Context, addr types.Address) (*rpc.Subscription, error) {
	s.log.Info("createUnreceivedBlockSubscriptionByAddress")
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}
	rpcSub := notifier.CreateSubscription()

	go func() {
		accountBlockHashCh := make(chan []*OnroadMsg, 128)
		acSub := s.eventSystem.SubscribeOnroadBlocksByAddr(addr, accountBlockHashCh)
		for {
			select {
			case h := <-accountBlockHashCh:
				notifier.Notify(rpcSub.ID, h)
			case <-rpcSub.Err():
				acSub.Unsubscribe()
				return
			case <-notifier.Closed():
				acSub.Unsubscribe()
				return
			}
		}
	}()

	return rpcSub, nil
}

// Deprevated: use subscribe_createVmLogSubscription instead
func (s *SubscribeApi) NewLogs(ctx context.Context, param RpcFilterParam) (*rpc.Subscription, error) {
	return s.createVmLogSubscription(ctx, param.AddrRange, param.Topics)
}
func (s *SubscribeApi) CreateVmLogSubscription(ctx context.Context, param api.VmLogFilterParam) (*rpc.Subscription, error) {
	return s.createVmLogSubscription(ctx, param.AddrRange, param.Topics)
}
func (s *SubscribeApi) createVmLogSubscription(ctx context.Context, rangeMap map[string]*api.Range, topics [][]types.Hash) (*rpc.Subscription, error) {
	s.log.Info("createVmLogSubscription")
	p, err := api.ToFilterParam(rangeMap, topics)
	if err != nil {
		return nil, err
	}

	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return &rpc.Subscription{}, rpc.ErrNotificationsUnsupported
	}
	rpcSub := notifier.CreateSubscription()

	go func() {
		logsMsg := make(chan []*api.Logs, 128)
		sub := s.eventSystem.SubscribeLogs(p, logsMsg)

		for {
			select {
			case msg := <-logsMsg:
				notifier.Notify(rpcSub.ID, msg)
			case <-rpcSub.Err():
				sub.Unsubscribe()
				return
			case <-notifier.Closed():
				sub.Unsubscribe()
				return
			}
		}
	}()
	return rpcSub, nil
}

// Deprecated: use ledger_getVmLogsByFilter instead
func (s *SubscribeApi) GetLogs(param RpcFilterParam) ([]*api.Logs, error) {
	return api.GetLogs(s.vite.Chain(), param.AddrRange, param.Topics)
}
