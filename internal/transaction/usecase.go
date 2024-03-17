package transaction

import (
	"context"

	"github/joaltoroc/storicard/internal/transaction/dtos"
)

type UseCase interface {
	ProcessFile(ctx context.Context, payload dtos.Payload) (httpCode int, err error)
}
