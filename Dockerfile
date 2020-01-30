FROM golang:1.13

COPY *.go ./

ENV GO111MODULE=on

RUN go build *.go