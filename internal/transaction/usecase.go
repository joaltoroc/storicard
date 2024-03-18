package transaction

import (
	"context"

	"github/joaltoroc/storicard/internal/transaction/dtos"
	"github/joaltoroc/storicard/internal/transaction/entities"
)

type UseCase interface {
	ProcessFile(ctx context.Context, payload dtos.Payload) (httpCode int, err error)
	GetData(ctx context.Context) ([]entities.Transaction, error)
}
