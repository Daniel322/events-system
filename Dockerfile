FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go.mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldrflags="-s -w" -o app .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]