package service

import (
	"grpc-project/service/tag"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	tag.ProviderSet,
)
