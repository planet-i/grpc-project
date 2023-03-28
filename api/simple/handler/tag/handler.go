package tag

import (
	v1 "grpc-project/api/simple/v1"
	"grpc-project/logic/tag"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewHandler,
)

type Handler struct {
	v1.UnimplementedTagServiceServer
	logic *tag.Logic
}

func NewHandler(logic *tag.Logic) *Handler {
	return &Handler{
		logic: logic,
	}
}
