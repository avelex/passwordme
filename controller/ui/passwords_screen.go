package ui

import (
	"errors"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (ui *uiController) passwordsScreen() Screen {
	return Screen{
		Title: _PASSWORD_REPOSITORY,
		View:  ui.passwordsView,
	}
}

func (ui *uiController) passwordsView(w fyne.Window) fyne.CanvasObject {
	if !ui.app.PassfileExists() {
		passwordEntry := widget.NewPasswordEntry()
		passwordEntry.Validator = func(s string) error {
			if strings.TrimSpace(s) == "" {
				return errors.New("password required")
			}

			return nil
		}
		createStoreButton := widget.NewButton("Create", nil)

		content := container.NewVBox(
			container.NewCenter(
				widget.NewLabel("You don't have passwords store"),
			),
			layout.NewSpacer(),
			container.NewCenter(
				container.NewAdaptiveGrid(2,
					container.NewCenter(
						widget.NewLabel("Master Password:"),
					),
					passwordEntry,
					widget.NewButton("Import", func() {}),
					createStoreButton,
				),
			),
			layout.NewSpacer(),
		)

		createStoreButton.OnTapped = func() {
			if err := ui.app.CreatePassfile(passwordEntry.Text); err != nil {
				return
			}
			content.Objects = []fyne.CanvasObject{ui.passowrdList(w)}
			content.Refresh()
		}

		return content
	}

	if !ui.app.PassfileOpened() {
		content := container.NewMax(
			container.NewCenter(widget.NewLabel("Locked")),
		)

		masterPasswordEntry := widget.NewPasswordEntry()
		dial := dialog.NewForm(
			"Write Your Master Password",
			"Unlock", "Cancel",
			[]*widget.FormItem{
				widget.NewFormItem("Password", masterPasswordEntry),
			},
			func(unlock bool) {
				if !unlock {
					return
				}

				if err := ui.app.OpenPassfile(masterPasswordEntry.Text); err != nil {
					return
				}

				content.Objects = []fyne.CanvasObject{ui.passowrdList(w)}
				content.Refresh()
			},
			w,
		)

		dial.Resize(fyne.NewSize(250, 150))
		dial.Show()

		return content
	} else {
		return ui.passowrdList(w)
	}
}

func (ui *uiController) passowrdList(w fyne.Window) fyne.CanvasObject {
	passwords, _ := ui.app.ListPasswords()

	list := widget.NewList(
		func() int { return len(passwords) },
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(passwords[id])
		},
	)

	list.OnSelected = func(id widget.ListItemID) {
		password, err := ui.app.ShowPassword(passwords[id])
		if err != nil {
			log.Println(err)
			return
		}

		generatedPasswordEntry := widget.NewEntry()
		generatedPasswordEntry.SetText(password)

		d := dialog.NewCustom("Password", "OK", generatedPasswordEntry, w)
		d.Resize(fyne.NewSize(250, 150))
		d.Show()
	}

	return list
}
