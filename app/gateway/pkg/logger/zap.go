package logger

import (
	"go.uber.org/zap"
	"grpc-admin/app/gateway/conf"
	"grpc-admin/common/pkg"
)

var logger *zap.SugaredLogger

func NewZapLogger() *zap.SugaredLogger {
	if logger != nil {
		return logger
	}

	logger = pkg.NewZapLogger(pkg.ZapConfig{
		OutputDir: conf.AppConf.Logger.Zap.OutputDir,
		Format:    conf.AppConf.Logger.Zap.Format,
		Level:     conf.AppConf.Logger.Zap.Level,
	})

	return logger
}
