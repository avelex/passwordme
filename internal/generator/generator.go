package generator

import (
	"crypto/sha256"
	"math/big"
	"net/url"
	"strings"
	"unicode"
)

const (
	_DEFAULT_PASSWORD_LENGTH  = 12
	_DEFAULT_PASSWORD_SYMBOLS = true
	_DEFAULT_PASSWORD_NUMBERS = true

	_CHARS   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890~!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	_SYMBOLS = "~!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	_NUMBERS = "1234567890"
)

type PasswordGenerator struct {
}

type PasswordOpt func(p *Password)

func WithLength(length uint8) PasswordOpt {
	return func(p *Password) {
		p.Length = length
	}
}

type Password struct {
	WithNumbers       bool
	WithSpecialSymbol bool
	Length            uint8
	MasterPassword    string
	Host              string
	Promts            []string
}

func (p *PasswordGenerator) Generate(master string, domain *url.URL, prompts []string, opts ...PasswordOpt) string {
	password := &Password{
		WithNumbers:       _DEFAULT_PASSWORD_NUMBERS,
		WithSpecialSymbol: _DEFAULT_PASSWORD_SYMBOLS,
		Length:            _DEFAULT_PASSWORD_LENGTH,
		MasterPassword:    master,
		Host:              domain.Host,
		Promts:            prompts,
	}

	for _, opt := range opts {
		opt(password)
	}

	return generate(password)
}

func generate(password *Password) string {
	hash := sha256.New()
	hash.Write([]byte(password.MasterPassword))
	hash.Write([]byte(password.Host))
	for _, promt := range password.Promts {
		hash.Write([]byte(promt))
	}

	hashBytes := hash.Sum(nil)
	bigNum := new(big.Int).SetBytes(hashBytes)

	passwordBuilder := strings.Builder{}
	passwordBuilder.Grow(int(password.Length))

	add := big.NewInt(0)
	sum := new(big.Int).Add(add, bigNum)
	for i := 0; i < int(password.Length); i++ {
		sum = new(big.Int).Add(add, sum)
		cursor := new(big.Int).Mod(sum, big.NewInt(int64(len(_CHARS))))
		passwordBuilder.WriteByte(_CHARS[cursor.Int64()])
		add = new(big.Int).Add(add, cursor)
	}

	if password.WithNumbers {
		digitNotExists := false
		for _, v := range passwordBuilder.String() {
			if unicode.IsDigit(v) {
				digitNotExists = true
				break
			}
		}

		if !digitNotExists {
			cursor := new(big.Int).Mod(sum, big.NewInt(int64(len(_NUMBERS))))
			passwordBuilder.WriteByte(_NUMBERS[cursor.Int64()])
		}
	}

	if password.WithSpecialSymbol && !strings.ContainsAny(passwordBuilder.String(), _SYMBOLS) {
		cursor := new(big.Int).Mod(sum, big.NewInt(int64(len(_SYMBOLS))))
		passwordBuilder.WriteByte(_SYMBOLS[cursor.Int64()])
	}

	return passwordBuilder.String()
}
