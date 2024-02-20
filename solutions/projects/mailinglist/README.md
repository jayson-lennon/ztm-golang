# Setup

This project requires a `gcc` compiler installed and the `protobuf` code generation tools.


## Install protobuf compiler

Download the appropriate binary for your operating system from [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases).

Install by placing the extracted binary somewhere in your `$PATH`. Additional instructions available at [https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/).


## Install Go protobuf codegen tools

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`


## Generate Go code from .proto files

After creating your `.proto` file, this command will generate the necessary code.

```
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  Proto/mail.proto
```

