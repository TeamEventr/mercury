// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: price_tier.sql

package db

import (
	"time"

	"github.com/google/uuid"
)

type CreatePriceTierQueryParams struct {
	EventID       uuid.UUID
	Name          string
	ValidityStart time.Time
	ValidityEnd   time.Time
	Price         int32
	SeatAvailable int32
}
