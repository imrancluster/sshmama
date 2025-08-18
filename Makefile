APP=sshmama

.PHONY: build run test clean install

build:
	go build -ldflags="-s -w -X github.com/imrancluster/sshmama/pkg/version.Version=$$(git describe --tags --always 2>/dev/null || echo dev)" -o bin/$(APP) ./cmd/sshmama

install:
	go install -ldflags="-s -w -X github.com/imrancluster/sshmama/pkg/version.Version=$$(git describe --tags --always 2>/dev/null || echo dev)" ./cmd/sshmama

test:
	go test ./...

clean:
	rm -rf bin
