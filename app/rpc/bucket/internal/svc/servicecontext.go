package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"log"
	"oos-system/app/rpc/bucket/internal/config"
	"oos-system/app/rpc/model/bucketmodel"
	"oos-system/app/rpc/model/objecthashmodel"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	BucketModel    bucketmodel.BucketModel
	ObjectHashMode objecthashmodel.ObjectHashModel
	Casbin         *casbin.SyncedEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	// 加载casbin服务
	casbinSqlConn, err := xormadapter.NewAdapter("mysql", c.Mysql.DataSource, true)
	casbinModel, err := model.NewModelFromString(`
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
	`)
	// 在生成的casbin权限表中 p_type 代表p策略/g角色 v0代表角色 v1代表桶名 v2代表权限
	casbin, _ := casbin.NewSyncedEnforcer(casbinModel, casbinSqlConn)
	if err != nil {
		log.Fatalf("error: adapter: %s", err)
	}
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),

		BucketModel:    bucketmodel.NewBucketModel(sqlConn, c.CacheRedis),
		ObjectHashMode: objecthashmodel.NewObjectHashModel(sqlConn, c.CacheRedis),
		Casbin:         casbin,
	}
}
