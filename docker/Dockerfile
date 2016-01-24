FROM alpine:3.3
MAINTAINER James C. Scott <jcscott.iii@gmail.com>

RUN apk update
RUN apk add git go=1.5.3-r0 bash

ENV GOPATH /ws
ENV GOROOT /usr/lib/go
ENV GOBIN $GOPATH/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin
ENV GO15VENDOREXPERIMENT 1
