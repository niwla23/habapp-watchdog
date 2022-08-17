FROM golang:1.19-alpine as builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /habapp-watchdog

FROM debian:bullseye-slim
COPY --from=builder /habapp-watchdog /habapp-watchdog
ENTRYPOINT ["/habapp-watchdog"]