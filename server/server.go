package server

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewGRPC, NewGateway, NewHttp)
