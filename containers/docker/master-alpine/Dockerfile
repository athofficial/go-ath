FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/1.8 https://github.com/athofficial/go-ath && \
  (cd go-ath && make gath) && \
  cp go-ath/build/bin/gath /gath && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go-ath && rm -rf /var/cache/apk/*

EXPOSE 8696
EXPOSE 30696

ENTRYPOINT ["/gath"]
