package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// ตั่งค่า zaplog เปลี่ยนการตั้งค่าต่างๆ
func init() {
	var err error

	config := zap.NewProductionConfig()

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""
	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

// สร้าง custom log ประเภทต่างๆ
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	log.Warn(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	default:
	}
}
