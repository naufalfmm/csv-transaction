package dao

import (
	"time"

	"github.com/uptrace/bun"
)

type Balance struct {
	bun.BaseModel `bun:"balances"`
	ID            uint64    `bun:"id,pk,autoincrement"`
	Total         float64   `bun:"total,notnull"`
	FailedTotal   float64   `bun:"failed_total,notnull"`
	CreatedAt     time.Time `bun:"created_at,notnull"`
	UpdatedAt     time.Time `bun:"updated_at,notnull"`
}

