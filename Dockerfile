FROM golang:1.25-alpine AS builder
ARG VER
ARG BUILDDATE

LABEL maintainer="TianRosandhy <tianrosandhy@gmail.com> (https://tianrosandhy.com)"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add git && apk add bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# # Install the package
# RUN go install -v ./...

# Build the Go app
RUN go build -ldflags="-w -s -X 'main.version=$VER' -X 'main.builddate=$BUILDDATE'" -o /build

# Expose port 9999 to the outside world
EXPOSE 9999

ENV TZ=Asia/Makassar

# Run the executable
CMD [ "/build" ]