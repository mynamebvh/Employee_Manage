FROM golang:1.14 AS builder
WORKDIR /mynamebvh/employee_manager
ADD . .
RUN go build -o employee_manager