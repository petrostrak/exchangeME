package main

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (c *Config) getToolBar() (toolbar *widget.Toolbar) {
	toolbar = widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			c.refreshPriceContent()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
	)

	return
}
