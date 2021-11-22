FROM golang:latest

COPY . /go/src/testMongo

WORKDIR /go/src/testMongo

RUN go build -o api ./cmd/main.go

EXPOSE 8080