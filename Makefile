.PHONY: build

build:
	@go build -o ./build/grpc_server ./cmd/server/...

build-tracker:
	@go build -o ./build/grpc_tracker ./cmd/tracker/...

.PHONY: run

test:
	sh test.sh

run: build
	@./build/grpc_server

run-tracker: build-tracker
	@./build/grpc_tracker