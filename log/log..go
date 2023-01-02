package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

type Log struct {
	Logger *zap.Logger
}

func (l *Log) initLogger() {
	env := os.Getenv("ENV")
	currentTime := time.Now().UTC().Format("2006-01-02")
	logFilename := currentTime + "-" + "logs" + ".log"

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	if strings.EqualFold(env, "Prod") || strings.EqualFold(env, "Production") || strings.EqualFold(env, "Staging") {
		encConfig := zapcore.EncoderConfig{
			MessageKey: "msg",
			TimeKey:    "time",
			EncodeTime: zapcore.RFC3339TimeEncoder,
		}

		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   userHomeDir + "/log/fibac/" + logFilename,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
		})
		core := zapcore.NewCore(zapcore.NewJSONEncoder(encConfig), w, zap.InfoLevel)
		l.Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	} else {
		l.Logger, err = zap.NewDevelopment(zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
		if err != nil {
			panic(err)
		}
	}
}

func (l *Log) Get() *zap.Logger {
	if l.Logger == nil {
		l.initLogger()
	}
	return l.Logger
}
