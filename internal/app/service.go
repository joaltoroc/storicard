package app

import (
	handlerV1 "github/joaltoroc/storicard/internal/transaction/handler/v1"
	"github/joaltoroc/storicard/internal/transaction/repository"
	"github/joaltoroc/storicard/internal/transaction/storage"
	"github/joaltoroc/storicard/internal/transaction/usecase"
)

func (app *App) startService() error {
	storageS3 := storage.NewStorage(app.s3, app.cfg)
	repo := repository.NewRepository(app.db)
	useCase := usecase.NewUseCase(repo, storageS3)
	handler := handlerV1.NewHandlers(useCase)

	domain := app.echo.Group("/api/v1/transaction")
	PingPong(domain)

	handler.TransactionRoutes(domain)

	return nil
}
