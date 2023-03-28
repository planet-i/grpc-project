package handler

import (
	"grpc-project/api/simple/handler/tag"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	tag.ProviderSet,
)
