package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"github.com/petrostrak/exchangeME/repository"

	"fyne.io/fyne/v2/widget"
	_ "github.com/glebarez/go-sqlite"
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
	DB                  repository.Repository
	MainWindow          fyne.Window
	PriceContainer      *fyne.Container
	ToolBar             *widget.Toolbar
	PriceChartContainer *fyne.Container
	HTTPClient          *http.Client
}

func (c *Config) connectToSLQ() (*sql.DB, error) {
	path := ""

	if os.Getenv("DB_PATH") != "" {
		path = os.Getenv("DB_PATH")
	} else {
		path = c.App.Storage().RootURI().Path() + "/sql.db"
		c.InfoLog.Println("db in:", path)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (c *Config) setupDB(sqlDB *sql.DB) {
	c.DB = repository.NewSQLiteRepository(sqlDB)

	err := c.DB.Migrate()
	if err != nil {
		c.ErrorLog.Println(err)
		log.Panic()
	}
}
