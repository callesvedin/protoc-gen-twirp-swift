BINARY := protoc-gen-twirp_swift

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
	go install ${LDFLAGS} github.com/callesvedin/twirp-swift

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

run:
	mkdir -p example/swift_client && \
    protoc --twirp_out=. --go_out=. --twirp_swift_out=package_name=haberdasher:./example/swift_client ./example/service.proto

build_linux:
	GOOS=linux GOARCH=amd64 go build -o ${BINARY} ${LDFLAGS} github.com/callesvedin/twirp-swift



clean:
	-rm -f ${GOPATH}/bin/${BINARY}