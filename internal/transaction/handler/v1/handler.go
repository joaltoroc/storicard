package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github/joaltoroc/storicard/internal/transaction"
	"github/joaltoroc/storicard/internal/transaction/dtos"
	"github/joaltoroc/storicard/pkg/utils/response"
)

type handlers struct {
	useCase transaction.UseCase
}

func NewHandlers(useCase transaction.UseCase) *handlers {
	return &handlers{useCase}
}

func (handler *handlers) ProcessFile(c echo.Context) error {
	var (
		ctx, cancel = context.WithTimeout(c.Request().Context(), 30*time.Second)
		payload     dtos.Payload
	)
	defer cancel()

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error()),
		)
	}

	if err := payload.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponseError(
			http.StatusBadRequest,
			response.MsgFailed,
			err.Error()),
		)
	}

	httpCode, executionID, err := handler.useCase.ProcessFile(ctx, payload)
	if err != nil {
		return c.JSON(httpCode, response.NewResponseError(
			httpCode,
			response.MsgFailed,
			err.Error()),
		)
	}

	return c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, response.MsgSuccess, map[string]string{"executionID": executionID}))
}

func (handler *handlers) GetData(c echo.Context) error {
	var (
		ctx, cancel = context.WithTimeout(c.Request().Context(), 30*time.Second)
	)
	defer cancel()

	data, err := handler.useCase.GetData(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewResponseError(
			http.StatusInternalServerError,
			response.MsgFailed,
			err.Error()),
		)
	}

	return c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, response.MsgSuccess, data))
}

func (handler *handlers) GetDataByID(c echo.Context) error {
	var (
		ctx, cancel = context.WithTimeout(c.Request().Context(), 30*time.Second)
		executionID = c.Param("executionID")
	)
	defer cancel()

	data, err := handler.useCase.GetDataByID(ctx, executionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewResponseError(
			http.StatusInternalServerError,
			response.MsgFailed,
			err.Error()),
		)
	}

	return c.JSON(http.StatusOK, response.NewResponse(http.StatusOK, response.MsgSuccess, data))
}
