package repo

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"grpc-admin/app/thirdparty/internal/conf"
	"grpc-admin/app/thirdparty/internal/pkg/cache"
	"grpc-admin/app/thirdparty/internal/pkg/db"
	"grpc-admin/app/thirdparty/internal/pkg/logger"
)

type IThirdPartyRepo interface {
}

type ThirdPartyRepo struct {
	config conf.Conf
	db     *gorm.DB
	rds    *redis.Client
	logger *zap.SugaredLogger
}

func NewThirdPartyRepo() IThirdPartyRepo {
	return &ThirdPartyRepo{
		db:     db.NewGormDB(),
		rds:    cache.NewRedisCache(),
		logger: logger.NewZapLogger(),
	}
}
