package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

const (
	_ON_FLIGHT_GENERATION = "On-Flight Generation"
	_PASSWORD_REPOSITORY  = "Password Repository"
)

type Screen struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
}

var (
	ScreenIndex = map[string][]string{
		"": {_ON_FLIGHT_GENERATION, _PASSWORD_REPOSITORY},
	}
)

func (ui *uiController) makeNav(setScreen func(screen Screen)) fyne.CanvasObject {
	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return ScreenIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := ScreenIndex[uid]
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := ui.screens[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
			obj.(*widget.Label).TextStyle = fyne.TextStyle{}
		},
		OnSelected: func(uid string) {
			if screen, ok := ui.screens[uid]; ok {
				setScreen(screen)
			}
		},
	}

	tree.Select(_ON_FLIGHT_GENERATION)

	return tree
}
