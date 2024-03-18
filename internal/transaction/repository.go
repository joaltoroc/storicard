package transaction

import (
	"context"
	"github/joaltoroc/storicard/internal/transaction/entities"
)

type Repository interface {
	InsertData(ctx context.Context, transactions []entities.Transaction) error
	GetData(ctx context.Context) ([]entities.Transaction, error)
	GetDataByID(ctx context.Context, executionID string) ([]entities.Transaction, error)
}
