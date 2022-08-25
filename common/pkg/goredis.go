package pkg

import (
	"github.com/go-redis/redis/v8"
)

type GoRedisConfig struct {
	Host        string
	Password    string
	DB          int
}

func NewGoRedis(c GoRedisConfig) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:        c.Host,
		Password:    c.Password,
		DB:          c.DB,
	})

	return cli
}
