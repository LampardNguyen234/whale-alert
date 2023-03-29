FROM golang:1.18-alpine
RUN apk add build-base

WORKDIR /app

COPY go.mod .
RUN go mod tidy

COPY . .

ENV CONFIG "config.json"

RUN go build -o wa

EXPOSE 12321

CMD ./wa --config $CONFIG