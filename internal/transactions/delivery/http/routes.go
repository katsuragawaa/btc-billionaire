package http

import (
	"github.com/labstack/echo/v4"

	"github.com/katsuragawaa/btc-billionaire/internal/transactions"
)

func MapTransactionsRoutes(commGroup *echo.Group, h transactions.Handlers) {
	commGroup.POST("", h.Create())
	commGroup.GET("", h.GetPerHours())
	commGroup.GET("/balance", h.GetBalance())
}
