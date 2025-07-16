package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, feild ...zap.Field) {
	log.Info(message, feild...)
}

func Debug(message string, feild ...zap.Field) {
	log.Info(message, feild...)
}

func Error(message interface{}, feild ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), feild...)
	case string:
		log.Error(v, feild...)
	}

}
