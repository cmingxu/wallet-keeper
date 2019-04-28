package notifier

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

var ErrShouldRetry = errors.New("should retry")

func ShouldRetry(err error) bool {
	return err.Error() == "should retry"
}

// Receiver receive event notfication
type Receiver struct {
	retryCount uint     `json:"retryCount"`
	endpoint   string   `json:"endpoint"`
	eventTypes []string `json:"evnetTypes"`

	client *http.Client `json:"-"`
}

func NewReceiver(endpoint string, eventTypes []string, retriesCount uint) *Receiver {
	return &Receiver{
		retryCount: retriesCount,
		endpoint:   endpoint,
		eventTypes: eventTypes,
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: true,
			},
		},
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
	sendFunc := func(event Event) error {
		buf := bytes.NewBufferString("")
		err := json.NewEncoder(buf).Encode(event.GetEvent())
		if err != nil {
			log.Error(err)
			return err
		}

		post, _ := http.NewRequest(http.MethodPost, r.endpoint, buf)
		resp, err := r.client.Do(post)
		if err != nil {
			return ErrShouldRetry
		}

		if resp.StatusCode != http.StatusOK {
			return ErrShouldRetry
		}

		return nil
	}

	go func() {
		retryRemain := r.retryCount
		ticker := time.NewTicker(10 * time.Second)
		for {
			select {
			case now := <-ticker.C:
				log.Println(now)
				err := sendFunc(event)
				if err == nil {
					return
				} else {
					if !ShouldRetry(err) {
						return
					}

					if retryRemain > 0 {
						retryRemain = retryRemain - 1
					} else {
						return
					}
				}
			}
		}
	}()
}
