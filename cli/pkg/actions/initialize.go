package actions

import (
	"log"
)

type InitParams struct {
	Workdir        string
	KaravelVersion string
}

func Initialize(logger *log.Logger, params InitParams) error {
	logger.Printf("Initializing new Karavel %s project at %s\n", params.KaravelVersion, params.Workdir)

	// mkdir -p workdir
	// check if karavel.hcl already exists
	// if exists and !force
	//   exit with error
	// elif exists and force
	//   overwrite with boilerplate
	// else
	//    write boilerplate

	return nil
}
