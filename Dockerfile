FROM golang:1.18.3-alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o piadocas .

FROM scratch

COPY --from=builder /app/piadocas .

COPY --from=builder /app/.env .  

ENTRYPOINT [ "./piadocas" ]