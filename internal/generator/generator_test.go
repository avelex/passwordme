package generator

import (
	"log"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeneratePassword32Length(t *testing.T) {
	r := require.New(t)

	expected := "V*gN@qUb{5?Nx&n*!xxOWOZ8kAwhB'N7"

	pg := &PasswordGenerator{}
	url := &url.URL{
		Host: "github.com",
	}

	password := pg.Generate("test", url, nil, WithLength32())

	log.Println("your password", password)
	r.Equal(expected, password)
}

func TestGeneratePassword16Length(t *testing.T) {
	r := require.New(t)

	expected := "m,G8D~fY~&1c=8w/"

	pg := &PasswordGenerator{}
	url := &url.URL{
		Host: "github.com",
	}

	password := pg.Generate("test", url, nil, WithLength16())

	log.Println("your password", password)
	r.Equal(expected, password)
}
