package server

import (
	"context"
	"encoding/json"
	"errors"
	"grpc-project/common"
	"grpc-project/models"
	"grpc-project/utils"
	"time"

	"github.com/go-redis/redis/v8"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

// 定义拦截器结构体
type MyInterceptor struct {
	redisClient *redis.Client
}

// 定义拦截器构造函数
func NewMyInterceptor(redisClient *redis.Client) *MyInterceptor {
	return &MyInterceptor{redisClient: redisClient}
}

// 实现拦截器接口
func (i *MyInterceptor) TokenAuthIntercept(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 从 Authorization 请求中获取 Bearer token
	token, err := grpcAuth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return ctx, err
	}
	// 从token中获取用户ID
	userID, err := i.GetUserIDFromToken(ctx, token)
	if err != nil {
		return ctx, err
	}
	// 通过用户ID判断token是否有效
	if err := i.IsUsefulToken(ctx, userID, token); err != nil {
		return ctx, err
	}
	// 刷新token的过期时间
	if err := i.RefreshTokenTime(ctx, token); err != nil {
		return ctx, err
	}
	// 将UserID存入ctx
	newCtx := context.WithValue(ctx, common.UserIDWithCtx, userID)
	// 调用下一个处理器
	return handler(newCtx, req)
}

// GetUserIDFromToken 用token获取UserID
func (i *MyInterceptor) GetUserIDFromToken(ctx context.Context, token string) (string, error) {
	// 从redis中获取token为key的值
	value, err := i.redisClient.Get(ctx, token).Result()
	if err == redis.Nil {
		return "", errors.New("登录超时，请重新登录")
	} else if err != nil {
		return "", errors.New("系统内部错误")
	}

	var tokenValue models.TokenValue
	if err := json.Unmarshal([]byte(value), &tokenValue); err != nil {
		return "", errors.New("系统内部错误")
	}
	return tokenValue.UserID, nil
}

// IsUserfulToken 用UserID判断token是否有效
func (i *MyInterceptor) IsUsefulToken(ctx context.Context, userID, token string) error {
	value, err := i.redisClient.Get(ctx, userID).Result()
	if err != nil {
		return errors.New("系统内部错误")
	}
	var userIDValue models.UserIDValue
	if err := json.Unmarshal([]byte(value), &userIDValue); err != nil {
		return errors.New("获取用户信息失败")
	}
	isExist := utils.IsStringInSlice(token, userIDValue.UsefulToken)
	if !isExist {
		return errors.New("token无效")
	}
	return nil
}

// RefreshTokenTime 刷新token过期时间
func (i *MyInterceptor) RefreshTokenTime(ctx context.Context, token string) error {
	expireTime := 2 * time.Hour
	_, err := i.redisClient.Expire(ctx, token, expireTime).Result()
	if err != nil {
		return err
	}
	return nil
}
