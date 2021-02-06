package version

import (
	"runtime"
)

var (
	version      = ""
	gitCommit    = ""
	gitTreeState = ""
	buildDate    = ""
)

type Version struct {
	Version      string
	GitCommit    string
	GitTreeState string
	BuildDate    string
	Arch         string
	Os           string
	GoVersion    string
}

func Get() Version {
	return Version{
		Version:      version,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		Arch:         runtime.GOARCH,
		Os:           runtime.GOOS,
		GoVersion:    runtime.Version(),
	}
}

func Short() string {
	ver := version
	if gitTreeState != "clean" {
		ver += "-" + gitTreeState
	}
	return ver
}
