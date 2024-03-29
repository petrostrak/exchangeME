package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/petrostrak/exchangeME/repository"
)

func (c *Config) holdingsTab() *fyne.Container {
	c.Holdings = c.getHoldingSlice()
	c.HoldingsTable = c.getHoldingsTable()

	holdingsContainer := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, c.HoldingsTable),
	)

	return holdingsContainer
}

func (c *Config) getHoldingsTable() *widget.Table {
	t := widget.NewTable(
		func() (int, int) {
			return len(c.Holdings), len(c.Holdings[0])
		},
		func() fyne.CanvasObject {
			return container.NewVBox(widget.NewLabel(""))
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == len(c.Holdings[0])-1 && i.Row != 0 {
				// last cell in row, put a button
				w := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Delete?", "", func(deleted bool) {
						if deleted {
							id, _ := strconv.Atoi(c.Holdings[i.Row][0].(string))
							err := c.DB.DeleteHolding(int64(id))
							if err != nil {
								c.ErrorLog.Println(err)
							}
						}

						// refresh the holdings table
						c.refreshHoldingsTable()
					}, c.MainWindow)
				})
				w.Importance = widget.HighImportance

				o.(*fyne.Container).Objects = []fyne.CanvasObject{w}
			} else {
				// we are putting textual information
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(c.Holdings[i.Row][i.Col].(string)),
				}
			}
		})

	colwidth := []float32{50, 200, 200, 200, 110}
	for i := 0; i < len(colwidth); i++ {
		t.SetColumnWidth(i, colwidth[i])
	}

	return t
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

		currentAmount := float32(x.PurchasePrice)
		currentAmount = currentAmount / 100

		currentRow = append(currentRow, fmt.Sprintf("$%.2f", currentAmount))
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
