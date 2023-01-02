package models

import (
	"time"

	"github.com/google/uuid"

	"github.com/katsuragawaa/btc-billionaire/pkg/utils"
)

type Transaction struct {
	ID        uuid.UUID `json:"id" db:"id" validate:"omitempty,uuid"`
	Amount    float64   `json:"amount" db:"amount" validate:"required"`
	Datetime  time.Time `json:"datetime" db:"datetime" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type TransactionBase struct {
	Amount   float64        `json:"amount" db:"amount"`
	Datetime utils.JSONTime `json:"datetime" db:"datetime"`
}

type TransactionsList struct {
	Transactions []*TransactionBase `json:"transactions"`
}

type TransactionsBalance struct {
	Total float64 `json:"total_amount" db:"total_amount"`
}
