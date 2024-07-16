package mysqlOrm

import (
	"fmt"
	"time"

	"github.com/naufalfmm/csv-transaction/utils/logger"
	"github.com/naufalfmm/csv-transaction/utils/logger/zeroLogger"
)

type mysqlConfig struct {
	host     string
	port     string
	username string
	password string
	dbname   string

	logger           logger.Logger
	logMode          bool
	logSlowThreshold time.Duration

	retry     int
	waitSleep time.Duration
}

func generateDefault() (mysqlConfig, error) {
	log, err := zeroLogger.NewZeroLogger()
	if err != nil {
		return mysqlConfig{}, err
	}

	return mysqlConfig{
		host:     "localhost",
		port:     "5432",
		username: "root",

		logger:           log,
		logMode:          false,
		logSlowThreshold: 200 * time.Millisecond,
	}, nil
}

func (c mysqlConfig) Addr() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.username, c.password, c.host, c.port, c.dbname)
}

type MysqlConfig func(c *mysqlConfig)

func WithHostPort(host, port string) MysqlConfig {
	return func(c *mysqlConfig) {
		c.host = host
		c.port = port
	}
}

func WithUsernamePassword(username, password string) MysqlConfig {
	return func(c *mysqlConfig) {
		c.username = username
		c.password = password
	}
}

func WithDatabaseName(dbname string) MysqlConfig {
	return func(c *mysqlConfig) {
		c.dbname = dbname
	}
}

func WithLogger(logger logger.Logger) MysqlConfig {
	return func(c *mysqlConfig) {
		c.logger = logger
		c.logMode = true
	}
}

func WithLog(logger logger.Logger, slowThreshold time.Duration) MysqlConfig {
	return func(c *mysqlConfig) {
		c.logger = logger
		c.logSlowThreshold = slowThreshold
		c.logMode = true
	}
}

func WithRetry(retry int, waitSleep time.Duration) MysqlConfig {
	return func(c *mysqlConfig) {
		c.retry = retry
		c.waitSleep = waitSleep
	}
}
