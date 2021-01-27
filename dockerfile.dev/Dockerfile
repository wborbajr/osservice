# 1st - Stage
#
ARG GO_VERSION=1.15.7-alpine3.13
FROM golang:${GO_VERSION} as dev

LABEL maintainer="Waldir Borba Junior <wborbajr@gmail.com>" \
    version="V0.0.1.0" \
    description="OS Service Application | wborbajr/osservice:latest"

ENV APP_NAME="osservice" \
    APP_PATH="/osservice/app" \
    APP_PORT=3031

ENV APP_BUILD_NAME="${APP_NAME}"

WORKDIR ${APP_PATH}
COPY . ${APP_PATH}

ENV CGO_ENABLED=0 \
    GOOS=linux

EXPOSE ${APP_PORT}
ENTRYPOINT ["sh"]

# 2nd - Stage
#
FROM dev as build

RUN go mod download && go mod tidy
RUN go build -ldflags="-s -w" -v -o ${APP_BUILD_NAME} main.go
RUN chmod +x ${APP_BUILD_NAME}

# 3rd - Stage
#
FROM scratch AS prod

ENV APP_BUILD_PATH="/osservice/app" \
    APP_BUILD_NAME="osservice"

WORKDIR ${APP_BUILD_PATH}

COPY --from=build ${APP_BUILD_PATH}/${APP_BUILD_NAME} .
COPY --from=dev ${APP_BUILD_PATH}/.env .

EXPOSE ${APP_PORT}
ENTRYPOINT ["./osservice"]
CMD ""
