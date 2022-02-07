# syntax=docker/dockerfile:1

FROM golang:1.17

# Default start directory in image for subsequent commands
WORKDIR /app

# Copy dependency/module files and download modules to image
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy necessary source files
COPY src ./
COPY utils ./
COPY *.go ./

# Compile this application to a binary executable
RUN go build -o /computer-specs-viewer

# Execute this command when the image built from this Dockerfile is used to start a container
CMD [ "/computer-specs-viewer" ]