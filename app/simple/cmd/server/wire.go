//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"grpc-project/api/simple/handler"
	"grpc-project/logic"
	"grpc-project/repo"
	"grpc-project/server"
	"grpc-project/service"
)

func initServer() (*Server, error) {
	panic(wire.Build(
		repo.ProviderSet,    // 数据仓库，与数据库相关联，以及一些第三方服务相关联
		service.ProviderSet, // 服务单一
		logic.ProviderSet,   // 处理业务逻辑，封装报错
		handler.ProviderSet, // 转换数据，将client传入的数据转换成logic需要的数据，将logic返回的数据转换成client需要的数据
		server.ProviderSet,
		NewServer,
	))

	return &Server{}, nil
}
