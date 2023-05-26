package generator

import (
	"log"
	"net/url"
	"testing"
)

func TestGenerate(t *testing.T) {
	pg := &PasswordGenerator{}
	url := &url.URL{
		Host: "github.com",
	}

	password := pg.Generate("test", url, nil)
	log.Println("your password", password)
}
