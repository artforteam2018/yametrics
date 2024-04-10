package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var Log *zap.Logger

func Init() {
	log, err := zap.NewProductionConfig().Build()

	if err != nil {
		fmt.Println("Error zap initialization", err)
		panic("zap not initialized")
	}
	defer log.Sync()

	Log = log
}
