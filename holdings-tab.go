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
	c.HoldingsTable = c.getHoldingsTable()

	return nil
}

func (c *Config) getHoldingsTable() *widget.Table {
	data := c.getHoldingSlice()
	c.Holdings = data

	t := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return container.NewVBox(widget.NewLabel(""))
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == len(data[0])-1 && i.Row != 0 {
				// last cell in row, put a button
				w := widget.NewButtonWithIcon("Delete", theme.DeleteIcon(), func() {
					dialog.ShowConfirm("Delete?", "", func(deleted bool) {
						id, _ := strconv.Atoi(data[i.Row][0].(string))
						err := c.DB.DeleteHolding(int64(id))
						if err != nil {
							c.ErrorLog.Println(err)
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
					widget.NewLabel(data[i.Row][i.Col].(string)),
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
