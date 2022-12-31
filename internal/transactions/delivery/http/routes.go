package http

import (
	"github.com/katsuragawaa/btc-billionaire/internal/transactions"
	"github.com/labstack/echo/v4"
)

func MapTransactionsRoutes(commGroup *echo.Group, h transactions.Handlers) {
	commGroup.POST("", h.Create())
}
