package state

import "fmt"

type ActiveRefused struct {
	MoneySplit *MoneySplit
}

func (a *ActiveRefused) PaySplit() error {
	split := a.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementPaidMembers()
		split.SetState(split.Active)
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementPaidMembers()
		split.SetState(split.Finished)
		return nil
	} else {
		return fmt.Errorf("Can't pay the split(active_refused)")
	}
}

func (a *ActiveRefused) RejectSplit() error {
	split := a.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementRejectedMembers()
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementRejectedMembers()
		split.SetState(split.Refused)
		return nil
	} else {
		return fmt.Errorf("Can't reject the split(active_refused)")
	}
}

func (a *ActiveRefused) CloseSplit() error {
	split := a.MoneySplit
	split.SetState(split.Closed)
	return nil
}
