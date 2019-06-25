# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gath android ios gath-cross swarm evm all test clean
.PHONY: gath-linux gath-linux-386 gath-linux-amd64 gath-linux-mips64 gath-linux-mips64le
.PHONY: gath-linux-arm gath-linux-arm-5 gath-linux-arm-6 gath-linux-arm-7 gath-linux-arm64
.PHONY: gath-darwin gath-darwin-386 gath-darwin-amd64
.PHONY: gath-windows gath-windows-386 gath-windows-amd64

GOBIN = $(shell pwd)/build/bin
GO ?= latest

gath:
	build/env.sh go run build/ci.go install ./cmd/gath
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gath\" to launch gath."

swarm:
	build/env.sh go run build/ci.go install ./cmd/swarm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/swarm\" to launch swarm."

all:
	build/env.sh go run build/ci.go install

android:
	build/env.sh go run build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gath.aar\" to use the library."

ios:
	build/env.sh go run build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gath.framework\" to use the library."

test: all
	build/env.sh go run build/ci.go test

lint: ## Run linters.
	build/env.sh go run build/ci.go lint

clean:
	./build/clean_go_build_cache.sh
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

swarm-devtools:
	env GOBIN= go install ./cmd/swarm/mimegen

# Cross Compilation Targets (xgo)

gath-cross: gath-linux gath-darwin gath-windows gath-android gath-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gath-*

gath-linux: gath-linux-386 gath-linux-amd64 gath-linux-arm gath-linux-mips64 gath-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gath-linux-*

gath-linux-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gath
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gath-linux-* | grep 386

gath-linux-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gath
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gath-linux-* | grep amd64

gath-linux-arm: gath-linux-arm-5 gath-linux-arm-6 gath-linux-arm-7 gath-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gath-linux-* | grep arm

gath-linux-arm-5:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gath
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gath-linux-* | grep arm-5

gath-linux-arm-6:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gath
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gath-linux-* | grep arm-6

gath-linux-arm-7:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gath
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gath-linux-* | grep arm-7

gath-linux-arm64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gath
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gath-linux-* | grep arm64

gath-linux-mips:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gath
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gath-linux-* | grep mips

gath-linux-mipsle:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gath
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gath-linux-* | grep mipsle

gath-linux-mips64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gath
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gath-linux-* | grep mips64

gath-linux-mips64le:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gath
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gath-linux-* | grep mips64le

gath-darwin: gath-darwin-386 gath-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gath-darwin-*

gath-darwin-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gath
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gath-darwin-* | grep 386

gath-darwin-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gath
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gath-darwin-* | grep amd64

gath-windows: gath-windows-386 gath-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gath-windows-*

gath-windows-386:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gath
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gath-windows-* | grep 386

gath-windows-amd64:
	build/env.sh go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gath
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gath-windows-* | grep amd64
