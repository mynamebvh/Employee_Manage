FROM golang:1.18-alpine AS build
LABEL maintainer="Hoang Bui <mynamebvh@gmail.com>"
WORKDIR /backend
COPY . .
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]