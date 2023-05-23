GIT_VERSION = $(shell git describe --tags $(git rev-list --tags --max-count=1))
APP_VERSION = $(GIT_VERSION:v%=%)
ICON = assets/img/Icon.png

linux:
	CGO_ENABLED=1 fyne package -release -os linux -icon $(ICON) -name PasswordME --appID com.password.me --appVersion $(APP_VERSION)
windows:
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ fyne package -release -os windows -icon $(ICON) -name PasswordME --appID com.password.me --appVersion $(APP_VERSION)