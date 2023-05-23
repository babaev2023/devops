# Building binary
FROM golang:1.20-alpine AS builder

ARG CGO_ENABLED=0
ARG GOOS=linux

WORKDIR /go/src/hello

COPY src/main.go .

RUN go mod init \
    && go build -o /go/bin/hello

RUN ls -lah && ls -lah /go/bin/hello

# Building run container
FROM alpine:latest

WORKDIR /app

COPY --from=builder /go/bin/hello .

CMD ./hello
