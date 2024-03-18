package entities

import "time"

type (
	Transaction struct {
		ID              int64 `gorm:"primaryKey;autoIncrement:true"`
		ExecutionID     string
		FileID          int64
		TypeTransaction TypeTransaction
		Date            time.Time
		Value           float64
		CreatedAt       time.Time
	}

	TypeTransaction string
)

const (
	DebitType  TypeTransaction = "debit"
	CreditType TypeTransaction = "credit"
)
