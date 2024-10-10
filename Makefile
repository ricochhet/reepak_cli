LDFLAGS=-X 'main.buildDate=$(shell date)' -X 'main.gitHash=$(shell git rev-parse HEAD)' -X 'main.buildOn=$(shell go version)' -w -s

GO_BUILD=go build -trimpath -ldflags "$(LDFLAGS)"

.PHONY: all fmt mod lint test deadcode syso reepak-linux reepak-linux-arm reepak-darwin reepak-darwin-arm reepak-windows clean

all: reepak-linux reepak-linux-arm reepak-darwin reepak-darwin-arm reepak-windows 

fmt:
	gofumpt -l -w .

mod:
	go get -u
	go mod tidy

lint:
	golangci-lint run

test:
	go test ./...

deadcode:
	deadcode ./...

syso:
	windres gomclauncher.rc -O coff -o gomclauncher.syso

reepak-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD) -o reepak-linux

reepak-linux-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GO_BUILD) -o reepak-linux-arm

reepak-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o reepak-darwin

reepak-darwin-arm:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GO_BUILD) -o reepak-darwin-arm

reepak-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO_BUILD) -o reepak-windows.exe

clean:
	rm -f reepak-linux reepak-linux-arm reepak-darwin reepak-darwin-arm reepak-windows.exe