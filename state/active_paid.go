package state

import "fmt"

type ActivePaid struct {
	MoneySplit *MoneySplit
}

func (a *ActivePaid) canPay() error {
	split := a.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.setState(split.Paid)
		return nil
	} else {
		return fmt.Errorf("Can't pay the split(active_paid)")
	}
}

func (a *ActivePaid) canReject() error {
	split := a.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		split.setState(split.Active)
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.setState(split.Finished)
		return nil
	} else {
		return fmt.Errorf("Can't reject the split(active_paid)")
	}
}

func (a *ActivePaid) canClose() error {
	split := a.MoneySplit
	split.setState(split.Closed)
	return nil
}
