package state

import "fmt"

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
