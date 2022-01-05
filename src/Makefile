# Go parameters
BINARY_NAME=main.go 

all: run

build:
	go build $(BINARY_NAME)

clean:
	rm -rf ./bin/*

run:
	go run $(BINARY_NAME)

# Cross compilation
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(BINARY_NAME)

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(BINARY_NAME)