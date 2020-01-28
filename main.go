package main

import (
	"fmt"
	"log"

	"github.com/mercadolibre/state-machine-split/state"
)

func main() {
	fmt.Println("New split")
	moneySplit := state.NewMoneySplit(5, 1, 0)
	fmt.Println(moneySplit.State())

	err := moneySplit.CanPay()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(moneySplit.State())

	// err = moneySplit.CanClose()
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// fmt.Println(moneySplit.State())

	err = moneySplit.CanPay()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(moneySplit.State())

	err = moneySplit.CanPay()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(moneySplit.State())

	err = moneySplit.CanPay()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(moneySplit.State())

	// err = moneySplit.CanClose()
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// fmt.Println(moneySplit.State)

	err = moneySplit.CanPay()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(moneySplit.State())

	fmt.Printf("There is %d members\n", moneySplit.TotalMembers)
	fmt.Printf("There is %d paid members\n", moneySplit.PaidMembers)
	fmt.Printf("There is %d rejected members\n", moneySplit.RejectedMembers)
}
