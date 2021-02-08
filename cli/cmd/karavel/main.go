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
	"context"
	"github.com/mikamai/karavel/cli/internal/version"
	"github.com/mikamai/karavel/cli/pkg/logger"
	"github.com/spf13/cobra"
	"time"
)

func main() {
	log := logger.New(logger.LvlInfo)
	var debug bool
	var quiet bool
	var colors bool

	app := cobra.Command{
		Use:     "karavel",
		Short:   "Smooth sailing in the Cloud sea",
		Long:    ``,
		Version: version.Short(),
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			log.SetColors(colors)
			if debug {
				log.SetLevel(logger.LvlDebug)
			}
			if quiet {
				log.SetLevel(logger.LvlError)
			}
		},
	}

	app.PersistentFlags().BoolVar(&debug, "debug", false, "Output debug logs")
	app.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "Suppress all logs except errors")
	app.PersistentFlags().BoolVar(&colors, "colors", true, "Enable colored logs")

	app.AddCommand(NewInitCommand(log))
	app.AddCommand(NewRenderCommand(log))
	app.AddCommand(NewVersionCommand())

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	if err := app.ExecuteContext(ctx); err != nil {
		log.Fatal(err)
	}
}
