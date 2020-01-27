package state

import "fmt"

type Pending struct {
	MoneySplit *MoneySplit
}

func (p *Pending) PaySplit() error {
	split := p.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementPaidMembers()
		split.SetState(split.ActivePaid)
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementPaidMembers()
		split.SetState(split.Paid)
		return nil
	} else {
		return fmt.Errorf("Can't pay the split(pending)")
	}
}

func (p *Pending) RejectSplit() error {
	split := p.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementRejectedMembers()
		split.SetState(split.ActiveRefused)
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementRejectedMembers()
		split.SetState(split.Refused)
		return nil
	} else {
		return fmt.Errorf("Can't reject the split(pending)")
	}
}

func (p *Pending) CloseSplit() error {
	split := p.MoneySplit
	split.SetState(split.Closed)
	return nil
}
