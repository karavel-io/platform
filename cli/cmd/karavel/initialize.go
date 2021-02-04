package main

import (
	"fmt"
	"github.com/mikamai/karavel/cli/pkg/action"
	"github.com/urfave/cli/v2"
	"log"
	"path/filepath"
	"strings"
)

func NewInitCommand(logger *log.Logger) *cli.Command {
	var ver string

	return &cli.Command{
		Name:      "initialize",
		Aliases:   []string{"init"},
		Usage:     "Initialize a new Karavel project",
		ArgsUsage: "WORKDIR",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "karavel-version",
				Aliases:     []string{"kv"},
				Usage:       "Karavel Platform version to initialize",
				Value:       "latest",
				Destination: &ver,
			},
		},
		Action: func(ctx *cli.Context) error {
			cwd := ctx.Args().Get(0)
			if cwd == "" {
				return fmt.Errorf("argument 'workdir' must be provided")
			}
			cwd, err := filepath.Abs(cwd)
			if err != nil {
				return err
			}

			ver = strings.TrimPrefix(ver, "v")

			return action.Initialize(logger, action.InitParams{
				Workdir:        cwd,
				KaravelVersion: ver,
			})
		},
	}
}
