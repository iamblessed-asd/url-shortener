FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY templates ./templates
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o url-shortener ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/url-shortener .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./url-shortener"]
