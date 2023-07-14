GIT_VERSION = $(shell git describe --tags --abbrev=0)
APP_VERSION = $(GIT_VERSION:v%=%)

ICON = assets/img/Icon.png

APP_ID = com.password.me
APP_NAME = PasswordME
CLI_NAME = passwordme

INTERNAL_APP_PATH = github.com/avelex/passwordme/internal/app
BUILD_INFO = -s -w -X '$(INTERNAL_APP_PATH).version=$(GIT_VERSION)' -X '$(INTERNAL_APP_PATH).name=$(APP_NAME)'

build:
	CGO_ENABLED=1 fyne package -release -icon $(ICON) -name $(APP_NAME) --appID $(APP_ID) --appVersion $(APP_VERSION)

linux:
	CGO_ENABLED=1 fyne package -release -os linux -icon $(ICON) -name $(APP_NAME) --appID $(APP_ID) --appVersion $(APP_VERSION)

windows:
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ fyne package -release -os windows -icon $(ICON) -name $(APP_NAME) --appID $(APP_ID) --appVersion $(APP_VERSION)

cli:
	go build -o $(CLI_NAME) -ldflags="$(BUILD_INFO)" ./cmd/* 