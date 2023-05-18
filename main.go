package main

import (
	"bytes"
	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	//go:embed assets/img/Icon.png
	icon []byte
	//go:embed assets/img/logo.png
	logo []byte
)

func main() {
	icon := canvas.NewImageFromReader(bytes.NewReader(icon), "Icon")

	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	a.SetIcon(icon.Resource)

	mainWindow := a.NewWindow("PasswordME")
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.CenterOnScreen()
	mainWindow.SetFixedSize(true)

	tabs := container.NewAppTabs(
		container.NewTabItem("Generate On-Flight", container.NewCenter(widget.NewButton("Generate", func() {}))),
		container.NewTabItem("Passwords", widget.NewList(func() int { return 3 }, func() fyne.CanvasObject { return &fyne.Container{} }, func(lii widget.ListItemID, co fyne.CanvasObject) {})),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	unlockButton := widget.NewButton("Unlock", func() {
		mainWindow.SetContent(tabs)
	})

	image := canvas.NewImageFromReader(bytes.NewReader(logo), "Logo")
	image.FillMode = canvas.ImageFillContain

	con := container.NewAdaptiveGrid(1, container.NewPadded(image), container.NewCenter(unlockButton))

	mainWindow.SetContent(con)

	mainWindow.ShowAndRun()
}
