package state

type State interface {
	canPay() error
	canReject() error
	canClose() error
	stateName() string
}
