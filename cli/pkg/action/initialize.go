package action

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/mikamai/karavel/cli/pkg/utils"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	latestReleaseURL = "https://github.com/mikamai/karavel/releases/latest/download"
	releaseUrl       = "https://github.com/mikamai/karavel/releases/platform-%s/download"
)

type InitParams struct {
	Workdir         string
	Filename        string
	KaravelVersion  string
	Force           bool
	FileUrlOverride string
	SumUrlOverride  string
}

func Initialize(logger *log.Logger, params InitParams) error {
	workdir := params.Workdir
	ver := params.KaravelVersion
	filename := params.Filename
	force := params.Force

	logger.Printf("Initializing new Karavel %s project at %s\n", ver, workdir)

	logger.Println()

	var url string
	if ver == "latest" {
		url = latestReleaseURL
	} else {
		url = fmt.Sprintf(releaseUrl, ver)
	}

	cfgUrl := params.FileUrlOverride
	if cfgUrl == "" {
		cfgUrl = path.Join(url, filename)
	}

	sumUrl := params.SumUrlOverride
	if sumUrl == "" {
		sumUrl = path.Join(url, filename+".sha256")
	}

	logger.Printf("Fetching bootstrap config from %s with checksum %s", cfgUrl, sumUrl)
	logger.Println()

	// TODO: implement init
	// mkdir -p workdir
	// check if karavel.cfg already exists
	// if exists and !force
	//   exit with error
	// elif exists and force
	//   overwrite with boilerplate
	// else
	//    write boilerplate

	if err := os.MkdirAll(workdir, 0755); err != nil {
		return err
	}

	filedst := filepath.Join(workdir, filename)
	info, err := os.Stat(filedst)
	if err != nil && err != os.ErrNotExist {
		return err
	}

	if info != nil && !force {
		return errors.Errorf("Karavel config file %s already exists", filename)
	}

	if info != nil && force {
		logger.Printf("WARNING: Karavel config file %s already exists and will be overwritten", filename)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	cfg, err := download(ctx, logger, cfgUrl)
	if err != nil {
		return err
	}
	logger.Println()

	shaHex, err := download(ctx, logger, sumUrl)
	if err != nil {
		return err
	}
	logger.Println()

	shaHex = bytes.TrimSpace(shaHex)
	shaD := make([]byte, hex.DecodedLen(len(shaHex)))
	if _, err := hex.Decode(shaD, shaHex); err != nil {
		return err
	}

	if len(shaD) != sha256.Size {
		return errors.Errorf("invalid SHA-256 checksum length: %d", len(shaD))
	}
	var sha [sha256.Size]byte
	copy(sha[:], shaD)

	sum := sha256.Sum256(cfg)
	if sum != sha {
		return errors.Errorf("checksum mismatch: wanted %x, got %x", sha, sum)
	}

	logger.Printf("Checksum successfully validated. Writing config file to %s", filedst)
	return ioutil.WriteFile(filedst, cfg, 0655)
}

func download(ctx context.Context, logger *log.Logger, url string) ([]byte, error) {
	f, err := ioutil.TempFile("", path.Base(url))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	defer os.Remove(f.Name())

	if err := utils.DownloadWithProgress(ctx, logger, url, f.Name()); err != nil {
		return nil, err
	}
	return ioutil.ReadFile(f.Name())
}
