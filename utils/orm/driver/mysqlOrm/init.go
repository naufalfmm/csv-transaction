package mysqlOrm

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/naufalfmm/csv-transaction/utils/orm"
	"github.com/naufalfmm/csv-transaction/utils/orm/driver/mysqlOrm/mysqlOrmLogger"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func NewMysql(configs ...MysqlConfig) (orm.Orm, error) {
	conf, err := generateDefault()
	if err != nil {
		return nil, err
	}

	for _, config := range configs {
		config(&conf)
	}

	sqldb, err := sql.Open("mysql", conf.Addr())
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, mysqldialect.New())

	for i := 0; i < conf.retry; i++ {
		if err := db.Ping(); err == nil {
			break
		}
	}

	if conf.logger != nil {
		db.AddQueryHook(mysqlOrmLogger.NewLogQueryHook(
			mysqlOrmLogger.WithLogger(conf.logger),
			mysqlOrmLogger.WithSlowThreshold(conf.logSlowThreshold),
		))
	}

	return &postgresOrm{
		db: db,
	}, nil
}
