package transactions

import (
	"context"

	"github.com/naufalfmm/csv-transaction/models/dto"
	"github.com/naufalfmm/csv-transaction/persistents"
	"github.com/naufalfmm/csv-transaction/resources/db"
)

//go:generate mockgen -package=transactions -destination=../../mocks/usecases/transactions/init.go -source=init.go
type (
	Usecases interface {
		Record(ctx context.Context, req dto.RecordTransactionRequest) error
	}

	usecases struct {
		persistents persistents.Persistents

		d *db.DB
	}
)

func Init(persist persistents.Persistents, d *db.DB) (Usecases, error) {
	return &usecases{
		persistents: persist,
		d:           d,
	}, nil
}
