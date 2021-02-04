package action

import (
	"fmt"
	"log"
)

const (
	latestReleaseURL = "https://github.com/mikamai/karavel/releases/latest/download/karavel.cfg"
	releaseUrl       = "https://github.com/mikamai/karavel/releases/platform-%s/download/karavel.cfg"
)

type InitParams struct {
	Workdir        string
	KaravelVersion string
}

func Initialize(logger *log.Logger, params InitParams) error {
	workdir := params.Workdir
	ver := params.KaravelVersion
	logger.Printf("Initializing new Karavel %s project at %s\n", ver, workdir)

	logger.Println()

	var url string
	if ver == "latest" {
		url = latestReleaseURL
	} else {
		url = fmt.Sprintf(releaseUrl, ver)
	}

	logger.Printf("Fetching bootstrap config from %s", url)

	// TODO: implement init
	// mkdir -p workdir
	// check if karavel.cfg already exists
	// if exists and !force
	//   exit with error
	// elif exists and force
	//   overwrite with boilerplate
	// else
	//    write boilerplate

	return nil
}
