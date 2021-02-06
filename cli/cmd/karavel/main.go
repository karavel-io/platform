package main

import (
	"github.com/mikamai/karavel/cli/internal/version"
	"github.com/mikamai/karavel/cli/pkg/logger"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	flagDebug = func(i *bool) cli.Flag {
		return &cli.BoolFlag{
			Name:        "debug",
			Aliases:     []string{"d"},
			Usage:       "Output debug logs",
			EnvVars:     []string{"KARAVEL_DEBUG"},
			Destination: i,
		}
	}

	flagQuiet = func(i *bool) cli.Flag {
		return &cli.BoolFlag{
			Name:        "quiet",
			Aliases:     []string{"q"},
			Usage:       "Suppress logs except errors",
			EnvVars:     []string{"KARAVEL_QUIET"},
			Destination: i,
		}
	}
)

func main() {
	log := logger.New(logger.LvlInfo)

	app := &cli.App{
		Name:    "karavel",
		Usage:   "Smooth sailing in the Cloud sea",
		Version: version.Short(),
		Commands: []*cli.Command{
			NewInitCommand(log),
			NewRenderCommand(log),
			NewVersionCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
