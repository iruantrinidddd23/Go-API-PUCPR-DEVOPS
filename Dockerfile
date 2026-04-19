FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o go-api .

FROM alpine:3.20

WORKDIR /app

ENV APP_NAME="Go API PUCPR DevOps"
ENV APP_VERSION="1.0.0"
ENV PORT="8000"

COPY --from=builder /app/go-api ./go-api

EXPOSE 8000

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD wget -qO- http://127.0.0.1:8000/health || exit 1

CMD ["./go-api"]
