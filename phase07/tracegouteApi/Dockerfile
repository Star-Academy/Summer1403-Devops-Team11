#builder
FROM golang:1.22.5-alpine3.20 AS builder

WORKDIR /app 

RUN adduser -D appuser

RUN chown -R appuser /app 

USER appuser

COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./tracegoute

#runner
FROM scratch 

WORKDIR /app

#COPY --from=builder --chown=root:root app/.env ./.env
COPY --from=builder --chown=root:root app/tracegoute ./tracegoute

CMD ["./tracegoute"]
