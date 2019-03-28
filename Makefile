VERSION=$(shell cat ./VERSION)
PKG=git.jd.com/jex/jex-fooabr
GOBUILD=CGO_ENABLED=0 go build -a -ldflags "-X main.Version=${VERSION}"
CROSS_GOBUILD=CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -ldflags "-X main.Version=${VERSION}"
CMDS = $(shell go list ${PKG}/cmd )
PKG_ALL = $(shell go list ${PKG}/...)
DOCKER=$(shell which docker)
BUILD_DIR=./bin

all: binaries
	
binaries: go-web go-job

go-web:
	${GOBUILD}  -o ${BUILD_DIR}/web_${VERSION} ./cmd/web.go ./cmd/flags.go
	${CROSS_GOBUILD}  -o ${BUILD_DIR}/web_linux_${VERSION} ./cmd/web.go ./cmd/flags.go

go-job:
	${CROSS_GOBUILD} -o ${BUILD_DIR}/job ./cmd/job.go ./cmd/flags.go
	${CROSS_GOBUILD}  -o ${BUILD_DIR}/job_linux_${VERSION} ./cmd/job.go ./cmd/flags.go

install: binaries

test:
	go test ./...

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

