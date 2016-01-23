FROM alpine:3.3
MAINTAINER James C. Scott <jcscott.iii@gmail.com>

RUN apk update
RUN apk add git go=1.5.3-r0 bash


RUN mkdir -p /ws/src/github.com/gophergala2016/ring_leader
ENV GOPATH /ws/src
ENV GOROOT /usr/lib/go
ENV GOBIN $GOPATH/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin
WORKDIR /ws/src/github.com/gophergala2016/ring_leader
#VOLUME /ws/src/github.com/gophergala2016/ring_leader

#RUN go get ./...
#CMD go run server.go