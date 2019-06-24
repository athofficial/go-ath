# Build Gath in a stock Go builder container
FROM golang:1.11-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /go-ath
RUN cd /go-ath && make gath

# Pull Gath into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go-ubiq/build/bin/gath /usr/local/bin/

EXPOSE 8588 8589 30388 30388/udp
ENTRYPOINT ["gath"]
