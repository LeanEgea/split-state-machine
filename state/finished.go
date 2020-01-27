package state

import "fmt"

type Finished struct {
	MoneySplit *MoneySplit
}

func (f *Finished) PaySplit() error {
	return fmt.Errorf("The split is already finished")
}

func (f *Finished) RejectSplit() error {
	return fmt.Errorf("The split is already finished")
}

func (f *Finished) CloseSplit() error {
	return fmt.Errorf("The split is already finished")
}
