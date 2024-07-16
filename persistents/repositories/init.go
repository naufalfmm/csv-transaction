package repositories

import (
	"github.com/naufalfmm/csv-transaction/persistents/repositories/balances"
	"github.com/naufalfmm/csv-transaction/persistents/repositories/transactions"
	"github.com/naufalfmm/csv-transaction/resources/db"
	"github.com/naufalfmm/csv-transaction/utils/logger"
)

type Repositories struct {
	Balances     balances.Repositories
	Transactions transactions.Repositories
}

func Init(d *db.DB, l logger.Logger) (Repositories, error) {
	bal, err := balances.Init(d, l)
	if err != nil {
		return Repositories{}, err
	}

	trx, err := transactions.Init(d, l)
	if err != nil {
		return Repositories{}, err
	}

	return Repositories{
		Balances:     bal,
		Transactions: trx,
	}, nil
}
