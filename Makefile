.PHONY: all build assets clean

GOBIN:=$(shell pwd)/bin
ASSETS:=$(shell pwd)/assets
GO:=go1.18.10
GOFLAGS:=-tags static -ldflags "-s -w"

all: build install assets

build:
	GOOS=linux GOARCH=arm GOARM=7 $(GO) build $(GOFLAGS) -o $(GOBIN)/myapp .

install: build
	scp $(GOBIN)/myapp root@192.168.15.244:/mnt/us/extensions/hello/myapp

assets:
	scp -R $(ASSETS)/* root@192.168.15.244:/mnt/us/extensions/hello/assets

clean:
	rm -rf $(GOBIN)