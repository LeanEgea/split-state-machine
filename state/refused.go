package state

import "fmt"

type Refused struct {
	MoneySplit *MoneySplit
}

func (r *Refused) PaySplit() error {
	return fmt.Errorf("The split is already rejected")
}

func (r *Refused) RejectSplit() error {
	return fmt.Errorf("The split is already rejected")
}

func (r *Refused) CloseSplit() error {
	return fmt.Errorf("The split is already rejected")
}
