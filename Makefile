BIN_DIR := ./bin/
BIN_NAME := scaffolder
MODULE_NAME :=  "github.com/twistingmercury/scaffolder"

default: help

.PHONY: help
help:
	@echo "\nScaffolder makefile usage: make [target]"
	@echo "  Targets:"
	@echo "  » clean           Remove build artifacts and clean up the project"
	@echo "  » bin             Build the binary and output to _bin/ directory"
	@echo "  » test            Run all unit tests and generate coverage report"

.PHONY: clean
clean:
	@rm -rf $(BIN_DIR) > /dev/null 2>&1

.PHONY: bin
bin: clean
	go build \
	-ldflags "-s -w" \
	-o "$(BIN_DIR)$(BIN_NAME)" ./main.go

.PHONY: test
test:
	go test -v ./cmd -coverprofile=coverage.out
	go tool cover -html=coverage.out