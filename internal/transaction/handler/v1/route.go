package v1

import "github.com/labstack/echo/v4"

func (handler handlers) TransactionRoutes(domain *echo.Group) {
	domain.GET("/", handler.GetData)
	domain.POST("/process", handler.ProcessFile)
}
