BINARY_NAME=go-diff
VERSION=$(shell git branch --show-current | sed 's:.*/::')

build:
	GOARCH=amd64 GOOS=linux go build -o ../build/${BINARY_NAME}-${VERSION} main.go
	cp ../build/${BINARY_NAME}-${VERSION} ../build/${BINARY_NAME}
clean:
	go clean
	rm -rf ../build