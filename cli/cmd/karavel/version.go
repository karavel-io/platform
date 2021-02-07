package main

import (
	"fmt"
	"github.com/mikamai/karavel/cli/internal/version"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints the CLI version and exits",
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Printf("%+v", version.Get())
			return err
		},
	}
}
