PROJ_DIR     := $(CURDIR)
BINDIR       := $(PROJ_DIR)/bin
BINNAME      ?= karavel
INSTALL_PATH ?= /usr/local/bin
SHELL        = /usr/bin/env bash

SRC := $(shell find . -type f -name '*.go' -print) go.mod go.sum

GOBIN         = $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN         = $(shell go env GOPATH)/bin
endif
PKGS          = $(PROJ_DIR)/...

VERSION    = $(shell cat $(PROJ_DIR)/VERSION)
GIT_COMMIT = $(shell git rev-parse  HEAD)
GIT_DIRTY  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
BUILD_DATE = $(shell date +'%Y-%m-%d %H:%M:%S')

LDFLAGS    += -X 'github.com/mikamai/karavel/cli/internal/version.version=${VERSION}'
LDFLAGS    += -X 'github.com/mikamai/karavel/cli/internal/version.gitCommit=${GIT_COMMIT}'
LDFLAGS    += -X 'github.com/mikamai/karavel/cli/internal/version.gitTreeState=${GIT_DIRTY}'
LDFLAGS    += -X 'github.com/mikamai/karavel/cli/internal/version.buildDate=${BUILD_DATE}'
LDFLAGS    += $(EXT_LDFLAGS)

.PHONY: all
all: build

.PHONY: addlicense
addlicense:
	addlicense -c "MIKAMAI s.r.l" -l apache .

.PHONY: build
build: fmt vet $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	GO111MODULE=on go build -ldflags "$(LDFLAGS)" -o $(BINDIR)/$(BINNAME) $(PROJ_DIR)/cmd/karavel

.PHONY: install
install: build
	@install $(BINDIR)/$(BINNAME) $(INSTALL_PATH)/$(BINNAME)

.PHONY: clean
clean:
	rm -rf $(BINDIR)

.PHONY: test
test:
	go test $(PKGS)

.PHONY: fmt
fmt: addlicense
	go fmt $(PKGS)

.PHONY: vet
vet:
	go vet $(PKGS)
