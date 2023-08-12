package app

import "strings"

var (
	name    string = "PasswordME"
	version string = "dev"
)

func Name() string {
	return name
}

func CliName() string {
	return strings.ToLower(name)
}

func Version() string {
	return version
}
