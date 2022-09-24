GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOTOOL=$(GOCMD) tool
GOTHRESHOLD=github.com/jokeyrhyme/go-coverage-threshold/cmd/go-coverage-threshold
GOTHRESHOLDCMD=go-coverage-threshold
BINARY_NAME=bin/ccc

.PHONY: default deps build test clean

default: build

deps:
		$(GOCMD) mod tidy

build:
		$(GOBUILD) -o $(BINARY_NAME) ./cmd/ccc/main.go

test:
		$(GOGET) -u $(GOTHRESHOLD)
		$(GOTEST) -v ./... -coverprofile cp.out
		$(GOTOOL) cover -html=cp.out
		$(GOTHRESHOLDCMD) -t 0

clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
