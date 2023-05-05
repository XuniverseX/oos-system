package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"oos-system/app/rpc/metadata/internal/config"
	"sync"
)

type ServiceContext struct {
	Config    config.Config
	RWLockMap *sync.Map
	Redis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	var lockMap sync.Map
	return &ServiceContext{
		Config:    c,
		RWLockMap: &lockMap,
		Redis:     c.Redis.NewRedis(),
	}
}
