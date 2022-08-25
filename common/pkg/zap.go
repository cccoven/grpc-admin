package pkg

import (
	"fmt"
	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"grpc-admin/common/util"
	"log"
	"os"
	"path/filepath"
	"time"
)

type ZapConfig struct {
	OutputDir string
	Format    string
	Level     string
}

func NewZapLogger(c ZapConfig) *zap.SugaredLogger {
	writeSyncer := getLogWriter(c.OutputDir)
	encoder := getEncoder(c.Format)

	// NewCore 需要三个配置:
	// Encoder: 编码器，规定如何写入日志，使用 zapcore 自带的 NewJSONEncoder()，并使用预先设置的ProductionEncoderConfig()
	// WriteSyncer: 指定将日志写到哪里去，使用 zapcore.AddSync() 将打开的文件传递进去
	// LogLevel: 哪种级别的日志将被写入
	core := zapcore.NewCore(encoder, writeSyncer, getLogLevel(c.Level))

	// zap.AddCaller() 将调用函数信息记录到日志中的功能
	logger := zap.New(core, zap.AddCaller())
	// 使用 Sugar(), 可以使用 Infof、Errorf 等格式化方法
	sugaredLogger := logger.Sugar()
	
	return sugaredLogger
}

func getLogLevel(level string) zapcore.Level {
	// 	DebugLevel 可以打印出 Info, Debug, Warn(Error)
	// 	InfoLevel 可以打印出 Info, Warn(Error)
	// 	WarnLevel 只能打印出 Warn
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "httperror":
		return zapcore.ErrorLevel
	default:
		return zapcore.DebugLevel
	}
}

// 自定义日志输出时间格式
func customEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 15:04:05"))
}

func getEncoder(format string) zapcore.Encoder {
	// 以 json 形式输出日志到文件中，使用 zap.NewProductionEncoderConfig 提供默认的配置
	// return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	// 以普通控制台输出的形式写入文件中
	// return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	encoderConfig := zap.NewProductionEncoderConfig()
	// 修改时间编码器
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 使用自定义日志时间格式
	encoderConfig.EncodeTime = customEncodeTime
	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	if format == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(path string) zapcore.WriteSyncer {
	if exist, _ := util.PathExists(path); !exist {
		// 没有日志存储目录则创建
		fmt.Printf("Create %v directory\n", path)
		_ = os.Mkdir(path, os.ModePerm)
	}

	// 使用 rotatelogs 做日志分割归档
	fileWriter, err := zaprotatelogs.New(
		// 日志文件名
		filepath.Join(path, "%Y-%m-%d.log"),
		// 保存 7 天内的日志
		zaprotatelogs.WithMaxAge(7*24*time.Hour),
		// 每小时分割一次日志
		zaprotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		log.Fatal(err)
	}

	return zapcore.AddSync(fileWriter)
}
