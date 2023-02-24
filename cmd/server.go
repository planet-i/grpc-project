package main

import (
	"fmt"
	"net"
	"net/http"

	"grpc-project/common"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	httpServer *http.ServeMux
}

func NewServer(
	grpcServer *grpc.Server,
	httpServer *http.ServeMux,
) *Server {
	return &Server{
		grpcServer: grpcServer,
		httpServer: httpServer,
	}
}

func (s *Server) Start() error {
	eg := errgroup.Group{}
	eg.Go(func() error {
		return s.RunGrpcServer(common.Config.Server.GRPC.Host, common.Config.Server.GRPC.Port)
	})
	eg.Go(func() error {
		return s.RunHttpServer(common.Config.Server.Http.Host, common.Config.Server.Http.Port)
	})

	return eg.Wait()
}

func (s *Server) RunGrpcServer(host string, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	fmt.Printf("server grpc listen: %s:%d\n", host, port)
	return s.grpcServer.Serve(lis)
}

func (s *Server) RunHttpServer(host string, port int) error {
	s.HttpRouter()
	fmt.Printf("server http listen: %s:%d\n", host, port)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), s.httpServer)
}

func (s *Server) HttpRouter() {
	s.httpServer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})
}
