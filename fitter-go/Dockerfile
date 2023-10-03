FROM golang:1.21-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

# System dependencies
RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    && update-ca-certificates

### Development with hot reload and debugger
FROM base AS dev

WORKDIR /app

# Hot reloading mod
# RUN go install github.com/cosmtrek/air@latest && go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/cosmtrek/air@v1.40.4 && go install github.com/go-delve/delve/cmd/dlv@v1.21.0

EXPOSE 8080 8080
EXPOSE 2345 2345

CMD ["air", "--config", ".air.toml"]

ENTRYPOINT ["air"]
### END OF DEV

### Executable builder
FROM base AS builder
WORKDIR /app

# Application dependencies
COPY . /app
RUN go mod download \
    && go mod verify

RUN go build -buildvcs=false -o my-great-program -a .

### Production
FROM alpine:latest

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    && update-ca-certificates

# Copy executable
COPY --from=builder /app/my-great-program /usr/local/bin/my-great-program
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/my-great-program -buildvcs=false"]
 
