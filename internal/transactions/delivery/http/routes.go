package http

import (
	"github.com/labstack/echo/v4"

	"github.com/katsuragawaa/btc-billionaire/internal/transactions"
)

func MapTransactionsRoutes(transactionsGroup *echo.Group, h transactions.Handlers) {
	transactionsGroup.POST("", h.Create())
	transactionsGroup.GET("", h.GetPerHours())
	transactionsGroup.GET("/balance", h.GetBalance())
}
