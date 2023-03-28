package logic

import (
	"grpc-project/logic/tag"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	tag.ProviderSet,
)
