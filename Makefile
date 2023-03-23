.PHONY: build

build:
	@go build -o ./build/grpc_server ./cmd/server/...

.PHONY: run

test:
	@sh test.sh

run: build
	@./build/grpc_server