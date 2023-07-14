package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/avelex/passwordme/controller/ui"
	"github.com/avelex/passwordme/internal/app"
	"github.com/avelex/passwordme/internal/generator"
	"github.com/avelex/passwordme/internal/store"
)

var (
	//go:embed assets/img/Icon.png
	icon []byte
	//go:embed assets/img/logo.png
	logo []byte
	//go:embed assets/img/background.png
	background []byte
)

func main() {
	appDir, err := app.CreateAppDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	store := store.NewPasswordStore(appDir)
	generator := &generator.PasswordGenerator{}

	app := app.NewApp(generator, store)

	ui := ui.NewUI(app, icon, logo, background)
	ui.Run()
}
