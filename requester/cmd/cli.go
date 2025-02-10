package main

import (
	"requester/internal"
	"requester/internal/infra"
)

func main() {
	logger := infra.NewLogger()

	logger.Info("⚡ Requester 🔄")

	params := internal.NewParamsFromFlags()
	requester := internal.NewRequester(params, logger)
	requester.Run()
}
