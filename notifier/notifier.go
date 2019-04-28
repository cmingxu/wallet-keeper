package notifier

import (
	log "github.com/sirupsen/logrus"
)

// Engine of notification
type Notifier struct {
	receivers map[string]*Receiver
	eventChan chan Event

	stopCh chan struct{}
}

// event generation buffer
const EventBufSize = 10

func New() *Notifier {
	return &Notifier{
		receivers: make(map[string]*Receiver),
		eventChan: make(chan Event, EventBufSize)}
}

func (notifier *Notifier) InstallReceiver(name string, receiver *Receiver) {
	notifier.receivers[name] = receiver
}

func (notifier *Notifier) UninstallReceiver(name string) {
	delete(notifier.receivers, name)
}

func (notifier *Notifier) ListReceivers() map[string]*Receiver {
	return notifier.receivers
}

func (notifier *Notifier) EventChan() chan<- Event {
	return notifier.eventChan
}

func (notifier *Notifier) Stop() {
	close(notifier.stopCh)
}

func (notifier *Notifier) Start() {
	for {
		select {
		case <-notifier.stopCh:
			return

		case event, _ := <-notifier.eventChan:
			for name, receiver := range notifier.receivers {
				if receiver.Match(event) {
					log.Debugf(name, "received", event)
					receiver.Accept(event)
				}
			}
		}
	}
}
