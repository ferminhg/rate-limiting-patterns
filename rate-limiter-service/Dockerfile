FROM golang:1.23 AS development
RUN apt-get update && apt-get install -y make curl

ENV APP_HOME=/app

RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

ENV API_HOST="0.0.0.0"
ENV API_PORT=3010

COPY rate-limiter-service/go.mod rate-limiter-service/go.sum ./
RUN go mod download

COPY rate-limiter-service/ ./

RUN go env -w GOFLAGS=-buildvcs=false && \
  go env -w CGO_ENABLED=0

RUN go build -o /app/build/server ./cmd/server.go

ENTRYPOINT [ "make", "run-dev" ]

FROM scratch AS production
COPY --from=development /app/build/server /app/build/server
ENTRYPOINT ["/app/build/server"]
