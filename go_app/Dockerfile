FROM golang:latest

USER 1000:1000
WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download && go mod verify
