package main

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/petrostrak/exchangeME/repository"
)

func (c *Config) getToolBar() (toolbar *widget.Toolbar) {
	toolbar = widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			c.addHoldingsDialog()
		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			c.refreshPriceContent()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
	)

	return
}

func (c *Config) addHoldingsDialog() dialog.Dialog {
	addAmountEntry := widget.NewEntry()
	purchaseDateEntry := widget.NewEntry()
	purchasePriceEntry := widget.NewEntry()

	c.AddHoldingsPurchaseAmountEntry = addAmountEntry
	c.AddHoldingsPurchaseDateEntry = purchaseDateEntry
	c.AddHoldingsPurchasePriceEntry = purchasePriceEntry

	purchaseDateEntry.PlaceHolder = "YYYY-MM-DD"

	dateValidator := func(s string) error {
		if _, err := time.Parse("2006-01-02", s); err != nil {
			return err
		}

		return nil
	}
	purchaseDateEntry.Validator = dateValidator

	isIntValidator := func(s string) error {
		if _, err := strconv.Atoi(s); err != nil {
			return err
		}

		return nil
	}
	addAmountEntry.Validator = isIntValidator

	isFloatValidator := func(s string) error {
		if _, err := strconv.ParseFloat(s, 32); err != nil {
			return err
		}

		return nil
	}
	purchasePriceEntry.Validator = isFloatValidator

	// create a dialog
	addForm := dialog.NewForm(
		"Add Gold",
		"Add",
		"Cancel",
		[]*widget.FormItem{
			{Text: "Amount in toz", Widget: addAmountEntry},
			{Text: "Purchase Price", Widget: purchasePriceEntry},
			{Text: "Purchase Date", Widget: purchaseDateEntry},
		},
		func(valid bool) {
			if valid {
				amount, _ := strconv.Atoi(addAmountEntry.Text)
				purchaseDate, _ := time.Parse("2006-01-02", purchaseDateEntry.Text)
				purchasePrice, _ := strconv.ParseFloat(purchasePriceEntry.Text, 32)
				purchasePrice = purchasePrice * 100

				_, err := c.DB.InsertHolding(repository.Holdings{
					Amount:        amount,
					PurchaseDate:  purchaseDate,
					PurchasePrice: int(purchasePrice),
				})
				if err != nil {
					c.ErrorLog.Println(err)
				}

				c.refreshHoldingsTable()
			}
		}, c.MainWindow)

	//size and show the dialog
	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()

	return addForm
}
