package app

import (
	handlerV1 "github/joaltoroc/storicard/internal/transaction/handler/v1"
	"github/joaltoroc/storicard/internal/transaction/repository"
	"github/joaltoroc/storicard/internal/transaction/usecase"
)

func (app *App) startService() error {
	repo := repository.NewRepository(app.db)
	useCase := usecase.NewUseCase(repo)
	handler := handlerV1.NewHandlers(useCase)

	domain := app.echo.Group("/api/v1/transaction")
	PingPong(domain)

	handler.TransactionRoutes(domain)

	return nil
}
