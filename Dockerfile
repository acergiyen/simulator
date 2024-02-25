#!/bin/sh
# Builder
FROM golang:1.21.5-alpine as builder

WORKDIR /app

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/simulator

# Final docker image

FROM alpine:3.7

ENV SPEC_FILE_PATH="../../config"

WORKDIR /
COPY --from=builder /app .

CMD ["/simulator"]