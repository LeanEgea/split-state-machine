package state

import "fmt"

type ActivePaid struct {
	MoneySplit *MoneySplit
}

func (a *ActivePaid) PaySplit() error {
	split := a.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementPaidMembers()
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementPaidMembers()
		split.SetState(split.Paid)
		return nil
	} else {
		return fmt.Errorf("Can't pay the split(active_paid)")
	}
}

func (a *ActivePaid) RejectSplit() error {
	split := a.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementRejectedMembers()
		split.SetState(split.Active)
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementRejectedMembers()
		split.SetState(split.Finished)
		return nil
	} else {
		return fmt.Errorf("Can't reject the split(active_paid)")
	}
}

func (a *ActivePaid) CloseSplit() error {
	split := a.MoneySplit
	split.SetState(split.Closed)
	return nil
}
