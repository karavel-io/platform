package main

import (
	"fmt"
	"github.com/mikamai/karavel/cli/pkg/action"
	"github.com/mikamai/karavel/cli/pkg/logger"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
)

const DefaultFileName = "karavel.hcl"

func NewRenderCommand(log logger.Logger) *cli.Command {
	var debug bool
	var quiet bool
	var cpath string
	return &cli.Command{
		Name:  "render",
		Usage: "Render a Karavel project",
		Description: fmt.Sprintf(`
Render a Karavel project with the given config (defaults to '%s' in the current directory).

This command is idempotent and can be run multiple times without issues. 
It will respect changes made to files outside the 'vendor' directory, only adding or removing Karavel-specific entries.
It will, however, consider the 'vendor' directory as a fully-managed folder and may add, delete or modify any file inside it without warning.
`, DefaultFileName),
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Usage:       "Specify an alternate config file",
				Value:       DefaultFileName,
				Destination: &cpath,
				EnvVars:     []string{"KARAVEL_CONFIG_FILE"},
			},
			flagDebug(&debug),
			flagQuiet(&quiet),
		},
		Action: func(ctx *cli.Context) error {
			if debug {
				log.SetLevel(logger.LvlDebug)
			}

			if quiet {
				log.SetLevel(logger.LvlError)
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

			return action.Render(log, action.RenderParams{
				ConfigPath: cpath,
			})
		},
	}
}
