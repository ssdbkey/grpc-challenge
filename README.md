# gRPC Server and Client Challenge

## Requirements

- Go `1.17+`
- Go Modules

## Usage

- Clone this repository and go to folder:

```console
git clone https://github.com/ssdbkey/grpc-challenge.git
cd grpc-challenge
```

- Try to build binary:

```console
make build
```

- No errors? Let's run:

```console
make run
```

- Go to another console session and connect to gRPC server with [Evans](https://github.com/ktr0731/evans) (gRPC client):

```console
evans api/proto/tendermint.proto
```

- Try to run test:

```console
make test
```

- Try to run state tracker to get latest consecutive 5 blocks info(result is saved in exported.json):

```console
make run-tracker
```

- That's all!
