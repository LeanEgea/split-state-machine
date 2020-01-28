package state

import (
	"fmt"

	"github.com/mercadolibre/state-machine-split/utils"
)

type ActiveRefused struct {
	MoneySplit *MoneySplit
}

func (a *ActiveRefused) canPay() error {
	split := a.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		split.setState(split.Active)
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.setState(split.Finished)
		return nil
	} else {
		return fmt.Errorf("Can't pay the split(active_refused)")
	}
}

func (a *ActiveRefused) canReject() error {
	split := a.MoneySplit
	if split.TotalMembers > (split.PaidMembers + split.RejectedMembers + 1) {
		return nil
	} else if split.TotalMembers == (split.PaidMembers + split.RejectedMembers + 1) {
		split.setState(split.Refused)
		return nil
	} else {
		return fmt.Errorf("Can't reject the split(active_refused)")
	}
}

func (a *ActiveRefused) canClose() error {
	split := a.MoneySplit
	split.setState(split.Closed)
	return nil
}

func (a *ActiveRefused) stateName() string {
	return utils.ActiveRefused
}
