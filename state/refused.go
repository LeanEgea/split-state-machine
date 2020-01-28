package state

import (
	"fmt"

	"github.com/mercadolibre/state-machine-split/utils"
)

type Refused struct {
	MoneySplit *MoneySplit
}

func (r *Refused) canPay() error {
	return fmt.Errorf("The split is already rejected")
}

func (r *Refused) canReject() error {
	return fmt.Errorf("The split is already rejected")
}

func (r *Refused) canClose() error {
	return fmt.Errorf("The split is already rejected")
}

func (r *Refused) stateName() string {
	return utils.Refused
}
