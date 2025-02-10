package internal

import (
	"fmt"
	"net/http"
	"requester/internal/infra"
	"sync"
)

type Requester struct {
	params *Params
	logger infra.Logger
	client *http.Client
}

func NewRequester(params *Params, logger infra.Logger) *Requester {
	return &Requester{
		params: params,
		logger: logger,
		client: &http.Client{},
	}
}

func (r *Requester) Run() {
	r.logStartMessage()
	r.sendRequestsConcurrently()
}

func (r *Requester) logStartMessage() {
	r.logger.Info(fmt.Sprintf("🏁 Sending %d requests to %s", r.params.NumRequests, r.params.Host))
}

func (r *Requester) sendRequestsConcurrently() {
	var wg sync.WaitGroup
	for i := 0; i < r.params.NumRequests; i++ {
		wg.Add(1)
		go r.executeRequest(&wg)
	}
	wg.Wait()
}

func (r *Requester) executeRequest(wg *sync.WaitGroup) {
	defer wg.Done()
	r.logger.Info(fmt.Sprintf("Sending request to %s", r.params.Host))

	response, err := r.client.Get(r.params.Host)
	if err != nil {
		r.logger.Error(fmt.Sprintf("Error sending request: %s", err))
		return
	}
	defer response.Body.Close()

	r.logger.Info(fmt.Sprintf("Response: %s", response.Status))
}
