FROM golang:1.18.3-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

CMD CGO_ENABLED=0 go test -v  ./...