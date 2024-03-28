FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod go.sum Makefile ./

RUN go mod download

RUN make test

RUN make build


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/app .

EXPOSE 8080

CMD ["./app"]