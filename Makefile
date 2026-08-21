#!/usr/bin/make -f

PACKAGES=$(shell go list ./... | grep -v '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true

# docker variables
DOCKER_BUF := docker run --rm --volume "$(PWD)":/workspace --workdir /workspace bufbuild/buf:1.26.1

all: install

install: go.sum
	go install -mod=readonly ./cmd/myblockchain
	go install -mod=readonly ./cmd/myblockchaincli

go.sum: go.mod
	@echo "Ensuring go.mod and go.sum are up to date..."
	@go mod tidy

test:
	@go test -mod=readonly $(PACKAGES)

fmt:
	@go fmt ./...

lint:
	@golangci-lint run

build:
	@go build -o myblockaind ./cmd/myblockchain

build-cli:
	@go build -o myblockchainicli ./cmd/myblockchaincli

clean:
	@rm -f myblockaind myblockchainicli

help:
	@echo "Available make targets:"
	@echo "  make install        - Install the blockchain binary"
	@echo "  make build          - Build the blockchain daemon"
	@echo "  make build-cli      - Build the CLI"
	@echo "  make test           - Run tests"
	@echo "  make fmt            - Format code"
	@echo "  make lint           - Run linter"
	@echo "  make clean          - Clean built binaries"

.PHONY: all install go.sum test fmt lint build build-cli clean help
