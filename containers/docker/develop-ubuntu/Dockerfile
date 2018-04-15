FROM ubuntu:xenial

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang git make gcc libc-dev ca-certificates && \
  git clone --depth 1 https://github.com/atheioschain/go-atheios && \
  (cd go-atheios && make gath) && \
  cp go-atheios/build/bin/gath /gath && \
  apt-get remove -y golang git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /go-atheios

EXPOSE 8696
EXPOSE 30696

ENTRYPOINT ["/gath"]
