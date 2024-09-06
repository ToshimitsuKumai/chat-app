FROM golang:1.23.1-alpine3.20

WORKDIR /go/src/app/

COPY ./backend .

RUN apk upgrade --update && \
    apk --no-cache add git gcc musl-dev

RUN adduser -s /bin/bash -D appuser
RUN chown -R appuser:appuser /go/src/app

USER appuser

RUN go install github.com/air-verse/air@latest
