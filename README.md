# grpc-gateway-custom

This repo contains sample code that demonstrates some of the customization capabilities in [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)

## How To Run

**Install deps:**
1. Install protobuf
2. Install grpc binaries:
```shell script
go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go
```

**Run:**
1. `go run server/app/main.go`
2. `go run server/gateway/main.go`
3. `go run client/main.go`


## Examples

todo