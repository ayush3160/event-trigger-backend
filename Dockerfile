FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final image
FROM debian:bookworm-slim
WORKDIR /root/

COPY --from=builder /app/main .
RUN chmod +x main

EXPOSE 8000
CMD ["./main"]