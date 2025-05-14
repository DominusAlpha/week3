APP_NAME=chahlikBot
TARGET_DIR=build
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGET_OS=linux
IMAGE_TAG=$(APP_NAME):test
PLATFORMS=linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64
TARGET_DIR=target

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v	

define make
$(1): format
	GOOS=$(2) GOARCH=$(3) CGO_ENABLED=0 go build -v -o $(TARGET_DIR)/$(1)/${APP_NAME} -ldflags "-X="github.com/DominusAlpha/chahlikBot/cmd.appVersion=${VERSION}
endef

$(eval $(call make,linux,linux,amd64))
$(eval $(call make,linux-arm,linux,arm64))
$(eval $(call make,macos,darwin,amd64))
$(eval $(call make,macos-arm,darwin,arm64))
$(eval $(call make,windows,windows,amd64))

image:
	docker buildx build --platform=$(PLATFORM) --build-arg TARGET_OS=$(TARGET_OS) --output type=docker -t $(IMAGE_TAG) -f Dockerfile.test

docker-test:
	docker run --rm $(IMAGE_TAG)

clean:
	rm -rf ${TARGET_DIR} ${APP_NAME}
	docker rmi -f $(IMAGE_TAG) || true