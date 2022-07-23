package main

import "testing"

func TestConfig_getHoldings(t *testing.T) {
	all, err := testApp.currentHoldings()
	if err != nil {
		t.Error("failed to get current holdings from DB", err)
	}

	if len(all) != 2 {
		t.Error("wrong number of rows returned:", len(all))
	}
}
