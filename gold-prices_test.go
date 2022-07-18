package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestGold_GetPrices(t *testing.T) {
	client := NewTestClient(func(r *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
			Header:     make(http.Header),
		}
	})

	g := Gold{
		Prices: nil,
		Client: client,
	}

	p, err := g.GetPrices()
	if err != nil {
		t.Error(err)
	}

	if p.Price != 1711.565 {
		t.Error("wrong price returned:", p.Price)
	}
}
