FROM golang:1.7-alpine
LABEL maintainer="cassiobotaro@gmail.com"
RUN apk update && apk add curl git alpine-sdk
RUN mkdir -p /go/src/github.com/gofn/docker-in-docker
COPY  . /go/src/github.com/gofn/docker-in-docker
WORKDIR /go/src/github.com/gofn/docker-in-docker
RUN go get -u github.com/kardianos/govendor
RUN govendor sync
RUN go install
CMD ["docker-in-docker"]
