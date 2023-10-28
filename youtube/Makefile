proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/youtube.proto

build:
	go build .

dev:
	air .

.PHONY: proto
