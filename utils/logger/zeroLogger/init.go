package zeroLogger

import (
	"os"

	"github.com/naufalfmm/csv-transaction/utils/logger"
	"github.com/rs/zerolog"
)

func NewZeroLogger(configs ...LoggerConfig) (logger.Logger, error) {
	zerolog.TimestampFieldName = "timestamp"
	zerolog.LevelFieldName = "level"
	zerolog.ErrorFieldName = "error"

	conf := config{}
	for _, confi := range configs {
		confi(&conf)
	}

	return &zelog{
		zl: zerolog.New(os.Stdout),

		conf: conf,
	}, nil
}
