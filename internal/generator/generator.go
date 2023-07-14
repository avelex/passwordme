package generator

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"net/url"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	_DEFAULT_PASSWORD_LENGTH  = 32
	_DEFAULT_PASSWORD_SYMBOLS = true
	_DEFAULT_PASSWORD_NUMBERS = true

	_CHARS   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890~!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	_SYMBOLS = "~!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
	_NUMBERS = "1234567890"
)

type PasswordGenerator struct {
}

type PasswordOpt func(p *password)

func WithLength32() PasswordOpt {
	return func(p *password) {
		p.Length = 32
	}
}

func WithLength16() PasswordOpt {
	return func(p *password) {
		p.Length = 16
	}
}

type password struct {
	WithNumbers       bool
	WithSpecialSymbol bool
	Length            uint8
	MasterPassword    string
	Host              string
	Prompts           []string
}

func (p *PasswordGenerator) Generate(master string, domain *url.URL, prompts []string, opts ...PasswordOpt) string {
	password := &password{
		WithNumbers:       _DEFAULT_PASSWORD_NUMBERS,
		WithSpecialSymbol: _DEFAULT_PASSWORD_SYMBOLS,
		Length:            _DEFAULT_PASSWORD_LENGTH,
		MasterPassword:    master,
		Host:              domain.Hostname(),
		Prompts:           prompts,
	}

	for _, opt := range opts {
		opt(password)
	}

	return generate(password)
}

func generate(password *password) string {
	hash := sha256.New()
	hash.Write([]byte(password.MasterPassword))
	hash.Write([]byte(password.Host))
	for _, prompt := range password.Prompts {
		hash.Write([]byte(prompt))
	}

	hashBytes := hash.Sum(nil)

	passwordNumberHex := hex.EncodeToString(hashBytes)

	passwordBuilder := strings.Builder{}
	passwordBuilder.Grow(int(password.Length))

	shift := len(passwordNumberHex) / int(password.Length)

	for i := 0; i < len(passwordNumberHex); i += shift {
		end := i + shift
		if end > len(passwordNumberHex) {
			end = len(passwordNumberHex)
		}
		num := mustParseInt(passwordNumberHex[i:end], 16, 64)
		cursor := int(num) % len(_CHARS)
		passwordBuilder.WriteByte(_CHARS[cursor])
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
			passwordNumber := new(big.Int).SetBytes(hashBytes)
			passwordBuilder.WriteByte(chooseChar(passwordNumber, _NUMBERS))
		}
	}

	if password.WithSpecialSymbol && !strings.ContainsAny(passwordBuilder.String(), _SYMBOLS) {
		passwordNumber := new(big.Int).SetBytes(hashBytes)
		passwordBuilder.WriteByte(chooseChar(passwordNumber, _SYMBOLS))
	}

	return passwordBuilder.String()
}

func chooseChar(passwordNumber *big.Int, charset string) byte {
	mod := new(big.Int).Mod(passwordNumber, big.NewInt(int64(utf8.RuneCountInString(charset))))
	return charset[mod.Int64()]
}

func mustParseInt(s string, base, bitSize int) int64 {
	num, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		panic(err)
	}
	return num
}
