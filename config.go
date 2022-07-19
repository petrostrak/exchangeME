package main

import (
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	App            fyne.App
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	MainWindow     fyne.Window
	PriceContainer *fyne.Container
	ToolBar        *widget.Toolbar
	HTTPClient     *http.Client
}
