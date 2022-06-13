FROM golang:1.17.6-alpine as builder

ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN go mod vendor
RUN go test ./...