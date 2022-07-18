package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	exchangeApp Config
)

func main() {
	// create a fyne app
	exchange := app.NewWithID("app.netlify.petrostrak.exchangeMe.preferences")
	exchangeApp.App = exchange

	// create loggers
	exchangeApp.InfoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	exchangeApp.ErrorLog = log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open connection to DB

	// create a DB repository

	// create and size a fyne window
	win := exchange.NewWindow("exchangeME!")
	win.CenterOnScreen()
	win.Resize(fyne.Size{Width: 800, Height: 500})

	// show and run app
	win.ShowAndRun()
}
