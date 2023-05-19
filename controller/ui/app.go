package ui

import (
	"bytes"
	"errors"
	"image/color"
	"net/url"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/avelex/passwordme/internal/generator"
)

// TODO: remove magic numbers
type uiController struct {
	generator              *generator.PasswordGenerator
	logo, icon, background []byte
}

func NewUI(generator *generator.PasswordGenerator, icon, logo, background []byte) *uiController {
	return &uiController{
		generator:  generator,
		icon:       icon,
		logo:       logo,
		background: background,
	}
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

	tabs := container.NewAppTabs(
		ui.passwordOnFlightTab(mainWindow.Clipboard()),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	unlockButton := widget.NewButton("Unlock", func() {
		mainWindow.SetContent(tabs)
	})

	image := canvas.NewImageFromReader(bytes.NewReader(ui.logo), "Logo")
	image.FillMode = canvas.ImageFillContain

	con := container.NewAdaptiveGrid(1, container.NewPadded(image), container.NewCenter(unlockButton))

	mainWindow.SetContent(con)

	return app
}

func (ui *uiController) passwordOnFlightTab(clipboard fyne.Clipboard) *container.TabItem {
	const tabName = "On-Flight Generation"

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.Validator = func(s string) error {
		if strings.TrimSpace(s) == "" {
			return errors.New("required password")
		}

		return nil
	}

	siteURLEntry := widget.NewEntry()
	siteURLEntry.SetPlaceHolder("github.com")
	siteURLEntry.Validator = func(s string) error {
		if strings.TrimSpace(s) == "" {
			return errors.New("required url")
		}

		if !isValidUrl(s) {
			return errors.New("invalid url")
		}

		return nil
	}

	promptsEntry := widget.NewEntry()
	promptsEntry.SetPlaceHolder("alex, avelex")

	lengthText := widget.NewTextGridFromString("8")
	lengthSlider := widget.NewSlider(6, 32)
	lengthSlider.SetValue(8)
	lengthSlider.OnChanged = func(f float64) {
		lengthText.SetText(strconv.Itoa(int(f)))
	}

	lengthEntry := container.NewVBox(container.NewCenter(lengthText), lengthSlider)

	form := widget.NewForm(
		widget.NewFormItem("Master Password", passwordEntry),
		widget.NewFormItem("Site URL", siteURLEntry),
		widget.NewFormItem("Prompts", promptsEntry),
		widget.NewFormItem("Length", lengthEntry),
	)
	form.SubmitText = "Generate"

	box := container.NewAdaptiveGrid(1)

	form.OnSubmit = func() {
		master := passwordEntry.Text
		domain := siteURLEntry.Text
		prompts := strings.Split(promptsEntry.Text, ",")
		url, _ := url.Parse(domain)
		length := uint8(lengthSlider.Value)

		password := ui.generator.Generate(master, url, prompts, generator.WithLength(length))

		if len(box.Objects) > 1 {
			box.Remove(box.Objects[len(box.Objects)-1])
		}

		generatedPasswordEntry := widget.NewEntry()
		generatedPasswordEntry.SetText(password)

		yourPasswordText := canvas.NewText("Your password:", color.Black)
		yourPasswordText.Alignment = fyne.TextAlignLeading
		yourPasswordText.TextStyle = fyne.TextStyle{Bold: true, Symbol: true}

		box.Add(container.New(layout.NewFormLayout(),
			yourPasswordText,
			generatedPasswordEntry,
		))
		form.Refresh()
	}

	box.Add(form)

	return container.NewTabItem(tabName,
		container.New(
			layout.NewMaxLayout(),
			box,
		))
}
