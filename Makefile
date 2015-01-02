VERSION=$(shell git rev-parse --short HEAD)

# UTC time in ISO 8601
NOW=$(shell date -u +%Y-%m-%dT%H:%M)

all:
	go build -v -o coduno -ldflags "-X main.Version '${VERSION}' -X main.BuildTime '${NOW}' -X main.Build '${}'"
