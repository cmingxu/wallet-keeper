package omnilayer

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync/atomic"
	"time"
)

const (
	contentType     = "Content-Type"
	contentTypeJSON = "application/json"
)

type Client struct {
	id           uint64
	config       *ConnConfig
	httpClient   *http.Client
	sendPostChan chan *sendPostDetails

	shutdown chan struct{}
	done     chan struct{}
}

func (c *Client) do(cmd command) chan *response {
	body, err := marshalCmd(cmd)
	if err != nil {
		return newFutureError(err)
	}

	responseChan := make(chan *response, 1)
	jReq := &jsonRequest{
		id:           c.NextID(),
		cmd:          cmd,
		body:         body,
		responseChan: responseChan,
	}

	c.sendPost(jReq)

	return responseChan
}

func (c *Client) sendPost(jReq *jsonRequest) {
	req, err := http.NewRequest(http.MethodPost, "http://"+c.config.Host, bytes.NewReader(jReq.body))
	if err != nil {
		jReq.responseChan <- &response{result: nil, err: err}
		return
	}

	req.Close = true
	req.Header.Set(contentType, contentTypeJSON)
	req.SetBasicAuth(c.config.User, c.config.Pass)

	select {
	case <-c.shutdown:
		jReq.responseChan <- &response{err: errClientShutdown()}
	default:
		c.sendPostChan <- &sendPostDetails{
			jsonRequest: jReq,
			httpRequest: req,
		}
	}
}

func New(config *ConnConfig) *Client {
	httpClient := newHTTPClient()

	client := &Client{
		config:       config,
		httpClient:   httpClient,
		sendPostChan: make(chan *sendPostDetails, sendPostBufferSize),

		shutdown: make(chan struct{}, 1),
		done:     make(chan struct{}, 1),
	}

	go client.sendPostHandler()

	return client
}

func (c *Client) sendPostHandler() {
out:
	for {
		select {
		case details := <-c.sendPostChan:
			c.handleSendPostMessage(details)

		case <-c.shutdown:
			break out
		}
	}

cleanup:
	for {
		select {
		case details := <-c.sendPostChan:
			details.jsonRequest.responseChan <- &response{
				result: nil,
				err:    errClientShutdown(),
			}

		default:
			break cleanup
		}
	}

	close(c.done)
}

func (c *Client) handleSendPostMessage(details *sendPostDetails) {
	jReq := details.jsonRequest
	httpResponse, err := c.httpClient.Do(details.httpRequest)
	if err != nil {
		jReq.responseChan <- &response{err: err}
		return
	}

	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		jReq.responseChan <- &response{err: err}
		return
	}
	err = httpResponse.Body.Close()
	if err != nil {
		jReq.responseChan <- &response{err: err}
		return
	}

	var resp rawResponse
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		jReq.responseChan <- &response{err: err}
		return
	}

	res, err := resp.result()
	jReq.responseChan <- &response{result: res, err: err}
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: 5 * time.Second,
			ExpectContinueTimeout: 4 * time.Second,
			IdleConnTimeout:       5 * 60 * time.Second,
		},
	}
}

func (c *Client) NextID() uint64 {
	return atomic.AddUint64(&c.id, 1)
}

func (c *Client) Shutdown() {
	select {
	case <-c.shutdown:
		return
	default:
	}

	close(c.shutdown)
	<-c.done
}
