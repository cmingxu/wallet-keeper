VERSION=$(shell cat ./VERSION)
PKG=github.com/cmingxu/wallet-keeper
GOBUILD=GO111MODULE=on CGO_ENABLED=1 go build -a -ldflags "-X main.Version=${VERSION}"
CROSS_GOBUILD=CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -a -ldflags "-X main.Version=${VERSION}"
CMDS = $(shell go list ${PKG}/cmd )
PKG_ALL = $(shell go list ${PKG}/...)
DOCKER=$(shell which docker)
BUILD_DIR=./bin

all: build

build:
	${GOBUILD}  -o ${BUILD_DIR}/wallet-keeper ./cmd/*.go

install: binaries

test:
	go test ./... -v

.PHONY: clean
clean:
	@rm bin/*

.PHONY: coverage
coverage:
	go test -cover -coverprofile=test.coverage
	go tool cover -html=test.coverage -o coverage.html
	rm -f test.coverage

.PHONY: fmt
fmt:
	go fmt ${PKG_ALL}

