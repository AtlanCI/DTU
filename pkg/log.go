package pkg

import (
	"go.uber.org/zap"
)

var LogSugar zap.Logger

func init() {
	LogSugar, _ := zap.NewProduction()
	defer LogSugar.Sync()
}
