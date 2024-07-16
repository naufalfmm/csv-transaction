package usecases

import (
	"github.com/naufalfmm/csv-transaction/persistents"
	"github.com/naufalfmm/csv-transaction/resources/db"
	"github.com/naufalfmm/csv-transaction/usecases/transactions"
)

type Usecases struct {
	Transactions transactions.Usecases
}

func Init(persists persistents.Persistents, d *db.DB) (Usecases, error) {
	tr, err := transactions.Init(persists, d)
	if err != nil {
		return Usecases{}, err
	}

	return Usecases{
		Transactions: tr,
	}, nil
}
