package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/avelex/passwordme/internal/app"
	"github.com/urfave/cli/v2"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cli := &cli.App{
		Name:    app.CliName(),
		Usage:   "cross-platform password manager",
		Version: app.Version(),
		Action:  cli.ShowAppHelp,
		Commands: []*cli.Command{
			generatePasswordCommand,
		},
	}

	if err := cli.RunContext(ctx, os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
