package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (c *Config) pricesTab() *fyne.Container {
	chart := c.getChart()
	chartContainer := container.NewVBox(chart)
	c.PriceContainer = chartContainer

	return chartContainer
}

func (c *Config) getChart() *canvas.Image {
	apiURL := fmt.Sprintf("https://goldprice.org/charts/gold_3d_b_o_%s_x.png", strings.ToLower(currency))
	var img *canvas.Image

	err := c.downloadFile(apiURL, "gold.png")
	if err != nil {
		// user bundled image
		// generated with `fyne bundle unreachable.png >> bundled.go`
		img = canvas.NewImageFromResource(resourceUnreachablePng)
	} else {
		img = canvas.NewImageFromFile("gold.png")
	}

	img.SetMinSize(fyne.Size{Width: 770, Height: 410})

	img.FillMode = canvas.ImageFillOriginal

	return img
}

func (c *Config) downloadFile(URL, filename string) error {
	// get the response bytes from calling a url
	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("received wrong response code when downloading image")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}

	out, err := os.Create(fmt.Sprintf("./%s", filename))
	if err != nil {
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}
