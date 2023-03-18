# Build stage
FROM golang:1.18-alpine3.17 AS builder

WORKDIR /github.com/b0shka/walkom-backend
COPY . .
RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

# Run stage
FROM alpine:3.17

WORKDIR /root/

COPY --from=builder /github.com/b0shka/walkom-backend/.bin/app .
COPY --from=builder /github.com/b0shka/walkom-backend/configs configs/
COPY --from=builder /github.com/b0shka/walkom-backend/.env .

EXPOSE 8000
CMD ["./app"]