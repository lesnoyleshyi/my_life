FROM golang:latest as builder

RUN mkdir -p /api
ADD . /api
WORKDIR /api

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build ./cmd/api/main.go -o service

FROM scratch

COPY --from=builder /api /api
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

CMD ["/api/service"]