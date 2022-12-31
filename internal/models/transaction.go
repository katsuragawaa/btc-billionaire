package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Transaction struct {
	ID       uuid.UUID `json:"id" db:"id" validate:"omitempty,uuid"`
	Amount   float64   `json:"amount" db:"amount" validate:"required"`
	Datetime time.Time `json:"datetime" db:"datetime" validate:"required"`
}
