package main

import (
	"fmt"
	"github.com/mikamai/karavel/cli/pkg/action"
	"github.com/mikamai/karavel/cli/pkg/logger"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

const DefaultFileName = "karavel.hcl"

func NewRenderCommand(log logger.Logger) *cobra.Command {
	var cpath string
	cmd := &cobra.Command{
		Use:   "render",
		Short: "Render a Karavel project",
		Long: fmt.Sprintf(`
Render a Karavel project with the given config (defaults to '%s' in the current directory).

This command is idempotent and can be run multiple times without issues. 
It will respect changes made to files outside the 'vendor' directory, only adding or removing Karavel-specific entries.
It will, however, consider the 'vendor' directory as a fully-managed folder and may add, delete or modify any file inside it without warning.
`, DefaultFileName),
		RunE: func(cmd *cobra.Command, args []string) error {
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

	cmd.Flags().StringVarP(&cpath, "file", "f", DefaultFileName, "Specify an alternate config file")

	return cmd
}
