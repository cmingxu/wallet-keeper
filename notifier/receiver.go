package notifier

import (
	log "github.com/sirupsen/logrus"
)

// Receiver receive event notfication
type Receiver struct {
	retryCount uint
	endpoint   string
	eventTypes []string
}

func NewReceiver(endpoint string, eventTypes []string, retriesCount uint) *Receiver {
	return &Receiver{
		retryCount: retriesCount,
		endpoint:   endpoint,
		eventTypes: eventTypes,
	}
}

func (r *Receiver) Match(event Event) bool {
	for _, et := range r.eventTypes {
		if et == event.Type() {
			return true
		}
	}

	return false
}

func (r *Receiver) Accept(event Event) {
	log.Println(event)
}
