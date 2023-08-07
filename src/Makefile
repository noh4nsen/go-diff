BINARY_NAME=go-diff

build:
	GOARCH=amd64 GOOS=darwin go build -o ../build/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ../build/${BINARY_NAME} main.go
	GOARCH=amd64 GOOS=windows go build -o ../build/${BINARY_NAME}-windows.exe main.go

clean:
	go clean
	rm ../build/${BINARY_NAME}-darwin
	rm ../build/${BINARY_NAME}
	rm ../build/${BINARY_NAME}-windows.exe