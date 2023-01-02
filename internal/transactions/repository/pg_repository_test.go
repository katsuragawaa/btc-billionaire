package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/katsuragawaa/btc-billionaire/pkg/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"

	"github.com/katsuragawaa/btc-billionaire/internal/models"
)

func Test_transactionsRepo_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	repository := NewTransactionsRepository(sqlxDB)

	t.Run(
		"Create", func(t *testing.T) {
			amount := 1.0
			now := time.Now()

			rows := sqlmock.NewRows([]string{"amount", "datetime"}).AddRow(amount, now)

			transaction := &models.Transaction{
				Amount:   amount,
				Datetime: now,
			}

			mock.ExpectQuery(createTransaction).WithArgs(transaction.Amount, transaction.Datetime).WillReturnRows(rows)

			createdTransaction, err := repository.Create(context.Background(), transaction)

			require.NoError(t, err)
			require.NotNil(t, createdTransaction)
			require.Equal(t, createdTransaction, transaction)
		},
	)

	t.Run(
		"Create ERR", func(t *testing.T) {
			amount := 1.0
			now := time.Now()

			createErr := errors.New("create transaction error")

			transaction := &models.Transaction{
				Amount:   amount,
				Datetime: now,
			}

			mock.ExpectQuery(createTransaction).WithArgs(transaction.Amount, transaction.Datetime).WillReturnError(createErr)

			createdTransaction, err := repository.Create(context.Background(), transaction)

			require.NotNil(t, err)
			require.Nil(t, createdTransaction)
		},
	)
}

func Test_transactionsRepo_GetBalance(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	repository := NewTransactionsRepository(sqlxDB)

	t.Run(
		"Get balance", func(t *testing.T) {
			amount := 1.0

			rows := sqlmock.NewRows([]string{"total_amount"}).AddRow(amount)

			transaction := &models.TransactionsBalance{
				Total: amount,
			}

			mock.ExpectQuery(getBalance).WillReturnRows(rows)

			balance, err := repository.GetBalance(context.Background())

			require.NoError(t, err)
			require.NotNil(t, balance)
			require.Equal(t, balance, transaction)
		},
	)

	t.Run(
		"Get balance ERR", func(t *testing.T) {
			getBalanceErr := errors.New("get balance error")

			mock.ExpectQuery(getBalance).WillReturnError(getBalanceErr)

			balance, err := repository.GetBalance(context.Background())

			require.NotNil(t, err)
			require.Nil(t, balance)
		},
	)
}

func Test_transactionsRepo_GetPerHours(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	repository := NewTransactionsRepository(sqlxDB)

	t.Run(
		"Get transactions per hour", func(t *testing.T) {
			amount := 1.0
			start := time.Date(2023, 01, 01, 01, 00, 00, 00, time.UTC)
			datetime := start.AddDate(0, 6, 0)
			end := start.AddDate(1, 0, 0)

			rows := sqlmock.NewRows([]string{"amount", "datetime"}).AddRow(amount, datetime)

			transaction := &models.TransactionBase{
				Amount: amount,
				Datetime: utils.JSONTime{
					Time: datetime,
				},
			}
			transactionsList := &models.TransactionsList{
				Transactions: []*models.TransactionBase{transaction},
			}

			mock.ExpectQuery(getTransactionsByHour).WithArgs(end).WillReturnRows(rows)

			transactionsPerHour, err := repository.GetPerHours(context.Background(), start, end)

			require.NoError(t, err)
			require.NotNil(t, transactionsPerHour)
			require.Equal(t, transactionsPerHour, transactionsList)
		},
	)

	t.Run(
		"Get balance ERR", func(t *testing.T) {
			start := time.Date(2023, 01, 01, 01, 00, 00, 00, time.UTC)
			end := start.AddDate(1, 0, 0)

			getPerHoursErr := errors.New("get transactions per hour error")

			mock.ExpectQuery(getTransactionsByHour).WithArgs(end).WillReturnError(getPerHoursErr)

			transactionsPerHour, err := repository.GetPerHours(context.Background(), start, end)

			require.NotNil(t, err)
			require.Nil(t, transactionsPerHour)
		},
	)
}
