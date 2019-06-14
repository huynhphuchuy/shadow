# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang v1.12 base image
FROM golang:1.12

# Add Maintainer Info
LABEL maintainer="Huy Huynh <huynhphuchuy@live.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/shadow

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure

# This container exposes port 6969 to the outside world
EXPOSE 6969

# Run the executable
CMD ["sh", "SERVER.sh"]