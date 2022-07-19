package main

import (
	"log"
	"net/http"

	"fyne.io/fyne/v2"
)

type Config struct {
	App            fyne.App
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	MainWindow     fyne.Window
	PriceContainer *fyne.Container
	HTTPClient     *http.Client
}
