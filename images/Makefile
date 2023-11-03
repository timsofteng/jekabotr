build:
	go build .

dev:
	air .

test:
	go test ./... -coverprofile=cover.out

test-cover:
	go tool cover -html=cover.out

lint:
	golangci-lint run

.PHONY: proto
