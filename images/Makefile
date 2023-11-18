build:
	cd cmd && go build -o images .

dev:
	cd cmd && air .

test:
	go test ./... -coverprofile=cover.out

test-cover:
	go tool cover -html=cover.out

lint:
	golangci-lint run

.PHONY: proto
.PHONY: build
.PHONY: dev
