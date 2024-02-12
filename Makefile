.PHONY: build-wsc-dev run-wsc-dev build-rigelctl build-docker

build-wsc-dev:
	mkdir -p out
	cd server && go build -tags dev -o ../out/rigel-server.exe .

run-wsc-dev: build-wsc-dev
	./out/rigel-server.exe

build-rigelctl:
	mkdir -p out
	cd cmd/rigelctl && go build -o ../../out/rigelctl .


# Default build tag is empty for non-dev builds
BUILD_TAGS ?=

# Build container image for rigel web services server with optional build tags
# for dev build
#     make build-docker BUILD_TAGS=dev
# for non-dev build
#     make build-docker
build-docker:
	docker build --build-arg BUILD_TAGS=$(BUILD_TAGS) -t rigelwsc:latest .