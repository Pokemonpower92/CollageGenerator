# Makefile for a Go project with two applications: app1 and app2

# Go related variables
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_VET=$(GO_CMD) vet
GO_CLEAN=$(GO_CMD) clean
GO_FMT=$(GO_CMD) fmt

# Application directories
APP1_DIR=cmd/imagesetparser

# Output binary names
APP1_BINARY=bin/imagesetparser

# Build step for app1
imagesetparser: vet
	$(GO_BUILD) -C $(APP1_DIR) -o ../../$(APP1_BINARY)

vet: fmt
	$(GO_VET) ./...

fmt:
	$(GO_FMT) ./...

# Clean step
clean:
	$(GO_CLEAN) && \
	rm ./bin/*

# Composite targets
build: imagesetparser
all: vet build

.PHONY: build fmt vet clean all

