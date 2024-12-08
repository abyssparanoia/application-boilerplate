FROM golang:1.23.4-alpine AS builder

ARG SERVICE_NAME=default

WORKDIR /go/src/github.com/abyssparanoia/application-boilerplate/

ENV CGO_ENABLED=0

COPY . .

RUN go install -v -tags netgo -ldflags '-extldflags "-static"' ./cmd/rapid/


FROM alpine AS server

RUN apk add ca-certificates
COPY --from=builder /go/bin/rapid /bin/rapid

WORKDIR /go/src/github.com/abyssparanoia/application-boilerplate/

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT ["rapid","server","run"]