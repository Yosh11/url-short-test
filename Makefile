build:
	go build -v ./cmd/url-short/main.go

run: build
	./main.exe
.DEFAULT_GOAL := run
