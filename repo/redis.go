package repo

import (
	"context"

	"grpc-project/common"

	"github.com/go-redis/redis/v8"
)

func NewRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: common.Config.Data.Redis.Addr,
	})

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	return client, err
}
