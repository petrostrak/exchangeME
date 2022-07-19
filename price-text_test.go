package main

import "testing"

func TestApp_getPriceText(t *testing.T) {
	open, _, _ := testApp.getPriceText()
	if open.Text != "Open: $1710.5935 USD" {
		t.Error("wrong price returned", open.Text)
	}
}
