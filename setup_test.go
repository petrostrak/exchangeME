package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/petrostrak/exchangeME/repository"
)

var testApp Config

func TestMain(m *testing.M) {
	a := test.NewApp()
	testApp.App = a
	testApp.MainWindow = a.NewWindow("")
	testApp.HTTPClient = client
	testApp.DB = repository.NewTestRepository()
	os.Exit(m.Run())
}

var jsonToReturn = `
{
	"ts": 1658160618758,
	"tsj": 1658160615441,
	"date": "Jul 18th 2022, 12:10:15 pm NY",
	"items": [
	  {
		"curr": "USD",
		"xauPrice": 1711.565,
		"xagPrice": 18.8865,
		"chgXau": 0.9715,
		"chgXag": 0.138,
		"pcXau": 0.0568,
		"pcXag": 0.7361,
		"xauClose": 1710.59346,
		"xagClose": 18.74853
	  }
	]
  }
`

type RoundTripFunc func(*http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(r *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
