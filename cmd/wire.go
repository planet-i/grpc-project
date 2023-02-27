//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"grpc-project/server"
	"grpc-project/service"
)

func initServer() (*Server, error) {
	panic(wire.Build(
		service.ProviderSet,
		server.ProviderSet,
		NewServer,
	))

	return &Server{}, nil
}
