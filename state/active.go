package state

import "fmt"

type Active struct {
	MoneySplit *MoneySplit
}

func (a *Active) canPay() error {
	split := a.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.setState(split.Finished)
		return nil
	} else {
		return fmt.Errorf("Can't pay the split(active)")
	}
}

func (a *Active) canReject() error {
	split := a.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.setState(split.Finished)
		return nil
	} else {
		return fmt.Errorf("Can't reject the split(active)")
	}
}

func (a *Active) canClose() error {
	split := a.MoneySplit
	split.setState(split.Closed)
	return nil
}
