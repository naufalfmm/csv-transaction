package transactions

import (
	"context"
	"encoding/csv"
	"io"
	"os"
	"time"

	"github.com/naufalfmm/csv-transaction/models/dao"
	"github.com/naufalfmm/csv-transaction/models/dto"
)

func (u usecases) Record(ctx context.Context, req dto.RecordTransactionRequest) error {
	if req.Filename == "" {
		return nil
	}

	file, err := os.Open(req.Filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bal, err := u.persistents.Repositories.Balances.Create(ctx, dao.Balance{
		Total:       0,
		FailedTotal: 0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return err
	}

	trxs := []dao.Transaction{}

	csvReader := csv.NewReader(file)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		trx := dao.NewTransactionCsv(rec)
		trx.BalanceID = bal.ID

		bal.Total += trx.GetAmount()
		bal.FailedTotal += trx.GetFailedAmount()

		trxs = append(trxs, trx)
	}

	if _, err := u.persistents.Repositories.Transactions.BulkCreate(ctx, trxs); err != nil {
		return err
	}

	return u.persistents.Repositories.Balances.UpdateAmount(ctx, bal.ID, bal.Total, bal.FailedTotal)
}
