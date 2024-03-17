package usecase

import (
	"context"
	"net/http"

	"github/joaltoroc/storicard/internal/transaction"
	"github/joaltoroc/storicard/internal/transaction/dtos"
)

type (
	usecase struct {
		repo transaction.Repository
	}
)

func NewUseCase(repo transaction.Repository) transaction.UseCase {
	return &usecase{repo}
}

// ProcessFile implements transaction.UseCase.
func (u *usecase) ProcessFile(ctx context.Context, payload dtos.Payload) (httpCode int, err error) {
	return http.StatusOK, nil
}
