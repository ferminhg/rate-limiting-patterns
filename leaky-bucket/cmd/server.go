package main

import (
	"leaky-bucket/internal/infra"
	"log"
)

func main() {
	prometheusService := infra.NewPrometheus()
	app := infra.NewServer("localhost", "3010", prometheusService)
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
