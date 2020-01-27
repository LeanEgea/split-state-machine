package state

type State interface {
	PaySplit() error
	RejectSplit() error
	CloseSplit() error
	// Meter una funcion o atributo que te devuelva el nombre
}
