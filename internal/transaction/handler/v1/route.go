package v1

import "github.com/labstack/echo/v4"

func (handler handlers) TransactionRoutes(domain *echo.Group) {
	domain.POST("/process", handler.ProcessFile)
}
