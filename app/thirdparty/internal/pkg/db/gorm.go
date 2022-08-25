package db

import (
	"gorm.io/gorm"
	"grpc-admin/app/thirdparty/internal/conf"
	"grpc-admin/common/pkg"
)

var gormDB *gorm.DB

func NewGormDB() *gorm.DB {
	if gormDB != nil {
		return gormDB
	}

	gormDB = pkg.NewGorm(pkg.GormConfig{
		Source:       conf.AppConf.Database.Source,
		User:         conf.AppConf.Database.User,
		Password:     conf.AppConf.Database.Password,
		Host:         conf.AppConf.Database.Host,
		DbName:       conf.AppConf.Database.DbName,
		Charset:      conf.AppConf.Database.Charset,
		ParseTime:    conf.AppConf.Database.ParseTime,
		Loc:          conf.AppConf.Database.Loc,
		MaxIdleConns: conf.AppConf.Database.MaxIdleConns,
		MaxOpenConns: conf.AppConf.Database.MaxOpenConns,
	})

	return gormDB
}
