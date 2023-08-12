package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/avelex/passwordme/internal/app"
	"github.com/avelex/passwordme/internal/generator"
	"github.com/avelex/passwordme/internal/store"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

var (
	//go:embed assets/img/Icon.png
	icon []byte
	//go:embed all:frontend/dist
	assets embed.FS
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

	err = wails.Run(&options.App{
		Title:         "PasswordME",
		Width:         800,
		Height:        600,
		MinWidth:      800,
		MinHeight:     600,
		MaxWidth:      800,
		MaxHeight:     600,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []interface{}{
			app,
		},
		Linux: &linux.Options{
			Icon: icon,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
