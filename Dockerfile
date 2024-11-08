FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY ./internal ./internal
COPY ./db ./db
COPY ./common ./common

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:3.20
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 6969
ENV APP_ENV="production"
CMD ["./app"]
