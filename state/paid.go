package state

import "fmt"

type Paid struct {
	MoneySplit *MoneySplit
}

func (p *Paid) canPay() error {
	return fmt.Errorf("The split is already paid")
}

func (p *Paid) canReject() error {
	return fmt.Errorf("The split is already paid")
}

func (p *Paid) canClose() error {
	return fmt.Errorf("The split is already paid")
}
