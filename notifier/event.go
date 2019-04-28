package notifier

type Event interface {
	GetEvent() map[string]interface{}
	Type() string
}
