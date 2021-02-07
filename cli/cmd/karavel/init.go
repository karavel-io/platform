package main

import (
	"github.com/mikamai/karavel/cli/pkg/action"
	"github.com/mikamai/karavel/cli/pkg/logger"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

func NewInitCommand(log logger.Logger) *cobra.Command {
	var ver string
	var filename string
	var force bool
	var cfgUrl string
	var sumUrl string

	cmd := &cobra.Command{
		Use:   "init [WORKDIR]",
		Short: "Initialize a new Karavel project",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var cwd string
			if len(args) > 0 {
				cwd = args[0]
			}

			if cwd == "" {
				d, err := os.Getwd()
				if err != nil {
					return err
				}
				cwd = d
			}
			cwd, err := filepath.Abs(cwd)
			if err != nil {
				return err
			}

			ver = strings.TrimPrefix(ver, "v")

			if (cfgUrl != "" && sumUrl == "") || (sumUrl != "" && cfgUrl == "") {
				return errors.New("both --config-url and --checksum-url must be provided")
			}

			return action.Initialize(log, action.InitParams{
				Workdir:         cwd,
				Filename:        filename,
				KaravelVersion:  ver,
				Force:           force,
				FileUrlOverride: cfgUrl,
				SumUrlOverride:  sumUrl,
			})
		},
	}

	cmd.Flags().StringVarP(&ver, "version", "v", "latest", "Karavel Platform version to initialize")
	cmd.Flags().StringVarP(&filename, "output-file", "o", DefaultFileName, "Karavel config file name to create")
	cmd.Flags().BoolVar(&force, "force", false, "Overwrite the config file even if it already exists")
	cmd.Flags().StringVar(&cfgUrl, "config-url", "", "Override the official URL pointing to the Karavel config file to download. Requires setting --checksum-url too")
	cmd.Flags().StringVar(&sumUrl, "checksum-url", "", "Override the official URL pointing to the Karavel config file checksum to download. Requires setting --config-url too")

	return cmd
}
