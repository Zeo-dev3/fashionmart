FROM golang:1.22.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o main ./cmd/api

# runtime
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app .

ENV config=docker

EXPOSE 3000

CMD ["./main"]