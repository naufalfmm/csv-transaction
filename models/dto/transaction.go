package dto

import "strconv"

type TransactionRequest struct {
	Code   string
	Name   string
	Type   string
	Amount float64
	Status string
	Notes  string
}

func NewTransactionRequest(recs []string) TransactionRequest {
	amount, _ := strconv.ParseFloat(recs[3], 64)

	return TransactionRequest{
		Code:   recs[0],
		Name:   recs[1],
		Type:   recs[2],
		Amount: amount,
		Status: recs[4],
		Notes:  recs[5],
	}
}

type RecordTransactionRequest struct {
	Filename string
}
