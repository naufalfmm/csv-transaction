package dao

import (
	"strconv"
	"time"

	"github.com/naufalfmm/csv-transaction/consts"
	"github.com/uptrace/bun"
)

type Transaction struct {
	bun.BaseModel `bun:"transactions"`
	ID            uint64    `bun:"id,pk,autoincrement"`
	BalanceID     uint64    `bun:"balance_id,notnull"`
	Code          string    `bun:"code,notnull"`
	Name          string    `bun:"name,notnull"`
	Type          string    `bun:"type,notnull"`
	Amount        float64   `bun:"amount,notnull"`
	Status        string    `bun:"status,notnull"`
	Notes         string    `bun:"notes,notnull"`
	CreatedAt     time.Time `bun:"created_at,notnull"`
	UpdatedAt     time.Time `bun:"updated_at,notnull"`
}

func NewTransactionCsv(recs []string) Transaction {
	amount, _ := strconv.ParseFloat(recs[3], 64)

	return Transaction{
		Code:      recs[0],
		Name:      recs[1],
		Type:      recs[2],
		Amount:    amount,
		Status:    recs[4],
		Notes:     recs[5],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (t Transaction) GetAmount() float64 {
	if t.Status != consts.SuccessStatus {
		return 0.
	}

	return consts.CalFactor[t.Type] * t.Amount
}

func (t Transaction) GetFailedAmount() float64 {
	if t.Status != consts.FailedStatus {
		return 0.
	}

	return consts.CalFactor[t.Type] * t.Amount
}
