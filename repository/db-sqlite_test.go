package repository

import (
	"testing"
	"time"
)

func TestSQLiteRepository_Migrate(t *testing.T) {
	err := testRepo.Mirate()

	if err != nil {
		t.Error("migrate failed:", err)
	}
}

func TestSQLiteRepository_InsertHolding(t *testing.T) {
	h := Holdings{
		Amount:        1,
		PurchaseDate:  time.Now(),
		PurchasePrice: 1000,
	}

	result, err := testRepo.InsertHolding(h)

	if err != nil {
		t.Error("insert failed:", err)
	}

	if result.ID <= 0 {
		t.Error("invalid id sent back:", result.ID)
	}
}

func TestSQLiteRepository_AllHoldings(t *testing.T) {
	h, err := testRepo.AllHoldings()
	if err != nil {
		t.Error("get all failed:", err)
	}

	if len(h) != 1 {
		t.Errorf("wrong number of rows returned; expected 1, but got %d:", len(h))
	}
}
