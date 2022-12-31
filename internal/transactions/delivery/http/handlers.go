package http

import (
	"github.com/katsuragawaa/btc-billionaire/config"
	"github.com/katsuragawaa/btc-billionaire/internal/transactions"
	"github.com/katsuragawaa/btc-billionaire/pkg/logger"
	"github.com/labstack/echo/v4"
)

type transactionsHandlers struct {
	cfg    *config.Config
	logger logger.Logger
}

func (t *transactionsHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		t.logger.Info("Transactions handler - Create")
		return nil
	}
}

func NewTransactionsHandlers(cfg *config.Config, logger logger.Logger) transactions.Handlers {
	return &transactionsHandlers{cfg: cfg, logger: logger}
}
