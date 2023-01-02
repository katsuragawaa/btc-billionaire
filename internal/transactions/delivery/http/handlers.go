package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/katsuragawaa/btc-billionaire/config"
	"github.com/katsuragawaa/btc-billionaire/internal/models"
	"github.com/katsuragawaa/btc-billionaire/internal/transactions"
	"github.com/katsuragawaa/btc-billionaire/pkg/logger"
	"github.com/katsuragawaa/btc-billionaire/pkg/utils"
)

type transactionsHandlers struct {
	cfg     *config.Config
	usecase transactions.UseCase
	logger  logger.Logger
}

func NewTransactionsHandlers(cfg *config.Config, usecase transactions.UseCase, logger logger.Logger) transactions.Handlers {
	return &transactionsHandlers{cfg: cfg, usecase: usecase, logger: logger}
}

// Create
//
//	@Summary		Create new transaction
//	@Description	Create new bitcoin transaction
//	@Tags			Transaction
//	@Accept			json
//	@Produce		json
//	@Param			transaction	body		models.TransactionBase	true	"Send new transaction"
//	@Success		201			{object}	models.Transaction
//	@Failure		400
//	@Failure		500
//	@Router			/transactions [post]
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
//
//	@Summary		Get transactions within a time interval
//	@Description	Get bitcoin transactions within a time interval
//	@Tags			Transaction
//	@Accept			json
//	@Produce		json
//	@Param			startDatetime	query		string	true	"interval start datetime"
//	@Param			endDatetime		query		string	true	"interval end datetime"
//	@Success		200				{object}	models.TransactionsList
//	@Failure		400
//	@Failure		500
//	@Router			/transactions [get]
func (t *transactionsHandlers) GetPerHours() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		start := c.QueryParam("startDatetime")
		end := c.QueryParam("endDatetime")
		if start == "" || end == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "please specify a start and an end datetime")
		}

		startTime, err := utils.ParseTime(start)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid start datetime, please format as 2006-01-02T15:04:05-07:00")
		}
		endTime, err := utils.ParseTime(end)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid end datetime, please format as 2006-01-02T15:04:05-07:00")
		}

		transactionsList, err := t.usecase.GetPerHours(ctx, startTime, endTime)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, transactionsList)
	}
}

// GetBalance
//
//	@Summary		Get wallet total balance
//	@Description	Get total bitcoin balance in the wallet
//	@Tags			Transaction
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.TransactionsBalance
//	@Failure		500
//	@Router			/transactions/balance [get]
func (t *transactionsHandlers) GetBalance() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		balance, err := t.usecase.GetBalance(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, balance)
	}
}
