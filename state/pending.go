package state

import "fmt"

type Pending struct {
	MoneySplit *MoneySplit
}

func (p *Pending) canPay() error {
	split := p.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		split.setState(split.ActivePaid)
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.setState(split.Paid)
		return nil
	} else {
		return fmt.Errorf("Can't pay the split(pending)")
	}
}

func (p *Pending) canReject() error {
	split := p.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		split.setState(split.ActiveRefused)
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.setState(split.Refused)
		return nil
	} else {
		return fmt.Errorf("Can't reject the split(pending)")
	}
}

func (p *Pending) canClose() error {
	split := p.MoneySplit
	split.setState(split.Closed)
	return nil
}
