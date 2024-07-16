package log

import (
	"github.com/naufalfmm/csv-transaction/resources/config"
	"github.com/naufalfmm/csv-transaction/utils/logger"
	"github.com/naufalfmm/csv-transaction/utils/logger/zeroLogger"
)

func NewLogger(c *config.EnvConfig) (logger.Logger, error) {
	return zeroLogger.NewZeroLogger(
		zeroLogger.WithEnabled(c.LogMode),
	)
}
