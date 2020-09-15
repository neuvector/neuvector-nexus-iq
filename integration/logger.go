package integration

import (
	"github.com/neuvector/neuvector-nexus-iq/internal/logger"
)

var log logger.Logger

func init() {
	l, err := logger.NewLogger(logger.Configuration{
		Level: logger.Info,
	})
	if err != nil {
		panic(err)
	}
	log = l
}

func SetLogger(l logger.Logger) {
	log = l
}
