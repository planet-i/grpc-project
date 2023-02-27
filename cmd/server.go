package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"

	"grpc-project/api"
	"grpc-project/common"
	"grpc-project/service"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer  *grpc.Server
	grpcGateway *runtime.ServeMux
	httpServer  *http.ServeMux
	init        *service.InitService
}

func NewServer(
	grpcServer *grpc.Server,
	grpcGateway *runtime.ServeMux,
	httpServer *http.ServeMux,
	init *service.InitService,
) *Server {
	return &Server{
		grpcServer:  grpcServer,
		grpcGateway: grpcGateway,
		httpServer:  httpServer,
		init:        init,
	}
}

type registerServiceHandlerFromEndpoint func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

func (s *Server) Start() error {
	// Register grpc server
	api.RegisterInitServiceServer(s.grpcServer, s.init)

	endpoints := []registerServiceHandlerFromEndpoint{
		api.RegisterInitServiceHandlerFromEndpoint,
	}

	for _, endpoint := range endpoints {
		err := endpoint(
			context.Background(),
			s.grpcGateway,
			fmt.Sprintf("%s:%d", common.Config.Server.GRPC.Host, common.Config.Server.GRPC.Port),
			[]grpc.DialOption{
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			},
		)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}

	eg := errgroup.Group{}
	// 启动Grpc服务
	eg.Go(func() error {
		return s.RunGrpcServer(common.Config.Server.GRPC.Host, common.Config.Server.GRPC.Port)
	})
	// 启动HTTP服务
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
		s.grpcGateway.ServeHTTP(w, r)
	})
}
