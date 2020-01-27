package main

import (
	"fmt"
	"log"

	"github.com/mercadolibre/state-machine-split/state"
)

func main() {
	fmt.Println("New split")
	moneySplit := state.NewMoneySplit(5, 1)
	err := moneySplit.PaySplit()
	if err != nil {
		log.Fatalf(err.Error())
	}
	// err = moneySplit.CloseSplit()
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	err = moneySplit.PaySplit()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = moneySplit.PaySplit()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = moneySplit.PaySplit()
	if err != nil {
		log.Fatalf(err.Error())
	}
	// err = moneySplit.CloseSplit()
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	err = moneySplit.PaySplit()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("There is %d members\n", moneySplit.TotalMembers)
	fmt.Printf("There is %d paid members\n", moneySplit.PaidMembers)
	fmt.Printf("There is %d rejected members\n", moneySplit.RejectedMembers)
}
