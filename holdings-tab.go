package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/petrostrak/exchangeME/repository"
)

func (c *Config) holdingsTab() *fyne.Container {
	return nil
}

func (c *Config) getHoldingsTable() *widget.Table {
	return nil
}

func (c *Config) getHoldingSlice() [][]any {
	var slice [][]any

	return slice
}

func (c *Config) currentHoldings() ([]repository.Holdings, error) {
	holdings, err := c.DB.AllHoldings()
	if err != nil {
		c.ErrorLog.Println(err)
		return nil, err
	}

	return holdings, nil
}
