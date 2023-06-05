package ui

import (
	"bytes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/avelex/passwordme/internal/generator"
	"github.com/avelex/passwordme/internal/store"
)

type uiController struct {
	store                  *store.PasswordStore
	generator              *generator.PasswordGenerator
	logo, icon, background []byte
	screens                map[string]Screen
}

func NewUI(generator *generator.PasswordGenerator, store *store.PasswordStore, icon, logo, background []byte) *uiController {
	controller := &uiController{
		generator:  generator,
		store:      store,
		icon:       icon,
		logo:       logo,
		background: background,
	}

	controller.screens = map[string]Screen{
		_ON_FLIGHT_GENERATION: controller.generationScreen(),
		_PASSWORD_REPOSITORY:  controller.passwordsScreen(),
	}

	return controller
}

func (ui *uiController) Run() {
	app := ui.init()
	app.Run()
}

func (ui *uiController) init() fyne.App {
	icon := canvas.NewImageFromReader(bytes.NewReader(ui.icon), "Icon")

	app := app.New()
	app.Settings().SetTheme(theme.LightTheme())
	app.SetIcon(icon.Resource)

	mainWindow := app.NewWindow("PasswordME")
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.CenterOnScreen()
	mainWindow.SetFixedSize(true)
	mainWindow.Show()

	content := container.NewMax()
	setScreen := func(s Screen) {
		content.Objects = []fyne.CanvasObject{s.View(mainWindow)}
		content.Refresh()
	}

	nav := ui.makeNav(setScreen)

	unlockButton := widget.NewButton("Unlock", func() {
		split := container.NewHSplit(
			container.NewBorder(nil, widget.NewLabel("v"+app.Metadata().Version), nil, nil, nav),
			content,
		)
		split.Offset = 0.25
		mainWindow.SetContent(split)
	})

	image := canvas.NewImageFromReader(bytes.NewReader(ui.logo), "Logo")
	image.FillMode = canvas.ImageFillContain

	welcomeContent := container.NewAdaptiveGrid(1, container.NewPadded(image), container.NewCenter(unlockButton))
	mainWindow.SetContent(welcomeContent)

	return app
}
