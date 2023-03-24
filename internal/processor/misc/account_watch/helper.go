package account_watcher

import (
	"github.com/LampardNguyen234/whale-alert/common"
	"github.com/LampardNguyen234/whale-alert/internal/store"
)

func (p *AccountWatchProcessor) getTxMonitoredDetails(from, to string) (string, string, Direction, bool) {
	toDetail, _ := p.getMonitoredAccountDetail(to)
	fromDetail, _ := p.getMonitoredAccountDetail(from)
	d := directionIn

	if toDetail == nil && fromDetail == nil {
		return "", "", d, false
	}

	if fromDetail != nil {
		from = fromDetail.String()
		d = directionOut
	}
	if toDetail != nil {
		to = toDetail.String()
	}
	if fromDetail != nil && toDetail != nil {
		d = directionBothWay
	}

	return from, to, d, true
}

func (p *AccountWatchProcessor) getMonitoredAccountDetail(addr string) (*store.AccountDetail, error) {
	addr, err := common.AccountAddressToHex(addr)
	if err != nil {
		return nil, err
	}

	tmpAccount, ok := p.cachedAccounts.Get(addr)
	if !ok {
		d, err := p.Db.GetAccountDetail(addr)
		if err != nil {
			return nil, err
		}

		p.cachedAccounts.Add(d.Address, d)
		return d, nil
	}

	d := tmpAccount.(*store.AccountDetail)
	return d, nil
}

func (p *AccountWatchProcessor) loadMonitoredAccounts() error {
	allAccounts, err := p.Db.GetAllMonitoredAccounts()
	if err != nil {
		return err
	}

	for addr, d := range allAccounts {
		p.cachedAccounts.Add(addr, d)
	}

	return nil
}
