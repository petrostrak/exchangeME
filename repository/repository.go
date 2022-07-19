package repository

import (
	"errors"
	"time"
)

var (
	errUpdateFailed = errors.New("update failed")
	errDeleteFailed = errors.New("delete failed")
)

type Repository interface {
	Migrate() error
	InsertHolding(Holdings) (*Holdings, error)
	AllHoldings() ([]Holdings, error)
	GetHoldingByID(int) (*Holdings, error)
	UpdateHolding(int64, Holdings) error
	DeleteHolding(int64) error
}

type Holdings struct {
	ID            int64     `json:"id"`
	Amount        int       `json:"amount"`
	PurchaseDate  time.Time `json:"purchase_date"`
	PurchasePrice int       `json:"purchase_price"`
}
