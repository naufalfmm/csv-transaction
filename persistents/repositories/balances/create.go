package balances

import (
	"context"

	"github.com/naufalfmm/csv-transaction/models/dao"
)

func (r repositories) Create(ctx context.Context, bal dao.Balance) (dao.Balance, error) {
	if _, err := r.db.GetDB().NewInsert().Model(&bal).Exec(ctx); err != nil {
		r.log.Error(ctx, LogMsgCreate).Err(err).Send()
		return dao.Balance{}, err
	}

	return bal, nil
}
