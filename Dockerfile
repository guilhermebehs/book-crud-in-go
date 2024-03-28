FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk update && apk add make

COPY go.mod Makefile ./

COPY . .

RUN go mod download

RUN make test

RUN make build


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/app .

EXPOSE 8080

CMD ["./app"]