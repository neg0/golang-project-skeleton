FROM golang:1.12.9-buster

RUN apt-get update -y && apt-get install -y \
    curl

RUN curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

HEALTHCHECK \
  --interval=1m \
  --timeout=5s \
  CMD go version || exit 1