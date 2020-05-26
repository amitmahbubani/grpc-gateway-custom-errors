package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/amitmahbubani/grpc-gateway-custom-errors/errors"
	userpb "github.com/amitmahbubani/grpc-gateway-custom-errors/proto_generated/proto/user"
)

const serverAddress = "127.0.0.1:8050"

func main() {
	// Create a TCP listener
	tcp, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("failed to listen on tcp: %v", err)
	}

	s := grpc.NewServer()

	userServer := &UserServer{}

	userpb.RegisterUserServiceServer(s, userServer)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Printf("starting grpc server on: %s", serverAddress)

	if err := s.Serve(tcp); err != nil {
		log.Fatalf("failed serving grpc: %v", err)
	}
}

type UserServer struct{}

func (u *UserServer) Get(ctx context.Context, request *userpb.UserGetRequest) (*userpb.UserResponse, error) {
	// To illustrate an error response, we're failing if len(id)>5
	if len(request.GetId()) > 5 {
		err := &errors.AppError{
			Code:    "validation_failure",
			Message: "id should be less than 5 chars",
			Field:   "id",
		}
		return nil, toGrpcError(err)
	}

	resp := userpb.UserResponse{
		Id:   request.GetId(),
		Name: "Random Name",
		Age:  30,
	}

	return &resp, nil
}

func (u *UserServer) Create(ctx context.Context, request *userpb.UserCreateRequest) (*userpb.UserResponse, error) {
	// To illustrate an error response, we're failing if age < 0
	if request.Age < 0 {
		err := &errors.AppError{
			Code:    "validation_failure",
			Message: "validation error: age must be a positive integer",
			Field:   "age",
		}
		return nil, toGrpcError(err)
	}

	resp := userpb.UserResponse{
		Id:   "12345",
		Name: request.GetName(),
		Age:  request.GetAge(),
	}

	return &resp, nil
}

// toGrpcError converts an error to grpc.Status compatible
func toGrpcError(err error) error {
	ierr, ok := err.(errors.IError)
	if !ok {
		// If err does not implement errors.IError, return a default Status
		return status.Error(codes.Unknown, err.Error())
	}

	protoErr, ok := ierr.(errors.ProtoConstructable)
	if !ok {
		// If not of errors.ProtoConstructable, we could map err.GetCode() to a
		// grpc Code, etc.
		return status.Error(codes.Internal, err.Error())
	}

	grpcStatus := status.New(codes.Internal, err.Error())
	richErr, detailErr := grpcStatus.WithDetails(protoErr.ToErrorProto())
	if detailErr != nil {
		return grpcStatus.Err()
	}

	return richErr.Err()
}
