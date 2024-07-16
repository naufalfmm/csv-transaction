package consts

const (
	SuccessStatus = "SUCCESS"
	FailedStatus  = "FAILED"
	PendingStatus = "PENDING"

	DebitType  = "DEBIT"
	CreditType = "CREDIT"
)

var (
	CalFactor = map[string]float64{
		DebitType:  1,
		CreditType: -1,
	}
)
