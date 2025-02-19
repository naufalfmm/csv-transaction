package db

import (
	"context"
	"database/sql"

	"github.com/naufalfmm/csv-transaction/utils/orm"
)

type Trx struct {
	trx orm.Orm
}

func (t *Trx) Begin(ctx context.Context, o orm.Orm) {
	tr, _ := o.BeginTx(ctx, &sql.TxOptions{})
	t.trx = tr
}

func (t *Trx) Commit() {
	t.trx.Commit() //nolint:errcheck
}

func (t *Trx) Rollback() {
	t.trx.Rollback() //nolint:errcheck
}
