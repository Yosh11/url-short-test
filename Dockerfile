FROM golang:latest

COPY . /go/src/url-short-test

WORKDIR /go/src/url-short-test

RUN go build -o api ./cmd/main.go

EXPOSE 8080