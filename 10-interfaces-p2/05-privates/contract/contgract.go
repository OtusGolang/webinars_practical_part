package contract

type Contract interface {
	Sign() error
	cancel() error
}
