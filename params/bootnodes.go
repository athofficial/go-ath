// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main atheios network.
var MainnetBootnodes = []string{
	// atheios Go Bootnodes
	"enode://4c0b55e67005b76136ac81e89e126c05cb2d6b5ab8aaf73c12d94c28026ab484fe6d852f56263a674fa996a9b6d465fc8cf17d15d5c3a42f188459cf275617ce@35.230.80.37:30696",
	"enode://002197c1271ab789548ce78d5f9be3c08100c9fa189c533a2ed14caf01d8ce9d6d691bddf13105b3958f339852b1b2da28d76b18dc2123be7b70bb837bb78852@35.234.112.176:30696",
	"enode://92b81230dceb5a0615aafc293481167388e84c350410242e8b687971b3bd2d830dee3aee50c1e0206e1e068039f4a7cb9613b9e3f5a0e6e7fa35be9afcf90b02@35.229.212.25:30696",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// atheios test network.
var TestnetBootnodes = []string{
	"enode://c1b2a9293599b6690ae8cd1b38a953e52189f04d18419bb781d2e4d0b4b32951c8d2cf38fa0f6ed3971a223caee892024e99be4f67dc9fffcc818c2d561d6a16@35.197.51.231:30696",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
