package logger

import (
	"go.uber.org/zap"
	"grpc-admin/app/thirdparty/internal/conf"
	"grpc-admin/common/pkg"
)

var zapLogger *zap.SugaredLogger

func NewZapLogger() *zap.SugaredLogger {
	if zapLogger != nil {
		return zapLogger
	}

	zapLogger = pkg.NewZapLogger(pkg.ZapConfig{
		OutputDir: conf.AppConf.Logger.Zap.OutputDir,
		Format:    conf.AppConf.Logger.Zap.Format,
		Level:     conf.AppConf.Logger.Zap.Level,
	})

	return zapLogger
}
