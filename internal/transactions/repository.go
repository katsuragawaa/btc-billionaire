package transactions

import (
	"context"

	"github.com/katsuragawaa/btc-billionaire/internal/models"
)

type Repository interface {
	Create(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
}
