# This Dockerfile has been think to be optimized on its size and on its reconstruction, using as possible the docker cache

# Download a light image of golang
FROM golang:1.24-alpine AS build_image
ENV GOROOT="/usr/local/go"

# If you have any question on this Dockerfile
LABEL maintainer="ABET Jean-Baptiste"

#   Because sqlite is a CGO enabled package.
ENV CGO_ENABLED=1

RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev

# Copy the application
COPY . /app

# Set the working directory (inside the docker container) to be in the folder /app. Every next RUN command will be located in this directory
WORKDIR /app/

# Build the api application to create the executable file
RUN go build -o ./main ./cmd/api

#use a very light image to host the app
FROM alpine:3.15

RUN apk --no-cache add tzdata

WORKDIR /app/

# copy from builded image
COPY --from=build_image /app/main ./api/main

WORKDIR /app/api

CMD ["./main"]