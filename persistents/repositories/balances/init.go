package balances

import (
	"context"

	"github.com/naufalfmm/csv-transaction/models/dao"
	"github.com/naufalfmm/csv-transaction/resources/db"
	"github.com/naufalfmm/csv-transaction/utils/logger"
)

const (
	LogMsgCreate       = "create-balance"
	LogMsgUpdateAmount = "update-amount-balance"
)

//go:generate mockgen -package=balances -destination=../../../mocks/persistents/repositories/balances/init.go -source=init.go
type (
	Repositories interface {
		Create(ctx context.Context, bal dao.Balance) (dao.Balance, error)
		UpdateAmount(ctx context.Context, id uint64, total, totalFailed float64) error
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
