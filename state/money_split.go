package state

import "fmt"

type MoneySplit struct {
	Pending       State
	ActivePaid    State
	ActiveRefused State
	Active        State
	Closed        State
	Finished      State
	Refused       State
	Paid          State

	CurrentState State

	TotalMembers    int
	PaidMembers     int
	RejectedMembers int
}

//Hacer otro NewMoneSplit que se construya solo en base al id de evento.
//lo obtiene de la DB(EventWithdetails)
//valida que sea un split
//con la cantidad de detalles obtenidos y segun el estado de cada uno llama a NewMoneySplit(totalMembers int, paidMembers int)

func NewMoneySplit(totalMembers int, paidMembers int) *MoneySplit {
	//TODO: validar que los totalMemberssean mas que los paidMembers
	split := &MoneySplit{
		TotalMembers:    totalMembers,
		PaidMembers:     paidMembers,
		RejectedMembers: 0,
	}
	pending := &Pending{
		MoneySplit: split,
	}
	activePaid := &ActivePaid{
		MoneySplit: split,
	}
	activeRefused := &ActiveRefused{
		MoneySplit: split,
	}
	active := &Active{
		MoneySplit: split,
	}
	closed := &Closed{
		MoneySplit: split,
	}
	finished := &Finished{
		MoneySplit: split,
	}
	refused := &Refused{
		MoneySplit: split,
	}
	paid := &Paid{
		MoneySplit: split,
	}

	split.SetState(pending)
	split.Pending = pending
	split.ActivePaid = activePaid
	split.ActiveRefused = activeRefused
	split.Active = active
	split.Closed = closed
	split.Finished = finished
	split.Refused = refused
	split.Paid = paid
	return split
}

func (split *MoneySplit) PaySplit() error {
	return split.CurrentState.PaySplit()
}

func (split *MoneySplit) RejectSplit() error {
	return split.CurrentState.RejectSplit()
}

func (split *MoneySplit) CloseSplit() error {
	return split.CurrentState.CloseSplit()
}

func (split *MoneySplit) SetState(s State) {
	split.CurrentState = s
}

func (split *MoneySplit) IncrementPaidMembers() {
	split.incrementPaidMembers(1)
}

func (split *MoneySplit) incrementPaidMembers(newPaidMembers int) {
	fmt.Printf("Adding %d paids\n", newPaidMembers)
	split.PaidMembers = split.PaidMembers + newPaidMembers
}

func (split *MoneySplit) IncrementRejectedMembers() {
	split.incrementRejectedMembers(1)
}

func (split *MoneySplit) incrementRejectedMembers(newRejectedMembers int) {
	fmt.Printf("Adding %d rejects\n", newRejectedMembers)
	split.RejectedMembers = split.RejectedMembers + newRejectedMembers
}
