FROM golang:1.23-alpine

WORKDIR /go/src/github.com/abyssparanoia/application-boilerplate/
COPY . .

ENV GO111MODULE=off

RUN apk upgrade \
    && apk add git alpine-sdk \
    && go get -u github.com/tockins/realize

ENV GO111MODULE=on

RUN go mod download 