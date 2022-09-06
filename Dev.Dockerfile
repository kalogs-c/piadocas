FROM golang:1.18.3-alpine

WORKDIR /app

COPY . .

CMD [ "tail", "-f", "/dev/null" ]