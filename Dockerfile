FROM golang:1.18-alpine AS build
LABEL maintainer="Hoang Bui <mynamebvh@gmail.com>"
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o employee_manager


FROM alpine:latest
WORKDIR /employee_manager
COPY --from=build /build .
EXPOSE 8080
CMD ["./employee_manager"]