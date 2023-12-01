package server

import (
	"github.com/go-redis/redis/v8"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewGRPC(redisClient *redis.Client) *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			//grpcAuth.UnaryServerInterceptor(tokenAuth),
			NewMyInterceptor(redisClient).TokenAuthIntercept,
			grpcRecovery.UnaryServerInterceptor(RecoveryInterceptor()),
		)),
	)
	return server
}

// RecoveryInterceptor panic时返回Unknown错误吗
func RecoveryInterceptor() grpcRecovery.Option {
	return grpcRecovery.WithRecoveryHandler(func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	})
}
