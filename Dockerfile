FROM golang:1.16-alpine AS builder
WORKDIR /app

ARG CMD

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main cmd/${CMD}/main.go

FROM alpine:3.14

WORKDIR /app

RUN addgroup -S app
RUN adduser -S app -G app
USER app

COPY --from=builder /app/main .

CMD ["/app/main"]