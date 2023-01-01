package transactions

import (
	"context"
	"time"

	"github.com/katsuragawaa/btc-billionaire/internal/models"
)

type Repository interface {
	Create(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
	GetPerHours(ctx context.Context, start time.Time, end time.Time) (*models.TransactionsList, error)
	GetBalance(ctx context.Context) (*models.TransactionsBalance, error)
}
