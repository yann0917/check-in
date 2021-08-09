# Go parameters
.PHONY: build test clean run build-race build-linux build-osx build-windows test-race enable-race

all: clean setup build-linux build-osx build-windows

BUILD_ENV=CGO_ENABLED=0
BUILD=`date +%FT%T%z`
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD}"

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
TARGET_EXEC=check-in

setup:
	mkdir -p build/linux
	mkdir -p build/osx
	mkdir -p build/windows

build-linux: setup
	$(BUILD_ENV) GOARCH=amd64 GOOS=linux $(GOBUILD) $(LDFLAGS) -o build/linux/$(TARGET_EXEC)

build-osx: setup
	$(UILD_ENV) GOARCH=amd64 GOOS=darwin $(GOBUILD) $(LDFLAGS) -o build/osx/$(TARGET_EXEC)

build-windows: setup
	$(BUILD_ENV) GOARCH=amd64 GOOS=windows $(GOBUILD) $(LDFLAGS) -o build/windows/$(TARGET_EXEC).exe

default: build-windows

build:
	$(GOBUILD) $(RACE) $(LDFLAGS) -o $(TARGET_EXEC) -v .

test:
	$(GOTEST) $(RACE) -v ./test

enable-race:
	$(eval RACE = -race)

run:
	$(GOBUILD) $(RACE) -o $(TARGET_EXEC) -v .
	 ./$(TARGET_EXEC)

clean:
	$(GOCLEAN)
	rm -rf build