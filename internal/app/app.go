package app

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/avelex/passwordme/internal/generator"
	"github.com/avelex/passwordme/internal/store"
)

type PasswordMeUsecases interface {
	GeneratePassword(master string, domain *url.URL, prompts []string, opts ...generator.PasswordOpt) string
	CreatePassfile(masterPassword string) error
	DeletePassfile() error
	ImportPassfile() error
	OpenPassfile(masterPassword string) error
	PassfileOpened() bool
	PassfileExists() bool
	ListPasswords() ([]string, error)
	SavePassword(name string, password string) error
	ShowPassword(name string) (string, error)
}

type app struct {
	generator *generator.PasswordGenerator
	store     *store.PasswordStore
}

func CreateAppDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(configDir, strings.ToLower(name))

	if err := os.MkdirAll(appDir, os.ModeDir|0700); err != nil {
		return "", err
	}

	return appDir, nil
}

func NewApp(gen *generator.PasswordGenerator, store *store.PasswordStore) *app {
	return &app{
		generator: gen,
		store:     store,
	}
}

func (a *app) GeneratePassword(master, domain, length string, prompts []string) string {
	url := &url.URL{
		Host: domain,
	}

	opt := generator.WithLength16()
	switch length {
	case "32":
		opt = generator.WithLength32()
	}

	return a.generator.Generate(master, url, prompts, opt)
}

func (a *app) CreatePassfile(masterPassword string) error {
	return a.store.Create(masterPassword)
}

func (a *app) DeletePassfile() error {
	return a.store.Delete()
}

func (a *app) ImportPassfile() error {
	return a.store.Import()
}

func (a *app) OpenPassfile(masterPassword string) error {
	return a.store.Open(masterPassword)
}

func (a *app) PassfileOpened() bool {
	return a.store.Opened()
}

func (a *app) PassfileExists() bool {
	return a.store.Exists()
}

func (a *app) ListPasswords() ([]string, error) {
	return a.store.List()
}

func (a *app) SavePassword(name string, password string) error {
	return a.store.Save(name, password)
}

func (a *app) ShowPassword(name string) (string, error) {
	return a.store.ShowPassword(name)
}

func (a *app) Version() string {
	return Version()
}
