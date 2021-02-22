FROM golang:1.15.8-alpine AS builder
RUN apk add --update git
COPY . /app
WORKDIR /app
RUN go build -o api ./cmd/api/
RUN chmod 700 api

FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/api .
COPY --from=builder /app/roster.json .
CMD ["./api"]