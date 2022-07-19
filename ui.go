package main

import "fyne.io/fyne/v2/container"

func (c *Config) makeUI() {
	// get current price of gold
	openPrice, currentPrice, priceChange := c.getPriceText()

	// put price info into a container
	priceContent := container.NewGridWithColumns(3, openPrice, currentPrice, priceChange)
	c.PriceContainer = priceContent

	// get toolbar
	toolbar := c.getToolBar(c.MainWindow)

	// add container to window
	finalContent := container.NewVBox(priceChange, toolbar)

	c.MainWindow.SetContent(finalContent)
}
