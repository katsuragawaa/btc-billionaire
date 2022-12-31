package http

import (
	"net/http"

	"github.com/katsuragawaa/btc-billionaire/config"
	"github.com/katsuragawaa/btc-billionaire/internal/models"
	"github.com/katsuragawaa/btc-billionaire/internal/transactions"
	"github.com/katsuragawaa/btc-billionaire/pkg/logger"
	"github.com/katsuragawaa/btc-billionaire/pkg/utils"
	"github.com/labstack/echo/v4"
)

type transactionsHandlers struct {
	cfg    *config.Config
	logger logger.Logger
}

func NewTransactionsHandlers(cfg *config.Config, logger logger.Logger) transactions.Handlers {
	return &transactionsHandlers{cfg: cfg, logger: logger}
}

// Create
// @Summary Create new transaction
// @Description Create new bitcoin transaction
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Success 201 {object} models.Transaction
// @Failure 400
// @Router /comments [post]
func (t *transactionsHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		t.logger.Info("Transactions handler - Create")

		var transaction models.Transaction
		if err := utils.BindRequest(c, &transaction); err != nil {
			t.logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, transaction)
	}
}
