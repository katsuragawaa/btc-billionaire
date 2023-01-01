package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/katsuragawaa/btc-billionaire/internal/models"
	"github.com/katsuragawaa/btc-billionaire/internal/transactions"
)

type transactionsRepo struct {
	db *sqlx.DB
}

func NewTransactionsRepository(db *sqlx.DB) transactions.Repository {
	return &transactionsRepo{db: db}
}

func (r *transactionsRepo) Create(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	t := &models.Transaction{}
	if err := r.db.QueryRowxContext(ctx, createTransaction, transaction.Amount, transaction.Datetime).StructScan(t); err != nil {
		return nil, errors.Wrap(err, "transactionsRepo.Create.StructScan")
	}

	return t, nil
}

func (r *transactionsRepo) GetPerHours(ctx context.Context, start time.Time, end time.Time) (*models.TransactionsList, error) {
	rows, err := r.db.QueryxContext(ctx, getTransactionsByHour, end)
	if err != nil {
		return nil, errors.Wrap(err, "transactionsRepo.GetPerHours")
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	acc := 0.0
	transactionsList := make([]*models.TransactionBase, 0)
	for rows.Next() {
		transaction := &models.TransactionBase{}
		if err = rows.StructScan(transaction); err != nil {
			return nil, errors.Wrap(err, "transactionsRepo.GetPerHours.StructScan")
		}

		acc += transaction.Amount
		transaction.Amount = acc

		if start.Before(transaction.Datetime) {
			transactionsList = append(transactionsList, transaction)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "transactionsRepo.GetPerHours.rows.Err")
	}

	return &models.TransactionsList{Transactions: transactionsList}, nil
}
