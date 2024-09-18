.PHONY: all lint build

export CGO_ENABLED=1

all: lint build

lint:
	go list -f '{{.Dir}}/...' -m | xargs golangci-lint run --timeout 30m -v
	
build:
	go list -f '{{.Dir}}/...' -m | xargs go build -v
	