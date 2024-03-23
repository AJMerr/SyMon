.DEFAULT_GOAL := build

.PRINT: fmt vet build 

fmt:
				go fmt ./...

vet:
				go vet ./...

build: vet 
				go build

clean:
				go clean
