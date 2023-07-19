SOURCES := $(wildcard *.go cmd/*/*.go)

VERSION=$(shell git describe --tags --long --dirty 2>/dev/null)

ifeq ($(VERSION),)
	VERSION=UNKNOWN
endif

app: $(SOURCES)
	go build -ldflags "-X main.version=${VERSION}" -o $@ ./cmd/app

docker: $(SOURCES) build/Dockerfile.dev
	docker build -t ph-locations:latest . -f build/Dockerfile.dev --build-arg VERSION=$(VERSION)
