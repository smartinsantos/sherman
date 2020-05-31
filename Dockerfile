FROM golang:1.14.3-alpine

ARG APP_PORT
ENV GO111MODULE=on

RUN apk update && apk add bash

WORKDIR /app

COPY ./ /app

RUN /app/bin/cmd/setup

EXPOSE $APP_PORT

CMD ["./bin/cmd/watch"]