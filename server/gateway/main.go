package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/amitmahbubani/grpc-gateway-custom-errors/errors"
	usergw "github.com/amitmahbubani/grpc-gateway-custom-errors/proto_generated/proto/user"
)

var grpcServerAddress = "127.0.0.1:8050"

var gatewayAddress = "127.0.0.1:8020"

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	runtime.GlobalHTTPErrorHandler = CustomErrorHandler
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

func CustomErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(status.Code(err)))

	s := status.Convert(err)
	for _, d := range s.Details() {
		switch info := d.(type) {
		case *errors.Error:
			jErr := json.NewEncoder(w).Encode(info)
			if jErr != nil {
				w.Write([]byte(fallback))
			}
			return
		}
	}

	runtime.DefaultHTTPError(ctx, mux, marshaler, w, nil, err)
}
