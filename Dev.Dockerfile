FROM golang:1.18.3-alpine

WORKDIR /app

COPY . .

RUN apk add build-base

CMD [ "tail", "-f", "/dev/null" ]