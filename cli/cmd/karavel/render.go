package main

import (
	"fmt"
	"github.com/mikamai/karavel/cli/pkg/action"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
)

const DefaultFileName = "karavel.hcl"

func NewRenderCommand(logger *log.Logger) *cli.Command {
	debug := false
	return &cli.Command{
		Name:  "render",
		Usage: "Render a new Karavel project",
		Description: `
Render a new Karavel project with the given config (defaults to 'karavel.hcl' in the current directory). 
`,
		ArgsUsage: "[CONFIG]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Usage:       "Output debug logs",
				Destination: &debug,
			},
		},
		Action: func(ctx *cli.Context) error {
			cpath := ctx.Args().Get(0)
			if cpath == "" {
				cpath = DefaultFileName
			}
			cpath, err := filepath.Abs(cpath)
			if err != nil {
				return err
			}

			cstat, err := os.Stat(cpath)
			if err != nil {
				return err
			}

			if cstat.IsDir() && cstat.Name() != DefaultFileName {
				cpath = filepath.Join(cpath, DefaultFileName)
			}

			cstat, err = os.Stat(cpath)
			if err != nil {
				return err
			}

			if cstat.IsDir() {
				return fmt.Errorf("invalid config file %s, is a directory", cpath)
			}

			return action.Render(logger, action.RenderParams{
				ConfigPath: cpath,
				Debug:      debug,
			})
		},
	}
}
