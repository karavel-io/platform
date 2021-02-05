package main

import (
	"github.com/mikamai/karavel/cli/pkg/action"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"log"
	"path/filepath"
	"strings"
)

func NewInitCommand(logger *log.Logger) *cli.Command {
	var ver string
	var filename string
	var force bool
	var cfgUrl string
	var sumUrl string

	return &cli.Command{
		Name:      "init",
		Usage:     "Initialize a new Karavel project",
		ArgsUsage: "WORKDIR",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "version",
				Aliases:     []string{"v"},
				Usage:       "Karavel Platform version to initialize",
				Value:       "latest",
				Destination: &ver,
			},
			&cli.PathFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Karavel config file name to create",
				Value:       DefaultFileName,
				Destination: &filename,
			},
			&cli.BoolFlag{
				Name:        "force",
				Aliases:     []string{"f"},
				Usage:       "Overwrite the config file even if it already exists",
				Value:       false,
				Destination: &force,
			},
			&cli.PathFlag{
				Name:        "config-url",
				Usage:       "URL pointing to the Karavel config file to download. Requires setting --checksum-url too",
				DefaultText: "the official Karavel config file URL",
				Destination: &cfgUrl,
			},
			&cli.PathFlag{
				Name:        "checksum-url",
				Usage:       "URL pointing to the Karavel config file checksum to download. Requires setting --config-url too",
				DefaultText: "the official Karavel config file checksum URL",
				Destination: &sumUrl,
			},
		},
		Action: func(ctx *cli.Context) error {
			cwd := ctx.Args().Get(0)
			if cwd == "" {
				return errors.New("argument 'workdir' must be provided")
			}
			cwd, err := filepath.Abs(cwd)
			if err != nil {
				return err
			}

			ver = strings.TrimPrefix(ver, "v")

			if (cfgUrl != "" && sumUrl == "") || (sumUrl != "" && cfgUrl == "") {
				return errors.New("both --config-url and --checksum-url must be provided")
			}

			return action.Initialize(logger, action.InitParams{
				Workdir:         cwd,
				Filename:        filename,
				KaravelVersion:  ver,
				Force:           force,
				FileUrlOverride: cfgUrl,
				SumUrlOverride:  sumUrl,
			})
		},
	}
}
