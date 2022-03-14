package reporters

import (
	"graph-svc/pkg/config"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const permissions os.FileMode = 0744

func NewExternalLogFile(cfg config.LogFileConfig) *lumberjack.Logger {
	if err := os.MkdirAll(cfg.GetFileDir(), permissions); err != nil {
		return nil
	}

	return &lumberjack.Logger{}
}
