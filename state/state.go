package state

type State interface {
	canPay() error //ver un nombre mejor
	canReject() error
	canClose() error
	// Meter una funcion o atributo que te devuelva el nombre
}
