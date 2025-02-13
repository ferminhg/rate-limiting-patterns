package internal

import (
	"flag"
)

type Params struct {
	NumRequests int
	Host        string
}

func NewParams(numRequests int, host string) *Params {
	return &Params{
		NumRequests: numRequests,
		Host:        host,
	}
}

func NewParamsFromFlags() *Params {
	numRequests := flag.Int("n", 10, "number of requests to send")
	host := flag.String("h", "http://localhost:3010", "host to send the requests to")
	flag.Parse()

	return NewParams(*numRequests, *host)
}
