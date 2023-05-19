package generator

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGenerate(t *testing.T) {
	pg := PasswordGenerator{}
	url, _ := url.Parse("https://github.com")
	password := pg.Generate("my_usually_password", url, []string{"avelex"}, WithLength(12))
	fmt.Printf("your fantastic password: %v\n", password)
}
