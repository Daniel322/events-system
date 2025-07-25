FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

RUN apk add --no-cache bash curl zip

RUN curl -fsSL \
https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
sh 

EXPOSE 8080

CMD ["./app"]