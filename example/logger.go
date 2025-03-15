package example

import (
	"github.com/rafae1f/go-common/logger"
)

func Logger() {
	logger.Initialize("production")
	logger.Info("This info log")
}
