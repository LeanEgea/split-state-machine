package state

import "fmt"

type Paid struct {
	MoneySplit *MoneySplit
}

func (p *Paid) PaySplit() error {
	return fmt.Errorf("The split is already paid")
}

func (p *Paid) RejectSplit() error {
	return fmt.Errorf("The split is already paid")
}

func (p *Paid) CloseSplit() error {
	return fmt.Errorf("The split is already paid")
}
