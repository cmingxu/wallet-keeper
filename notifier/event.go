package notifier

type Event interface {
	GetEvent() map[string]string
	Type() string
}
