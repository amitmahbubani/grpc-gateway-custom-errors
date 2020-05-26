package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	usergw "github.com/amitmahbubani/grpc-gateway-custom/proto_generated/proto/user"
)

var grpcServerAddress = "127.0.0.1:8050"

var gatewayAddress = "127.0.0.1:8020"

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := usergw.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcServerAddress, opts)
	if err != nil {
		log.Fatalf("failed to regsiter user gateway: %v", err)
	}

	log.Printf("starting grpc server on: %s", gatewayAddress)

	err = http.ListenAndServe(gatewayAddress, mux)
	if err != nil {
		log.Fatalf("failed to start grpc-gateway server: %v", err)
	}
}
