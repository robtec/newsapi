COMMIT_ID=$(shell git rev-parse --short HEAD)
VERSION=$(shell cat VERSION)

OSARCH = "darwin/amd64 linux/amd64 windows/amd64"

DIST=dist

BIN=news

all: clean deps build

clean:
	@echo ">> cleaning..."
	@rm -rf $(DIST)

deps:
	@echo ">> installing deps..."
	@go get -u github.com/golang/dep/cmd/dep
	@go get -u github.com/mitchellh/gox
	@dep ensure

build: clean deps
	@echo ">> building..."
	@echo "Commit: $(COMMIT_ID)"
	@echo "Version: $(VERSION)"
	@gox -osarch $(OSARCH) -ldflags "-X main.Version=$(VERSION) -X main.CommitId=$(COMMIT_ID)" -output "$(DIST)/$(BIN)-{{.OS}}-{{.Arch}}" ./cmd/news

install:
	@go install -ldflags "-X main.Version=$(VERSION) -X main.CommitId=$(COMMIT_ID)" ./cmd/news

.PHONY: all clean build install
