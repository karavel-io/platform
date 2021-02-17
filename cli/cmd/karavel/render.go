// Copyright 2021 MIKAMAI s.r.l
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	var skipGit bool

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
				SkipGit:    skipGit,
			})
		},
	}

	cmd.Flags().StringVarP(&cpath, "file", "f", DefaultFileName, "Specify an alternate config file")
	cmd.Flags().BoolVar(&skipGit, "skip-git", false, "Skip the git integration to discover remote repositories for Argo. WARNING: this will render the Argo component inoperable")

	return cmd
}
