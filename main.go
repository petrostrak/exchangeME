package main

import (
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	var exchangeApp Config

	// create a fyne app
	exchange := app.NewWithID("app.netlify.petrostrak.exchangeMe.preferences")
	exchangeApp.App = exchange
	exchangeApp.HTTPClient = &http.Client{}

	// create loggers
	exchangeApp.InfoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	exchangeApp.ErrorLog = log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open connection to DB
	sqlDB, err := exchangeApp.connectToSLQ()
	if err != nil {
		log.Panic(err)
	}

	// create a DB repository
	exchangeApp.setupDB(sqlDB)

	currency = exchange.Preferences().StringWithFallback("currency", "JPY")

	// create and size a fyne window
	exchangeApp.MainWindow = exchange.NewWindow("exchangeME!")
	exchangeApp.MainWindow.Resize(fyne.NewSize(770, 410))
	exchangeApp.MainWindow.SetFixedSize(true)
	exchangeApp.MainWindow.SetMaster()
	exchangeApp.MainWindow.CenterOnScreen()

	exchangeApp.makeUI()

	// show and run app
	exchangeApp.MainWindow.ShowAndRun()
}
