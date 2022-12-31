package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
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
