package balances

import (
	"context"

	"github.com/naufalfmm/csv-transaction/models/dao"
)

func (r repositories) UpdateAmount(ctx context.Context, id uint64, total, totalFailed float64) error {
	if _, err := r.db.GetDB().
		NewUpdate().
		Model((*dao.Balance)(nil)).
		Set("total = total + ?", total).
		Set("failed_total = failed_total + ?", totalFailed).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		r.log.Error(ctx, LogMsgUpdateAmount).Err(err).Send()
		return err
	}

	return nil
}
