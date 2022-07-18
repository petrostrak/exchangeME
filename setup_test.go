package main

import (
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
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
