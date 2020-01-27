package state

import "fmt"

type Closed struct {
	MoneySplit *MoneySplit
}

func (c *Closed) PaySplit() error {
	return fmt.Errorf("The split is already closed")
}

func (c *Closed) RejectSplit() error {
	return fmt.Errorf("The split is already closed")
}

func (c *Closed) CloseSplit() error {
	return fmt.Errorf("The split is already closed")
}
