FROM golang:1.14

ARG APP_PORT
ENV GO111MODULE=on

WORKDIR /app

COPY ./ /app

RUN /app/bin/cmd/setup

EXPOSE $APP_PORT

CMD ["./bin/cmd/watch"]