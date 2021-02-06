BINDIR  := $(CURDIR)/bin
BINNAME ?= karavel
INSTALL_PATH ?= /usr/local/bin
SHELL   = /usr/bin/env bash

SRC := $(shell find . -type f -name '*.go' -print) go.mod go.sum

GOBIN         = $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN         = $(shell go env GOPATH)/bin
endif

VERSION    = $(shell cat $(CURDIR)/VERSION)
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

.PHONY: build
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	GO111MODULE=on go build -ldflags "$(LDFLAGS)" -o $(BINDIR)/$(BINNAME) ./cmd/karavel

.PHONY: install
install: build
	@install $(BINDIR)/$(BINNAME) $(INSTALL_PATH)/$(BINNAME)

.PHONY: clean
clean:
	rm -rf $(BINDIR)

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...
