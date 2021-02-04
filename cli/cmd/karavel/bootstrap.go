package main

import (
	"fmt"
	"github.com/mikamai/karavel/cli/pkg/actions"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
)

const DefaultFileName = "karavel.cfg"

func NewBootstrapCommand(logger *log.Logger) *cli.Command {
	return &cli.Command{
		Name:  "bootstrap",
		Usage: "Bootstrap a new Karavel project",
		Description: `
Bootstrap a new Karavel project with the given config (defaults to 'karavel.cfg' in the current directory). 
`,
		ArgsUsage: "[CONFIG]",
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

			return actions.Bootstrap(logger, actions.BootstrapParams{
				ConfigPath: cpath,
			})
		},
	}
}
