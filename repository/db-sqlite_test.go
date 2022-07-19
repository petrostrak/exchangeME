package repository

import "testing"

func TestSQLiteRepository_Migrate(t *testing.T) {
	err := testRepo.Mirate()

	if err != nil {
		t.Error("migrate failed:", err)
	}
}
