FROM golang:latest as build

MAINTAINER czarhao

ENV GOPROXY https://goproxy.io/

ENV GO111MODULE on

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .

RUN go mod download

WORKDIR /app

ADD . .

RUN go build chat

EXPOSE 8080

CMD ["./chat"]