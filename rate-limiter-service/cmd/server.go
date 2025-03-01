package main

import (
	"leaky-bucket/internal/infra"
	"leaky-bucket/internal/infra/buckets"
	"log"
)

func main() {
	prometheusService := infra.NewPrometheus()
	leakyBucket := buckets.NewInMemoryLeakyBucket(10)

	app := infra.NewServer("localhost", "3010", prometheusService, leakyBucket)
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
