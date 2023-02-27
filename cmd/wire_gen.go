// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"grpc-project/server"
	"grpc-project/service"
)

// Injectors from wire.go:

func initServer() (*Server, error) {
	grpcServer := server.NewGRPC()
	serveMux := server.NewGateway()
	httpServeMux := server.NewHttp()
	initService := service.NewInitService()
	mainServer := NewServer(grpcServer, serveMux, httpServeMux, initService)
	return mainServer, nil
}
