FROM golang:latest

MAINTAINER Yaroslav <yarl9812@tutanota.com>

RUN mkdir -p /go/src/landing
RUN go get github.com/pilu/fresh
RUN go get github.com/lib/pq
RUN go get golang.org/x/crypto/bcrypt

COPY . /go/src/landing

CMD cd src/landing && go get && fresh

EXPOSE 8000