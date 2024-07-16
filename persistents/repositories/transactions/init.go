package transactions

import (
	"context"

	"github.com/naufalfmm/csv-transaction/models/dao"
	"github.com/naufalfmm/csv-transaction/resources/db"
	"github.com/naufalfmm/csv-transaction/utils/logger"
)

const (
	LogMsgCreate = "create-transaction"
)

//go:generate mockgen -package=transactions -destination=../../../mocks/persistents/repositories/transactions/init.go -source=init.go
type (
	Repositories interface {
		BulkCreate(ctx context.Context, trxs []dao.Transaction) ([]dao.Transaction, error)
	}

	repositories struct {
		db  *db.DB
		log logger.Logger
	}
)

func Init(d *db.DB, l logger.Logger) (Repositories, error) {
	return &repositories{
		db:  d,
		log: l,
	}, nil
}
