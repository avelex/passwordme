linux:
	CGO_ENABLED=1 fyne package -release -os linux -icon assets/img/Icon.png -name PasswordME --appID com.password.me --appVersion 0.1.0
windows:
	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ fyne package -release -os windows -icon assets/img/Icon.png -name PasswordME --appID com.password.me --appVersion 0.1.0