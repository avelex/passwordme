package ui

import (
	"errors"
	"image/color"
	"net/url"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/avelex/passwordme/internal/generator"
)

func (ui *uiController) generationScreen() Screen {
	return Screen{
		Title: _ON_FLIGHT_GENERATION,
		View:  ui.passwordGenerationView,
	}
}

func (ui *uiController) passwordGenerationView(_ fyne.Window) fyne.CanvasObject {
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.Validator = func(s string) error {
		if strings.TrimSpace(s) == "" {
			return errors.New("required password")
		}

		return nil
	}

	siteDomainEntry := widget.NewEntry()
	siteDomainEntry.SetPlaceHolder("github.com")
	siteDomainEntry.Validator = func(s string) error {
		if strings.TrimSpace(s) == "" {
			return errors.New("required domain")
		}

		if !isValidUrl(s) {
			return errors.New("invalid domain")
		}

		return nil
	}

	promptsEntry := widget.NewEntry()
	promptsEntry.SetPlaceHolder("alex, avelex")

	lengthRG := widget.NewRadioGroup([]string{"16", "32"}, func(s string) {

	})
	lengthRG.Required = true
	lengthRG.Horizontal = true
	lengthRG.SetSelected("16")

	form := widget.NewForm(
		widget.NewFormItem("Master Password", passwordEntry),
		widget.NewFormItem("Site Domain", siteDomainEntry),
		widget.NewFormItem("Prompts", promptsEntry),
		widget.NewFormItem("Length", lengthRG),
	)
	form.SubmitText = "Generate"

	grid := container.NewAdaptiveGrid(1)

	form.OnSubmit = func() {
		master := passwordEntry.Text
		domain := siteDomainEntry.Text
		prompts := strings.Split(promptsEntry.Text, ",")

		url := &url.URL{
			Host: domain,
		}

		length, _ := strconv.Atoi(lengthRG.Selected)

		var opt generator.PasswordOpt
		switch length {
		case 16:
			opt = generator.WithLength16()
		default:
			opt = generator.WithLength32()
		}

		password := ui.generator.Generate(master, url, prompts, opt)

		if len(grid.Objects) > 1 {
			grid.Remove(grid.Objects[len(grid.Objects)-1])
		}

		generatedPasswordEntry := widget.NewEntry()
		generatedPasswordEntry.SetText(password)

		yourPasswordText := canvas.NewText("Your password:", color.Black)
		yourPasswordText.Alignment = fyne.TextAlignLeading
		yourPasswordText.TextStyle = fyne.TextStyle{Bold: true, Symbol: true}

		savePasswordButton := widget.NewButton("Save", func() {
			ui.store.Save(domain, password)
		})

		youPasswordLayout := container.New(layout.NewFormLayout(),
			yourPasswordText,
			generatedPasswordEntry,
		)

		grid.Add(container.NewVBox(youPasswordLayout, savePasswordButton))
		form.Refresh()
	}

	grid.Add(form)

	return grid
}
