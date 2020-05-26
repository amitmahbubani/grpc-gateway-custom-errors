package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	userpb "github.com/amitmahbubani/grpc-gateway-custom/proto_generated/proto/user"
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
		err := status.Error(codes.InvalidArgument, "validation error: id should be less than 5 chars")
		return nil, err
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
		err := status.Error(codes.InvalidArgument, "validation error: age must be a positive integer")
		return nil, err
	}

	resp := userpb.UserResponse{
		Id:   "12345",
		Name: request.GetName(),
		Age:  request.GetAge(),
	}

	return &resp, nil
}
