package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func (ui *uiController) passwordsScreen() Screen {
	return Screen{
		Title: _PASSWORD_REPOSITORY,
		View:  ui.passwordsView,
	}
}

func (ui *uiController) passwordsView(_ fyne.Window) fyne.CanvasObject {
	list := widget.NewList(
		func() int { return len(ui.store.List()) },
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(ui.store.List()[id])
		},
	)

	return list
}
