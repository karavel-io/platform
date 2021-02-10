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

package action

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/mikamai/karavel/cli/internal/utils"
	"github.com/mikamai/karavel/cli/pkg/logger"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	latestReleaseURL = "https://github.com/mikamai/karavel/releases/latest/download"
	releaseUrl       = "https://github.com/mikamai/karavel/releases/v%s/download"
)

type InitParams struct {
	Workdir         string
	Filename        string
	KaravelVersion  string
	Force           bool
	FileUrlOverride string
	SumUrlOverride  string
}

func Initialize(log logger.Logger, params InitParams) error {
	workdir := params.Workdir
	ver := params.KaravelVersion
	filename := params.Filename
	force := params.Force

	log.Infof("Initializing new Karavel v%s project at %s", ver, workdir)
	log.Info()

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

	log.Infof("Fetching bootstrap config from %s with checksum %s", cfgUrl, sumUrl)
	log.Info()

	if err := os.MkdirAll(workdir, 0755); err != nil {
		return err
	}

	filedst := filepath.Join(workdir, filename)
	info, err := os.Stat(filedst)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if info != nil && !force {
		return errors.Errorf("Karavel config file %s already exists", filename)
	}

	if info != nil && force {
		log.Warnf("Karavel config file %s already exists and will be overwritten", filename)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	cfg, err := download(ctx, log, cfgUrl)
	if err != nil {
		return err
	}
	log.Info()

	shaHex, err := download(ctx, log, sumUrl)
	if err != nil {
		return err
	}
	log.Info()

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

	log.Infof("Checksum successfully validated. Writing config file to %s", filedst)
	return ioutil.WriteFile(filedst, cfg, 0655)
}

func download(ctx context.Context, log logger.Logger, url string) ([]byte, error) {
	f, err := ioutil.TempFile("", path.Base(url))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	defer os.Remove(f.Name())

	if err := utils.DownloadWithProgress(ctx, log, url, f.Name()); err != nil {
		return nil, err
	}
	return ioutil.ReadFile(f.Name())
}
