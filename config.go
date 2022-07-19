package main

import (
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Config is the type used to share data with various parts of our application.
// It includes the parts of our GUI that are dynamic and will need to be updated,
// such as the holdings table, gold price info, and the chart. In order to refresh
// those things, we need a reference to them, and this is a convenient place to put
// them, instead of package level variables.
type Config struct {
	App                 fyne.App
	InfoLog             *log.Logger
	ErrorLog            *log.Logger
	MainWindow          fyne.Window
	PriceContainer      *fyne.Container
	ToolBar             *widget.Toolbar
	PriceChartContainer *fyne.Container
	HTTPClient          *http.Client
}
