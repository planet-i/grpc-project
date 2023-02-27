package server

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func NewGateway() *runtime.ServeMux {
	mux := runtime.NewServeMux()
	return mux
}
