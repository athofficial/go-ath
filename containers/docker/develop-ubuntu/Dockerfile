FROM ubuntu:xenial

ENV PATH=/usr/lib/go-1.9/bin:$PATH

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang-1.9 git make gcc libc-dev ca-certificates && \
  git clone --depth 1 https://github.com/athofficial/go-ath && \
  (cd go-ubiq && make gath) && \
  cp go-ubiq/build/bin/gath /gath && \
  apt-get remove -y golang-1.9 git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /go-ubiq

EXPOSE 8696
EXPOSE 30696

ENTRYPOINT ["/gath"]
