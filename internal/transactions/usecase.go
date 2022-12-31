package transactions

import (
	"context"

	"github.com/katsuragawaa/btc-billionaire/internal/models"
)

type UseCase interface {
	Create(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error)
}
