.PHONY: build
build:
	go build -v ./cmd/url-short/main.go

.DEFAULT_GOAL := build
