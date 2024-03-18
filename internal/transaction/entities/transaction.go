package entities

import "time"

type (
	Transaction struct {
		ID              int64 `gorm:"primaryKey;autoIncrement:true"`
		ExecutionID     string
		FileID          int64
		TypeTransaction TransactionType
		Date            time.Time
		Value           float64
		CreatedAt       time.Time
	}

	TransactionType string
)

const (
	DebitType  TransactionType = "debit"
	CreditType TransactionType = "credit"
)
