# Docker in Docker

How to run docker inside docker

# How to run?
```
docker build -t did .
docker run --privileged -v /var/run/docker.sock:/var/run/docker.sock did:latest

# or using docker-compose
docker-compose up
```
