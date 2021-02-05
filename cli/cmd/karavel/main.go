package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var Version = "unstable"

func main() {
	logger := log.New(os.Stderr, "", 0)

	app := &cli.App{
		Name:    "karavel",
		Usage:   "Smooth sailing in the Cloud sea",
		Version: Version,
		Commands: []*cli.Command{
			NewInitCommand(logger),
			NewRenderCommand(logger),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatalf("Error: %s", err)
	}
}
