.PHONY: all build package clean

GOBIN:=$(shell pwd)/bin
GO:=go1.18.10
GOFLAGS:=-tags static -ldflags "-s -w"

all: build

build:
	GOOS=linux GOARCH=arm GOARM=7 $(GO) build $(GOFLAGS) -o $(GOBIN)/myapp .

install: build
	scp $(GOBIN)/myapp root@192.168.15.244:/mnt/us/extensions/hello/myapp

clean:
	rm -rf $(GOBIN)