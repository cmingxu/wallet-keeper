package omnilayer

type ErrClientShutdown struct{}

func (*ErrClientShutdown) Error() string {
	return "the client has been shutdown"
}

func errClientShutdown() error {
	return &ErrClientShutdown{}
}
