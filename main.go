package main

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/avelex/passwordme/controller/ui"
	"github.com/avelex/passwordme/internal/generator"
	"github.com/avelex/passwordme/internal/store"
)

const _APP_DIRNAME = "passwordme"

var (
	//go:embed assets/img/Icon.png
	icon []byte
	//go:embed assets/img/logo.png
	logo []byte
	//go:embed assets/img/background.png
	background []byte
)

func main() {
	appDir, err := initAppDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	store := store.NewPasswordStore(appDir)
	generator := &generator.PasswordGenerator{}

	ui := ui.NewUI(generator, store, icon, logo, background)
	ui.Run()
}

func initAppDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(configDir, _APP_DIRNAME)

	if err := os.MkdirAll(appDir, os.ModeDir|0700); err != nil {
		return "", err
	}

	return appDir, nil
}
