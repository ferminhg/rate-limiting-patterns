.PHONY: requester leaky-bucket

requester:
ifdef n
	@cd requester && go run cmd/cli.go -n ${n}
else
	@cd requester && go run cmd/cli.go
endif

test:
	@cd requester && go test -v ./...

leaky-bucket:
	@cd leaky-bucket && go run cmd/server.go
