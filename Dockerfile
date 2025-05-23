FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod download

#COPY .env ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .
#COPY --from=builder /app/.env .

CMD ["./app"]
