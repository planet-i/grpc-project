package tag

import (
	"context"
	"fmt"
	v1 "grpc-project/api/simple/v1"
)

// CreateTag 创建标签
func (h *Handler) CreateTag(ctx context.Context, req *v1.NoParam) (*v1.BoolReply, error) {
	fmt.Println("handler CreateTag")
	return &v1.BoolReply{
		Code:   200,
		Result: true,
	}, nil
}

// DeleteTag 删除标签
func (h *Handler) DeleteTag(ctx context.Context, req *v1.NoParam) (*v1.BoolReply, error) {
	fmt.Println("handler DeleteTag")
	return &v1.BoolReply{
		Code:   200,
		Result: true,
	}, nil
}

// UpdateTag 编辑标签
func (h *Handler) UpdateTag(ctx context.Context, req *v1.IDReq) (*v1.BoolReply, error) {
	fmt.Println("handler UpdateTag")
	return &v1.BoolReply{
		Code:   200,
		Result: true,
	}, nil
}

// GetTag 获取标签详情
func (h *Handler) GetTag(ctx context.Context, req *v1.IDReq) (*v1.BoolReply, error) {
	fmt.Println("handler GetTag")
	return &v1.BoolReply{
		Code:   200,
		Result: true,
	}, nil
}

// ListTag 列举标签
func (h *Handler) ListTag(ctx context.Context, req *v1.NoParam) (*v1.BoolReply, error) {
	fmt.Println("handler ListTag")
	return &v1.BoolReply{
		Code:   200,
		Result: true,
	}, nil
}
