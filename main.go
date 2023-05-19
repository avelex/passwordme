package main

import (
	_ "embed"

	"github.com/avelex/passwordme/controller/ui"
	"github.com/avelex/passwordme/internal/generator"
)

const _APP_NAME = "PasswordME"

var (
	//go:embed assets/img/Icon.png
	icon []byte
	//go:embed assets/img/logo.png
	logo []byte
	//go:embed assets/img/background.png
	background []byte
)

func main() {
	generator := &generator.PasswordGenerator{}
	ui := ui.NewUI(generator, icon, logo, background)
	ui.Run()
}
