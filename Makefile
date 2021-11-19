BINARY := protoc-gen-twirp-swift
DIR=$(PWD)
TIMESTAMP := $(shell date -u "+%Y-%m-%dT%H:%M:%SZ")
COMMIT := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
#
LDFLAGS := -ldflags "-X main.Timestamp=${TIMESTAMP} -X main.Commit=${COMMIT} -X main.Branch=${BRANCH}"

all: clean test install

test:
	go test -v ./...

lint:
	golint -set_exit_status ./...

install:
	go install ${LDFLAGS} github.com/callesvedin/protoc-gen-twirp-swift

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

run:
	protoc --swift_out=. --twirp-swift_out=. ./service.proto

build_native:
	go build -o ${GOPATH}/bin/${BINARY} ${LDFLAGS} github.com/callesvedin/protoc-gen-twirp-swift

build_linux:
	GOOS=linux GOARCH=amd64 go build -o ${BINARY} ${LDFLAGS} github.com/callesvedin/protoc-gen-twirp-swift

clean:
	-rm -f ${GOPATH}/bin/${BINARY}