package main

import (
	"github.com/mikamai/karavel/cli/internal/version"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "", 0)

	app := &cli.App{
		Name:    "karavel",
		Usage:   "Smooth sailing in the Cloud sea",
		Version: version.Short(),
		Commands: []*cli.Command{
			NewInitCommand(logger),
			NewRenderCommand(logger),
			NewVersionCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatalf("Error: %s", err)
	}
}
