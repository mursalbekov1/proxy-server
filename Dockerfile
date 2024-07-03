FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o proxy-service ./cmd/proxy

FROM alpine as runner

COPY --from=builder /app/proxy-service .
COPY /config/config.yaml .

CMD ["./proxy-service"]
