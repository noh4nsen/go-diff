BINARY_NAME=go-diff

build:
	VERSION=$(shell git branch --show-current | sed 's:.*/::')
	GOARCH=amd64 GOOS=darwin go build -o ../build/${BINARY_NAME}-${VERSION}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ../build/${BINARY_NAME}-${VERSION} main.go
	GOARCH=amd64 GOOS=windows go build -o ../build/${BINARY_NAME}-${VERSION}-windows.exe main.go

clean:
	go clean
	rm -rf ../build