//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"grpc-project/server"
)

func initServer() (*Server, error) {
	panic(wire.Build(
		server.ProviderSet,
		NewServer,
	))

	return &Server{}, nil
}
