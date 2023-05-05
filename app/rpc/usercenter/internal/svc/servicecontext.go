package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"oos-system/app/rpc/model/bucketmodel"
	"oos-system/app/rpc/model/usermodel"
	"oos-system/app/rpc/usercenter/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	UserModel   usermodel.UserModel
	BucketModel bucketmodel.BucketModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		UserModel:   usermodel.NewUserModel(sqlConn, c.CacheRedis),
		BucketModel: bucketmodel.NewBucketModel(sqlConn, c.CacheRedis),
	}
}
