files = main.go
binary_name = main

all: deps build

build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags='-s -w -extldflags "-static"' -o $(binary_name) $(files)

deps:
	go get github.com/gorilla/mux
	go get github.com/tidwall/gjson

clean:
	go clean