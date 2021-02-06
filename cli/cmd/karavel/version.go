package main

import (
	"fmt"
	"github.com/mikamai/karavel/cli/internal/version"
	"github.com/urfave/cli/v2"
)

func NewVersionCommand() *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "Prints the CLI version and exits",
		Action: func(ctx *cli.Context) error {
			_, err := fmt.Printf("%+v", version.Get())
			return err
		},
	}
}
