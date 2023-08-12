GIT_VERSION = $(shell git describe --tags --abbrev=0)
APP_VERSION = $(GIT_VERSION:v%=%)

ICON = assets/img/Icon.png

APP_ID = com.password.me
APP_NAME = PasswordME
CLI_NAME = passwordme

INTERNAL_APP_PATH = github.com/avelex/passwordme/internal/app
BUILD_INFO = "-s -w -X '$(INTERNAL_APP_PATH).version=$(GIT_VERSION)' -X '$(INTERNAL_APP_PATH).name=$(APP_NAME)'"

localbuild:
	wails build -ldflags $(BUILD_INFO) -clean

linux:
	wails build -ldflags $(BUILD_INFO) -platform linux/amd64 -clean

windows:
	wails build -ldflags $(BUILD_INFO) -platform windows/amd64 -clean

cli:
	go build -o $(CLI_NAME) -ldflags="$(BUILD_INFO)" ./cmd/* 