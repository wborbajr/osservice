# 1st - Dependencies
#
# FROM golang:alpine as builder
ARG GO_VERSION=1.15.7-alpine3.13
FROM golang:${GO_VERSION} as builder

LABEL maintainer="Waldir Borba Junior <wborbajr@gmail.com>" \
    version="V0.0.1.0" \
    description="OS Service Application | wborbajr/osservice:latest"

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates gcc g++ libc-dev && update-ca-certificates
# Create appuser
RUN adduser -D -g '' appuser
#
WORKDIR /go/app
#
ENV APP_BUILD_NAME="app"
#
COPY go.mod go.sum ./
#
RUN go mod download
#
COPY . .
#
# CGO_ENABLED=0 GOARCH=arm GOARCH=amd64

RUN CGO_ENABLED=1 GOOS=linux go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -v -o ${APP_BUILD_NAME} main.go
#
RUN chmod +x ${APP_BUILD_NAME}

# 2nd - Distro Setup
#
# FROM alpine as linuxsetup
# ENV LANG=en_US.UTF-8 \
#     LANGUAGE=en_US.UTF-8
# ENV TZ America/Sao_Paulo
# RUN apk add --update --no-cache \
#     tzdata \
#     htop \
#     apk-cron \
#     bash \
#     ca-certificates \
#     && cp /usr/share/zoneinfo/${TZ} /etc/localtime \
#     && echo "${TZ}" > /etc/timezone

# RUN update-ca-certificates

# # Create the user and group files that will be used in the running
# # container to run the process as an unprivileged user.
# RUN mkdir /user && \
#     echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
#     echo 'nobody:x:65534:' > /user/group

# 3rd - Final
#
FROM alpine as final

WORKDIR /osservice/

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
#
COPY --from=builder /go/app/app .
COPY --from=builder /go/app/.env .
#
EXPOSE 3001
# Perform any further action as an unprivileged user.
USER appuser
# Run the service on container startup.
ENTRYPOINT ["./app"]
CMD ""
