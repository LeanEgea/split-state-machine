package state

import (
	"fmt"

	"github.com/mercadolibre/state-machine-split/utils"
)

type Finished struct {
	MoneySplit *MoneySplit
}

func (f *Finished) canPay() error {
	return fmt.Errorf("The split is already finished")
}

func (f *Finished) canReject() error {
	return fmt.Errorf("The split is already finished")
}

func (f *Finished) canClose() error {
	return fmt.Errorf("The split is already finished")
}

func (f *Finished) stateName() string {
	return utils.Finished
}
