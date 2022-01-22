# Author:  Niels A.D.
# Project: gowarcraft3 (https://github.com/nielsAD/gowarcraft3)
# License: Mozilla Public License, v2.0

THIRD_PARTY=third_party/StormLib/build/libstorm.a third_party/bncsutil/build/libbncsutil.a

GO_FLAGS=
GOTEST_FLAGS=-cover -cpu=1,2,4 -timeout=2m

GO=go
GOFMT=gofmt
GOLINT=$(shell $(GO) env GOPATH)/bin/golint

DIR_BIN=bin
DIR_PRE=github.com/nielsAD/gowarcraft3

PKG:=$(shell $(GO) list ./...)
DIR:=$(subst $(DIR_PRE),.,$(PKG))
CMD:=$(subst $(DIR_PRE)/cmd/,,$(shell $(GO) list ./cmd/...))
ARCH:=$(shell $(GO) env GOARCH)

ifeq ($(TEST_RACE),1)
	GOTEST_FLAGS+= -race
endif

.PHONY: all release check test fmt lint vet list clean $(CMD)

all: test release
release: $(CMD)

$(DIR_BIN):
	mkdir -p $@

$(PKG): $(THIRD_PARTY)
	$(GO) build $@

$(CMD): $(THIRD_PARTY) $(DIR_BIN)
	cd $(DIR_BIN); $(GO) build $(GO_FLAGS) $(DIR_PRE)/cmd/$@

third_party/%:
	$(MAKE) -C third_party $(subst third_party/,,$@)

check: $(THIRD_PARTY)
	$(GO) build $(PKG)

test: check fmt lint vet
	$(GO) test $(GOTEST_FLAGS) $(PKG)

fmt:
	$(GOFMT) -l $(filter-out .,$(DIR)) $(wildcard *.go)

lint:
	$(GOLINT) -set_exit_status $(PKG)

vet:
	$(GO) vet $(PKG)

list:
	@echo $(PKG) | tr ' ' '\n'

clean:
	-rm -r $(DIR_BIN)
	go clean $(PKG)
	$(MAKE) -C third_party clean
