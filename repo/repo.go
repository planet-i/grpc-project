package repo

import (
	"grpc-project/repo/tag"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewDB,
	NewRedis,

	tag.ProviderSet,
)
