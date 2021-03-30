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

package gitutils

import (
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/mikamai/karavel/cli/pkg/logger"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func GetOriginRemote(log logger.Logger, cwd string, urlOverride string) (string, string, error) {
	path, err := findGitRepoPath(log, cwd)
	if err != nil {
		return "", "", err
	}

	fs := osfs.New(path)
	if _, err := fs.Stat(git.GitDirName); err == nil {
		fs, err = fs.Chroot(git.GitDirName)
		if err != nil {
			return "", "", err
		}
	}

	s := filesystem.NewStorageWithOptions(fs, cache.NewObjectLRUDefault(), filesystem.Options{KeepDescriptors: true})
	r, err := git.Open(s, fs)
	if err != nil {
		return "", "", err
	}

	// remove the .git step
	path = filepath.Dir(path)
	if urlOverride != "" {
		return path, urlOverride, nil
	}

	rems, err := r.Remotes()
	if err != nil {
		return "", "", err
	}

	if len(rems) == 0 {
		return "", "", errors.Errorf("no remote found for git repo at %s", path)
	}

	if len(rems) == 1 {
		rem := rems[0].Config()
		urls := rem.URLs
		if len(urls) > 0 {
			return path, urls[0], nil
		}

		return "", "", errors.Errorf("git remote %s has no URL defined", rem.Name)
	}

	for _, rem := range rems {
		c := rem.Config()
		if c.Name == "origin" {
			urls := c.URLs
			if len(urls) > 0 {
				return path, urls[0], nil
			}

			return "", "", errors.Errorf("git remote %s has no URL defined", c.Name)
		}
	}
	return "", "", errors.Errorf("could not find a suitable remote for git repo %s", path)
}

func findGitRepoPath(log logger.Logger, path string) (string, error) {
	log.Debugf("looking for repo for %s", path)

	gitPath := filepath.Join(path, git.GitDirName)
	i, err := os.Stat(gitPath)
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}

	if os.IsNotExist(err) {
		up := filepath.Dir(path)
		log.Debugf("no git dir found, going up one notch from %s to %s", path, up)
		if up == "" || up == "." || up == "/" {
			return "", errors.New("reached top of fs without finding a git repository")
		}
		return findGitRepoPath(log, up)
	}

	if !i.IsDir() || i.Name() != git.GitDirName {
		var d string
		if i.IsDir() {
			d = " not"
		}

		return "", errors.Errorf("failed to find git repo, found %s which is%s a directory", i.Name(), d)
	}

	log.Debugf("found git repo at %s", gitPath)
	return gitPath, nil
}
