FROM golang:1.11.4-alpine3.8 as builder

RUN apk update && apk upgrade && \
    apk --update add git gcc make && \
    go get -u github.com/golang/dep/cmd/dep

COPY ./ /code/

RUN go get -u github.com/kardianos/govendor && \
    go get -u github.com/pilu/fresh

# RUN govendor 

CMD fresh
