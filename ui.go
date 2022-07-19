package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (c *Config) makeUI() {
	// get current price of gold
	openPrice, currentPrice, priceChange := c.getPriceText()

	// put price info into a container
	priceContent := container.NewGridWithColumns(3, openPrice, currentPrice, priceChange)
	c.PriceContainer = priceContent

	// get toolbar
	toolbar := c.getToolBar()
	c.ToolBar = toolbar

	priceTabContent := c.pricesTab()

	// get app tabs
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), priceTabContent),
		container.NewTabItemWithIcon("Holdings", theme.InfoIcon(), canvas.NewText("holdings content goes here", nil)),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// add container to window
	finalContent := container.NewVBox(priceChange, toolbar, tabs)

	c.MainWindow.SetContent(finalContent)
}
