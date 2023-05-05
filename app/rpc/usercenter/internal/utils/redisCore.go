package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"oos-system/app/rpc/usercenter/internal/config"
)

// RedisDb 全局定义
var (
	RedisDb *redis.Client
)

// NewRedisCore 创建 redis 链接
func NewRedisCore(c config.Config) {
	var ctx = context.Background()
	//fmt.Println("redisconfig", c.Redis.Host, c.Redis.Pass)
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass, // no password set
		DB:       0,            // use default DB
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		//连接失败
		println(err)
	}
}
