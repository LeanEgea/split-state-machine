package state

import "fmt"

type Active struct {
	MoneySplit *MoneySplit
}

func (a *Active) PaySplit() error {
	split := a.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementPaidMembers()
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementPaidMembers()
		split.SetState(split.Finished)
		return nil
	} else {
		return fmt.Errorf("Can't pay the split(active)")
	}
}

func (a *Active) RejectSplit() error {
	split := a.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementRejectedMembers()
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.IncrementRejectedMembers()
		split.SetState(split.Finished)
		return nil
	} else {
		return fmt.Errorf("Can't reject the split(active)")
	}
}

func (a *Active) CloseSplit() error {
	split := a.MoneySplit
	split.SetState(split.Closed)
	return nil
}
