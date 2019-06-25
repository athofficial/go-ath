.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/athofficial/go-ath/cmd/swarm
	go build -o bin/gath     github.com/athofficial/go-ath/cmd/gath
	go build -o bin/bootnode github.com/athofficial/go-ath/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/athofficial/go-ath/swarm/...
