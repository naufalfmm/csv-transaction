package db

import (
	"github.com/naufalfmm/csv-transaction/resources/config"
	"github.com/naufalfmm/csv-transaction/utils/logger"
	"github.com/naufalfmm/csv-transaction/utils/orm/driver/mysqlOrm"
)

func NewMysql(c *config.EnvConfig, log logger.Logger) (*DB, error) {
	confs := []mysqlOrm.MysqlConfig{
		mysqlOrm.WithHostPort(c.DbHost, c.DbPort),
		mysqlOrm.WithUsernamePassword(c.DbUsername, c.DbPassword),
		mysqlOrm.WithDatabaseName(c.DbName),
		mysqlOrm.WithRetry(c.DbRetry, c.DbWaitSleep),
	}

	if c.DbLogMode {
		confs = append(confs, mysqlOrm.WithLog(log, c.DbLogSlowThreshold))
	}

	o, err := mysqlOrm.NewMysql(confs...)
	if err != nil {
		return nil, err
	}

	return &DB{
		O: o,
	}, nil
}
