package http

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/katsuragawaa/btc-billionaire/config"
	"github.com/katsuragawaa/btc-billionaire/internal/models"
	"github.com/katsuragawaa/btc-billionaire/internal/transactions"
	"github.com/katsuragawaa/btc-billionaire/pkg/logger"
	"github.com/katsuragawaa/btc-billionaire/pkg/utils"
)

const layout = "2006-01-02T15:04:05-07:00"

type transactionsHandlers struct {
	cfg     *config.Config
	usecase transactions.UseCase
	logger  logger.Logger
}

func NewTransactionsHandlers(cfg *config.Config, usecase transactions.UseCase, logger logger.Logger) transactions.Handlers {
	return &transactionsHandlers{cfg: cfg, usecase: usecase, logger: logger}
}

// Create
// @Summary Create new transaction
// @Description Create new bitcoin transaction
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Success 201 {object} models.Transaction
// @Failure 400
// @Failure 500
// @Router /comments [post]
func (t *transactionsHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		transaction := &models.Transaction{}
		if err := utils.BindRequest(c, transaction); err != nil {
			t.logger.Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		createdTransaction, err := t.usecase.Create(ctx, transaction)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, createdTransaction)
	}
}

// GetPerHours
// @Summary Get transactions within a time interval
// @Description Get bitcoin transactions within a time interval
// @Tags Transaction
// @Accept json
// @Produce json
// @Success 200 {object} models.TransactionsList
// @Failure 500
// @Router /comments [get]
func (t *transactionsHandlers) GetPerHours() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		start := c.QueryParam("start")
		end := c.QueryParam("end")

		startTime, err := time.Parse(layout, start)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		endTime, err := time.Parse(layout, end)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		transactionsList, err := t.usecase.GetPerHours(ctx, startTime, endTime)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, transactionsList)
	}
}
