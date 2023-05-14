FROM golang:1.18.3-alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o piadocas .

RUN apk add -U --no-cache ca-certificates

FROM scratch

COPY --from=builder /app/piadocas .

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT [ "./piadocas" ]
