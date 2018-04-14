# Use Alpine Linux
FROM golang:1.9-alpine as builder 

# Install basic dependencies
RUN apk update
RUN apk add --no-cache wget
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://raw.githubusercontent.com/sgerrand/alpine-pkg-glibc/master/sgerrand.rsa.pub

# Add glibc dependency as previous to 1.6 Geth needs it in order to compile and work properly
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.25-r0/glibc-2.25-r0.apk
RUN apk add --no-cache make gcc musl-dev linux-headers git glibc-2.25-r0.apk

# Fetch gath source and build it
RUN git clone https://github.com/atheioschain/go-atheios go-atheios/
RUN cd go-atheios && make gath

# Pull gath into another container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/go-atheios/build/bin/gath /usr/local/bin/

# Expose ports
EXPOSE 8588 8588

# Entry point
ENTRYPOINT ["gath"]