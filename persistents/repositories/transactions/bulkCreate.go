package transactions

import (
	"context"

	"github.com/naufalfmm/csv-transaction/models/dao"
)

func (r repositories) BulkCreate(ctx context.Context, trxs []dao.Transaction) ([]dao.Transaction, error) {
	if len(trxs) == 0 {
		return nil, nil
	}

	if _, err := r.db.GetDB().NewInsert().Model(&trxs).Exec(ctx); err != nil {
		r.log.Error(ctx, LogMsgCreate).Err(err).Send()
		return nil, err
	}

	return trxs, nil
}
