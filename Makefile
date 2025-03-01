.PHONY: requester rate-limiter-service

requester:
ifdef n
	@cd requester && go run cmd/cli.go -n ${n}
else
	@cd requester && go run cmd/cli.go
endif

test:
	@cd requester && go test -v ./...

rate-limiter-service:
	@cd rate-limiter-service && go run cmd/server.go
