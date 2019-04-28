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

// Check if event type in receivier's eventTypes
func (r *Receiver) Match(event Event) bool {
	for _, et := range r.eventTypes {
		if et == event.Type() {
			return true
		}
	}

	return false
}

// Accept event and spawn new goroutine to post event back to the endpoint.
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

		// should retry if endpoint does not return status code 200
		if resp.StatusCode != http.StatusOK {
			return ErrShouldRetry
		}

		return nil
	}

	go func() {
		retryRemains := r.retryCount
		ticker := time.NewTicker(10 * time.Second)
		for {
			select {
			case now := <-ticker.C:
				log.Println(now)
				err := sendFunc(event)
				if err == nil {
					return
				} else {
					// stop retrying if serious error happend
					if !ShouldRetry(err) {
						return
					}

					if retryRemains <= 0 {
						log.Debugf("stop posting event after n retries")
						return
					}

					retryRemains = retryRemains - 1
				}
			}
		}
	}()
}
