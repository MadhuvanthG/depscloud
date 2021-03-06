GIT_SHA ?= $(shell git rev-parse HEAD)
VERSION ?= local
TIMESTAMP ?= $(shell date +%Y-%m-%dT%T)
LD_FLAGS := -X main.version=${VERSION} -X main.commit=${GIT_SHA} -X main.date=${TIMESTAMP}

build-deps:
	GO111MODULE=off go get -u golang.org/x/lint/golint
	GO111MODULE=off go get -u oss.indeed.com/go/go-groups

deps:
	go mod download
	go mod verify

fmt:
	go-groups -w .
	gofmt -s -w .

test:
	go vet ./...
	# golint -set_exit_status ./...
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

build:
	go build -ldflags="${LD_FLAGS}" .

install:
	go install -ldflags="${LD_FLAGS}" .

docker:
	docker build . \
		--build-arg GIT_SHA=${GIT_SHA} \
		--build-arg VERSION=${VERSION} \
		--build-arg TIMESTAMP=${TIMESTAMP} \
		-t depscloud/cli:latest \
		-f Dockerfile.dev
