package transaction

import (
	"context"

	"github/joaltoroc/storicard/internal/transaction/dtos"
	"github/joaltoroc/storicard/internal/transaction/entities"
)

type UseCase interface {
	ProcessFile(ctx context.Context, payload dtos.Payload) (int, string, error)
	GetData(ctx context.Context) ([]entities.Transaction, error)
	GetDataByID(ctx context.Context, executionID string) ([]entities.Transaction, error)
}
