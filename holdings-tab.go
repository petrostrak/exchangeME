package main

import (
	"fmt"
	"strconv"

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

	holdings, err := c.currentHoldings()
	if err != nil {
		c.ErrorLog.Println(err)
	}

	slice = append(slice, []any{"ID", "Amount", "Price", "Date", "Delete?"})

	for _, x := range holdings {
		var currentRow []any

		currentRow = append(currentRow, strconv.FormatInt(x.ID, 10))
		currentRow = append(currentRow, fmt.Sprintf("%d toz", x.Amount))
		currentRow = append(currentRow, fmt.Sprintf("$%2f", float32(x.PurchasePrice/100)))
		currentRow = append(currentRow, x.PurchaseDate.Format("2006-01-02"))
		currentRow = append(currentRow, widget.NewButton("Delete", func() {}))

		slice = append(slice, currentRow)
	}

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
