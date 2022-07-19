package main

import (
	"time"

	"fyne.io/fyne/v2"
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
	finalContent := container.NewVBox(priceContent, toolbar, tabs)

	c.MainWindow.SetContent(finalContent)

	go func() {
		for range time.Tick(time.Second * 30) {
			c.refreshPriceContent()
		}
	}()
}

func (c *Config) refreshPriceContent() {
	c.InfoLog.Println("refreshing prices")
	open, current, change := c.getPriceText()
	c.PriceContainer.Objects = []fyne.CanvasObject{open, current, change}
	c.PriceContainer.Refresh()

	chart := c.getChart()
	c.PriceChartContainer.Objects = []fyne.CanvasObject{chart}
	c.PriceChartContainer.Refresh()
}
