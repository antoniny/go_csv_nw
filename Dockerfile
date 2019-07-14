######################################
FROM golang:alpine as builder
#MAINTAINER Antoniny <antoniny@gmail.com>
#######################################

RUN mkdir /build 
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get -u 'github.com/antoniny/go_lang_text'
RUN go get -u "github.com/lib/pq"
COPY . /build/
WORKDIR /build 
RUN go get -d -v
#ADD . /build/
RUN go build main.go

ENTRYPOINT ["./main","file_input.txt"]


