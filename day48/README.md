# Minimal docker image

## Build go

`CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .`

## Build an image

`docker build . -t example-scratch`

## Run

`docker run --rm -it example-scratch`
