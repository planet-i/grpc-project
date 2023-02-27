package service

import (
	"github.com/google/wire"

	"grpc-project/api"
)

var ProviderSet = wire.NewSet(
	NewInitService,
)

type InitService struct {
	api.UnimplementedInitServiceServer
}

func NewInitService() *InitService {
	return &InitService{}
}
