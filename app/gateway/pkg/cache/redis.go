package cache

import (
	"github.com/go-redis/redis/v8"
	"grpc-admin/app/gateway/conf"
	"grpc-admin/common/pkg"
)

var cli *redis.Client

func NewRedisCache() *redis.Client {
	if cli != nil {
		return cli
	}

	cli = pkg.NewGoRedis(pkg.GoRedisConfig{
		Host:     conf.AppConf.Cache.Redis.Host,
		Password: conf.AppConf.Cache.Redis.Password,
		DB:       conf.AppConf.Cache.Redis.DB,
	})

	return cli
}
