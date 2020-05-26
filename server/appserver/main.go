package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

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
	resp := userpb.UserResponse{
		Id:   request.GetId(),
		Name: "Random Name",
		Age:  30,
	}

	return &resp, nil
}

func (u *UserServer) Create(ctx context.Context, request *userpb.UserCreateRequest) (*userpb.UserResponse, error) {
	resp := userpb.UserResponse{
		Id:   "12345",
		Name: request.GetName(),
		Age:  request.GetAge(),
	}

	return &resp, nil
}
