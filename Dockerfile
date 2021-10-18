FROM golang:1.15-alpine AS builder

MAINTAINER Karthik Pothineni

LABEL service=ip-checker

# Install golang linting
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.32.0

WORKDIR /etc/ip-checker

ENV CGO_ENABLED=0

COPY . .

RUN go mod download

# Run linter
RUN golangci-lint run -v -c golangci.yml

# Run tests
RUN go test -v -cover ./...

# Build application binary
RUN GOOS=linux GOARCH=amd64 go build -a -o ip-checker-svc -v -ldflags '-w' main.go

# Build a small image
FROM scratch

COPY --from=builder /etc/ip-checker/ip-checker-svc /

COPY --from=builder /etc/ip-checker/config/config.toml /config/

COPY --from=builder /etc/ip-checker/GeoLite2-Country /GeoLite2-Country/

ENTRYPOINT ["/ip-checker-svc"]