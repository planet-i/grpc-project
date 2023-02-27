package service

import (
	"context"
	"fmt"

	"grpc-project/api"
)

func (i *InitService) Init(ctx context.Context, req *api.NoParam) (*api.BoolReply, error) {
	fmt.Println("start initService")
	return &api.BoolReply{
		Code:   200,
		Result: true,
	}, nil
}
