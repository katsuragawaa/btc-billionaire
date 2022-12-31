package usecase

import (
	"context"

	"github.com/katsuragawaa/btc-billionaire/config"
	"github.com/katsuragawaa/btc-billionaire/internal/models"
	"github.com/katsuragawaa/btc-billionaire/internal/transactions"
	"github.com/katsuragawaa/btc-billionaire/pkg/logger"
)

type transactionsUC struct {
	cfg    *config.Config
	repo   transactions.Repository
	logger logger.Logger
}

func NewTransactionsUseCase(cfg *config.Config, repo transactions.Repository, logger logger.Logger) transactions.UseCase {
	return &transactionsUC{
		cfg:    cfg,
		repo:   repo,
		logger: logger,
	}
}

func (u *transactionsUC) Create(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	return u.repo.Create(ctx, transaction)
}
