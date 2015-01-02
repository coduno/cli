VERSION=$(shell git rev-parse --short HEAD)

# UTC time in ISO 8601
NOW=$(shell date -u +%Y-%m-%dT%H:%M)

OUTPUT?='coduno'

all: test
	@go env
	@go version

	go build -work -x -v -o ${OUTPUT} \
		-ldflags "-X main.Version '${VERSION}' -X main.BuildTime '${NOW}'"

test: get
	go test

get:
	go get -d
