FROM golang:1.21.0

WORKDIR /api

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /traceroute

EXPOSE 8080

CMD ["/traceroute"]
