package actions

import (
	"github.com/mikamai/karavel/cli/pkg/config"
	"log"
)

type BootstrapParams struct {
	ConfigPath string
}

func Bootstrap(logger *log.Logger, params BootstrapParams) error {
	cpath := params.ConfigPath
	logger.Printf("Bootstrapping new Karavel project with config file %s\n", cpath)

	cfg, err := config.ReadFrom(logger.Writer(), cpath)
	if err != nil {
		return err
	}

	if err := cfg.SetupHelm(); err != nil {
		return err
	}

	logger.Printf("Config: %+v", cfg)

	logger.Println()
	for _, cc := range cfg.Components {
		logger.Printf("Scaffolding component '%s' v%s", cc.Name, cc.Version)
		if err := cc.Render(logger); err != nil {
			return err
		}
	}

	return nil
}
