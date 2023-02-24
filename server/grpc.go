package server

import "google.golang.org/grpc"

func NewGRPC() *grpc.Server {
	server := grpc.NewServer()
	return server
}
