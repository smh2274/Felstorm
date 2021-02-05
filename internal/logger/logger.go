package logger

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var zlg *zap.Logger

// InitLogger 初始化Logger
func InitLogger(v *viper.Viper) (err error) {
	fileName := v.GetString("log.file")
	maxSize := v.GetInt("log.maxSize")
	maxAge := v.GetInt("log.maxAge")
	maxBackups := v.GetInt("log.maxBackups")
	level := v.GetString("log.level")

	writeSyncer := getLogWriter(fileName, maxSize, maxBackups, maxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	zlg = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
	zap.ReplaceGlobals(zlg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return
}

// getEncoder 设置日志参数
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// getLogWriter 获取日志流
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	hook := &lumberjack.Logger{
		Filename:   filename,  // 日志文件路径
		MaxSize:    maxSize,   // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackup, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,    // 文件最多保存多少天
		Compress:   true,      // 是否压缩
	}
	return zapcore.AddSync(hook)
}

func RecordErr(req interface{}, err error, desc ...string) {
	var info string
	if len(desc) > 0 {
		info = desc[0] + ": "
	}
	zap.L().Error(fmt.Sprintf("%s request: %#v, err: %v",info, req, err.Error()))
}

