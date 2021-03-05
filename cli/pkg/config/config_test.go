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

package config

import (
	"github.com/mikamai/karavel/cli/internal/helmw"
	"github.com/stretchr/testify/suite"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type ConfigTestSuite struct {
	suite.Suite
	logw      io.Writer
	stableCfg *os.File
	edgeCfg   *os.File
}

func (*ConfigTestSuite) prepareConfig(cfg string) *os.File {
	f, err := ioutil.TempFile("", "karavel")
	if err != nil {
		panic(f)
	}

	if err := ioutil.WriteFile(f.Name(), []byte(cfg), 0600); err != nil {
		panic(err)
	}

	return f
}

func (s *ConfigTestSuite) SetupSuite() {
	s.logw = ioutil.Discard

	s.stableCfg = s.prepareConfig(`
component "test" {
	version = "0.1.0"
	namespace = "test"

	hello = "world"
}
`)

	s.edgeCfg = s.prepareConfig(`
channel = "edge"

component "test" {
	version = "0.1.0-alpha"
	namespace = "test"

	hello = "world"
}
`)
}

func (s *ConfigTestSuite) TearDownSuite() {
	os.Remove(s.stableCfg.Name())
	os.Remove(s.edgeCfg.Name())
}

func (s *ConfigTestSuite) TestStable() {
	assert := s.Assert()
	f := s.stableCfg

	cfg, err := ReadFrom(s.logw, f.Name())
	if err != nil {
		assert.NoError(err)
	}

	assert.Equal(ChannelStable, cfg.Channel)
	assert.Equal(helmw.HelmDefaultStableRepo, cfg.HelmRepoUrl)
	assert.Equal(1, len(cfg.Components))

	c := cfg.Components[0]
	assert.Equal("test", c.Name)
	assert.Equal("0.1.0", c.Version)
	assert.Equal("test", c.Namespace)
	assert.Equal("{\"hello\":\"world\"}", c.JsonParams)
}

func (s *ConfigTestSuite) TestEdge() {
	assert := s.Assert()
	f := s.edgeCfg

	cfg, err := ReadFrom(s.logw, f.Name())
	if err != nil {
		assert.NoError(err)
	}

	assert.Equal(ChannelEdge, cfg.Channel)
	assert.Equal(helmw.HelmDefaultEdgeRepo, cfg.HelmRepoUrl)
	assert.Equal(1, len(cfg.Components))

	c := cfg.Components[0]
	assert.Equal("test", c.Name)
	assert.Equal("0.1.0-alpha", c.Version)
	assert.Equal("test", c.Namespace)
	assert.Equal("{\"hello\":\"world\"}", c.JsonParams)
}

func TestReadFrom(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
