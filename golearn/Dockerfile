# Build stage
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o golearn .

# Run stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates sqlite-libs

WORKDIR /app

COPY --from=builder /app/golearn .

EXPOSE 8080

ENV GIN_MODE=release
ENV PORT=8080
ENV DB_PATH=/app/data/golearn.db
ENV JWT_SECRET=super-secret-key-golearn-2024

RUN mkdir -p /app/data

CMD ["./golearn"]
