package main

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/avelex/passwordme/internal/generator"
	"github.com/go-playground/validator"
	"github.com/urfave/cli/v2"
)

var generatePasswordCommand = &cli.Command{
	Name:      "generate",
	Usage:     "On-Flight generation password without saving",
	UsageText: "passwordme generate [--len=<16|32>] <master> <domain> [promts...]",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:        "len",
			DefaultText: "32",
		},
	},
	Action: generatePassword,
}

func generatePassword(ctx *cli.Context) error {
	gen := generator.PasswordGenerator{}

	if ctx.Args().Len() < 2 {
		return errors.New("generate args must contain <master> and <domain>")
	}

	master := ctx.Args().Get(0)

	domain := ctx.Args().Get(1)
	if !isValidDomain(domain) {
		return fmt.Errorf("invalid domain: %v", domain)
	}

	url := &url.URL{
		Host: domain,
	}

	promts := make([]string, 0)

	for i := 2; i < ctx.Args().Len(); i++ {
		if promt := ctx.Args().Get(i); promt != "" {
			promts = append(promts, promt)
		}
	}

	opt := generator.WithLength32()
	if len := ctx.Int("len"); len != 0 {
		switch len {
		case 16:
			opt = generator.WithLength16()
		case 32:
			opt = generator.WithLength32()
		default:
			return errors.New("invalid password len, must be 16 or 32")
		}
	}

	password := gen.Generate(master, url, promts, opt)

	fmt.Println(password)

	return nil
}

func isValidDomain(str string) bool {
	return validator.New().Var(str, "hostname") == nil
}
