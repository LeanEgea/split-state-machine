package state

import (
	"fmt"

	"github.com/mercadolibre/state-machine-split/utils"
)

type Closed struct {
	MoneySplit *MoneySplit
}

func (c *Closed) canPay() error {
	return fmt.Errorf("The split is already closed")
}

func (c *Closed) canReject() error {
	return fmt.Errorf("The split is already closed")
}

func (c *Closed) canClose() error {
	return fmt.Errorf("The split is already closed")
}

func (c *Closed) stateName() string {
	return utils.Closed
}
