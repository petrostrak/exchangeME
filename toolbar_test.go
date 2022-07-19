package main

import "testing"

func TestApp_getToolbar(t *testing.T) {
	tb := testApp.getToolBar()

	if len(tb.Items) != 4 {
		t.Error("wrong number of toolbar-items", len(tb.Items))
	}
}
