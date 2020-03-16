GO ?= go

download:
	$(GO) mod download

build:
	$(GO) build -o ./core cmd/core/main.go

run:
	$(GO) run cmd/core/main.go
