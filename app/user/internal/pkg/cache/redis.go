package cache

import (
	"github.com/go-redis/redis/v8"
	"grpc-admin/app/user/internal/conf"
	"grpc-admin/common/pkg"
)

var goRedisCli *redis.Client

func NewRedisCache() *redis.Client {
	if goRedisCli != nil {
		return goRedisCli
	}

	goRedisCli = pkg.NewGoRedis(pkg.GoRedisConfig{
		Host:     conf.AppConf.Cache.Redis.Host,
		Password: conf.AppConf.Cache.Redis.Password,
		DB:       conf.AppConf.Cache.Redis.DB,
	})

	return goRedisCli
}
