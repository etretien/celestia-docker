FROM golang:1.20-alpine as builder

RUN apk --no-cache add make gcc musl-dev git

WORKDIR /build
RUN git clone https://github.com/celestiaorg/celestia-app.git && \
    cd celestia-app && \
    git checkout v0.13.2 && \
    make build

FROM alpine:3.17

RUN apk --no-cache add curl jq bash

WORKDIR /root

COPY --from=builder /build/celestia-app/build/celestia-appd /root/celestia-appd
