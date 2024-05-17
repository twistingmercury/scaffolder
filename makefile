BIN_DIR := "./_bin/"
BIN_NAME := "scaffolder"
BUILD_DATE := $(shell date +"%Y-%m-%dT")
BUILD_VER := "n/a"
GIT_COMMIT := "n/a"
MODULE_NAME :=  "github.com/twistingmercury/scaffolder"

ifeq ($(shell git rev-parse --is-inside-work-tree 2>/dev/null),true)
  TAG := $(shell git describe --tags --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2>/dev/null)
  ifdef TAG
    BUILD_VER := $(shell echo $(TAG) | sed 's/v//')
  else
    BUILD_VER := ""
  endif
  GIT_COMMIT := $(shell git rev-parse --short HEAD)pwd
endif

default: help

.PHONY: help
help:
	@echo "\nScaffolder makefile usage: make [target]"
	@echo "  Targets:"
	@echo "  » clean           Remove build artifacts"
	@echo "  » build           Build the api binary"
	@echo "  » test            Run all unit tests"

.PHONY: clean
clean:
	@rm -rf $(BIN_DIR) > /dev/null 2>&1

.PHONY: build
build: clean
	go build \
	-ldflags "-X '$(MODULE_NAME)/cmd/conf.buildDate=$(BUILD_DATE)' \
	-X '$(MODULE_NAME)/cmd/conf.buildVer=$(BUILD_VER)' \
	-X '$(MODULE_NAME)/cmd/conf.buildCommit=$(GIT_COMMIT)' -s -w" \
	-o $(BIN_DIR)$(BIN_NAME) ./main.go

.PHONY: test
test:
	go test -v ./conf ./api -coverprofile=coverage.out
	go tool cover -html=coverage.out