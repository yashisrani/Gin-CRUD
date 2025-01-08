package utils

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yashisrani/Go-Gin-API/model"
)

var Logger logrus.Logger

func init() {
	Logger = *logrus.New()
	Logger.Out = os.Stdout
}

func SetLogger() {
	loglevel := flag.String(model.LogLevel, model.LogLevelInfo, "log-level (debug,error,info,warning)")

	flag.Parse()
	switch *loglevel {
	case model.LogLevelDebug:
		Logger.SetLevel(logrus.DebugLevel)

	case model.LogLevelError:
		Logger.SetLevel(logrus.ErrorLevel)

	case model.LogLevelWarn:
		Logger.SetLevel(logrus.WarnLevel)

	case model.LogLevelFatal:
		Logger.SetLevel(logrus.FatalLevel)

	case model.LogLevelPanic:
		Logger.SetLevel(logrus.PanicLevel)

	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
}

