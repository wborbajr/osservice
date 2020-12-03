# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
# 1st - Dependencies
#
FROM golang:1.15.5 as builder

WORKDIR /osservice

COPY . .

RUN go mod download
