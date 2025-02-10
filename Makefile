.PHONY: requester

requester:
	@cd requester && go run cmd/cli.go

test:
	@cd requester && go test -v ./...