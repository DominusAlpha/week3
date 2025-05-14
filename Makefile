VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGET_OS=linux

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v	

build: format
	CGO_ENABLED=0 GOOS=${TARGET_OS} GOARCH=${shell dpkg --print-architecture} go build -v -o chahlikBot -ldflags "-X="github.com/DominusAlpha/chahlikBot/cmd.appVersion=${VERSION}

clean:
	rm -rf chahlikBot	