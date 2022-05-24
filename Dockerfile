FROM golang:1.16-alpine3.14
RUN apk update && apk add alpine-sdk && apk add git curl
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app