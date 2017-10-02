FROM golang:1.8-alpine as build-stage
LABEL maintainer="cassiobotaro@gmail.com"
WORKDIR /go/src/example
COPY dockerize.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
LABEL maintainer="cassiobotaro@gmail.com"
COPY --from=build-stage /go/src/example/main /
CMD ["/main"]
