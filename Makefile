GOLANGCI_LINT_VERSION := v2.15.0

.PHONY: build test lint fmt

build:
	go build ./...

test:
	go test ./...

lint:
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run ./...

fmt:
	gofmt -w .
