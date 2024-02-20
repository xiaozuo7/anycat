package zaputil

import (
	"anycat/global/consts"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func CreateZapUtil() *zap.SugaredLogger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(consts.TimeForMate))
	}
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.MillisDurationEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig) // 日志格式

	//写入器
	lumberJackLogger := &lumberjack.Logger{
		Filename:   consts.LogFileLocation, //日志文件的位置
		MaxSize:    128,                    //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 1,                      //保留旧文件的最大个数
		MaxAge:     28,                     //保留旧文件的最大天数
		Compress:   false,                  //是否压缩/归档旧文件
	}
	writer := zapcore.AddSync(lumberJackLogger)
	zapCore := zapcore.NewCore(encoder, writer, zap.InfoLevel)                         // 日志级别
	return zap.New(zapCore, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel)).Sugar() // 抛出堆栈级别
}
