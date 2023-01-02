package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"

	"github.com/katsuragawaa/btc-billionaire/internal/models"
	mock_transactions "github.com/katsuragawaa/btc-billionaire/internal/transactions/mock"
	"github.com/katsuragawaa/btc-billionaire/pkg/converter"
	"github.com/katsuragawaa/btc-billionaire/pkg/logger"
	"github.com/katsuragawaa/btc-billionaire/pkg/utils"
)

const baseURL = "/api/v1/transactions"

func Test_transactionsHandlers_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewAPILogger(nil)
	mockTransactionsUC := mock_transactions.NewMockUseCase(ctrl)
	tHandlers := NewTransactionsHandlers(nil, mockTransactionsUC, apiLogger)

	transaction := &models.Transaction{
		Amount:   1,
		Datetime: time.Date(2023, 01, 01, 9, 00, 00, 00, time.UTC),
	}

	buf, err := converter.AnyToBytesBuffer(transaction)
	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, baseURL, strings.NewReader(buf.String()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	ctx := context.Background()

	handlerFunc := tHandlers.Create()

	id := uuid.New()
	now := time.Now()

	mockCreatedTransaction := &models.Transaction{
		ID:        id,
		Amount:    transaction.Amount,
		Datetime:  transaction.Datetime,
		CreatedAt: now,
		UpdatedAt: now,
	}

	mockTransactionsUC.EXPECT().Create(ctx, gomock.Eq(transaction)).Return(mockCreatedTransaction, nil)

	err = handlerFunc(c)
	require.NoError(t, err)
	require.Nil(t, err)
}

func Test_transactionsHandlers_GetPerHours(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewAPILogger(nil)
	mockTransactionsUC := mock_transactions.NewMockUseCase(ctrl)
	tHandlers := NewTransactionsHandlers(nil, mockTransactionsUC, apiLogger)

	query := "?startDatetime=2023-01-01T18:00:00-05:00&endDatetime=2023-12-31T21:00:00+08:00"

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, baseURL+query, strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	ctx := context.Background()

	handlerFunc := tHandlers.GetPerHours()

	now := time.Now()

	transaction := &models.TransactionBase{
		Amount: 1,
		Datetime: utils.JSONTime{
			Time: now,
		},
	}
	transactionsList := &models.TransactionsList{
		Transactions: []*models.TransactionBase{transaction},
	}

	mockTransactionsUC.EXPECT().GetPerHours(ctx, gomock.Any(), gomock.Any()).Return(transactionsList, nil)

	err := handlerFunc(c)
	require.NoError(t, err)
	require.Nil(t, err)
}

func Test_transactionsHandlers_GetBalance(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewAPILogger(nil)
	mockTransactionsUC := mock_transactions.NewMockUseCase(ctrl)
	tHandlers := NewTransactionsHandlers(nil, mockTransactionsUC, apiLogger)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, baseURL+"/balance", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	ctx := context.Background()

	handlerFunc := tHandlers.GetBalance()

	balance := &models.TransactionsBalance{
		Total: 1000,
	}

	mockTransactionsUC.EXPECT().GetBalance(ctx).Return(balance, nil)

	err := handlerFunc(c)
	require.NoError(t, err)
	require.Nil(t, err)
}
