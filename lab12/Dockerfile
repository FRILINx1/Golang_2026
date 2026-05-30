FROM golang:alpine AS builder

ARG SERVICE_DIR

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o service ${SERVICE_DIR}/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/service .

CMD ["./service"]
