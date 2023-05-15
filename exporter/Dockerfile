FROM golang:1.20-alpine as builder

WORKDIR /build
COPY . .

WORKDIR /build

RUN go install
RUN go build -o exporter .

FROM alpine:3.17

RUN apk --no-cache add curl jq bash

WORKDIR /root

COPY --from=builder /build/exporter /root/exporter

EXPOSE 10000

CMD [ "/root/exporter" ]
