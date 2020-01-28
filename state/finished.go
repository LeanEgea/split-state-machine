package state

import "fmt"

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
