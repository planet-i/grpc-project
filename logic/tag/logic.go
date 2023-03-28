package tag

import (
	"grpc-project/service/tag"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewLogic,
)

type Logic struct {
	service *tag.Service
}

func NewLogic(service *tag.Service) *Logic {
	return &Logic{
		service: service,
	}
}
