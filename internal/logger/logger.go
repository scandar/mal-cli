package logger

import (
	"go.uber.org/zap"
)

type AppLogger = *zap.SugaredLogger

var Instance AppLogger

func InitLogger(isDev bool) {
	if isDev {
		Instance = zap.Must(zap.NewDevelopment()).Sugar()
	} else {
		config := zap.NewProductionConfig()
		config.Encoding = "console"
		Instance = zap.Must(config.Build()).Sugar()
	}
	defer Instance.Sync()
}
