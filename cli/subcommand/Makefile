COMMIT_HASH=$(shell git rev-parse --short HEAD)
GOBIN=$(CURDIR)/bin

dev-tools:
	go install github.com/goreleaser/goreleaser@latest

build:
	go build -o ./bin/example -ldflags '-X main.CommitHash=$(COMMIT_HASH)' ./cmd/example
