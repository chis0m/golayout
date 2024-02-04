FROM golang:1.21.6-alpine3.18 as builder

WORKDIR /app
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates curl
RUN adduser -D -g '' golang

WORKDIR /app
COPY --from=builder /app/main .
COPY .env.deploy .env
COPY storage/goose/migrations db/goose/migrations
COPY docs/swagger.json docs/swagger.json
COPY wait-for.sh .
RUN chmod +x /app/wait-for.sh

USER golang

EXPOSE 9000

CMD ["/app/main"]
