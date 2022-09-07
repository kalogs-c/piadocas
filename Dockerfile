FROM golang:1.18.3-alpine as builder

WORKDIR /app

COPY . .

ENV DB_HOST=127.0.0.1
ENV DB_USER=killerik13
ENV DB_PASSWORD=monstrinho55
ENV DB_NAME=jokes
ENV DB_PORT=3306

RUN CGO_ENABLED=0 GOOS=linux go build -a -o piadocas .

FROM scratch

COPY --from=builder /app/piadocas .

ENTRYPOINT [ "./piadocas" ]