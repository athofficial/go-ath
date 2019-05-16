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
	"enode://eec3280e5d827c857fc66a68ac6d72f6a94cce6483f880abc2e1846fcb77ce64d13bf6a643ad6cdc2c71a7c1d379860b095fab9877ed7e25c0c9addbc8831b9f@213.249.173.253:30696",
	"enode://4c0b55e67005b76136ac81e89e126c05cb2d6b5ab8aaf73c12d94c28026ab484fe6d852f56263a674fa996a9b6d465fc8cf17d15d5c3a42f188459cf275617ce@35.230.80.37:30696",
	"enode://002197c1271ab789548ce78d5f9be3c08100c9fa189c533a2ed14caf01d8ce9d6d691bddf13105b3958f339852b1b2da28d76b18dc2123be7b70bb837bb78852@35.234.112.176:30696",
	"enode://f7c4467562a52a90336e72385cc91efa09866cd09e6ea13393dbf5a5d42836263bb2678d4160890d39aca9edf244b2852e2636ef92ae1a15ff9f958fa86a0fc0@35.229.212.25:30696",
	"enode://ddd52f85230ebd07259a3e22187091373f0c74b2dc7c777f1687cbafa78fb49c484c0d80b479455d63987819c5aa91e252ce18fe51f2270e75cef77acd29fd01@35.199.65.5:30696",
	"enode://95bee7fc08916698a691f53911e557121c09d45742f6bd9965bf38de3c85bd27cc3ed7d568ded858a1c3b7861e351f826e54be0dc820b8f2d8350771ac181703@35.201.8.2:30696",
	"enode://4c5021c5ff91b022470f56936aa291a93f18afe22819309d14e6887c9e1483cc285ced7a115c5ce0407db91ce462078987883a753a4b18768b03bbc732c91c2c@35.203.118.142:30696",
	"enode://505f5ee260b6d2ca815c68118e8865036aa16ba5833f93a795dd9a4a4f27fb0fe669f1495a8877eb79a9dd35d5fb7fc843a9f990a22311a3c01e6bd3d9488263@35.200.57.196:30696",
	"enode://e41e2b8c2be004cdc1756c8513c4bc8adf1ef50448d6f058e0427d70b6b74e814dc4de25089fec2e833d0b03709ec12162d0bc542a79f4234852678d33e7b8d4@95.179.156.43:30696",
	"enode://70637b20c02b3da0437b247b294d2d4b271e33a33089669f366bac97d96e9f8b1c85b507cdb1e4774f26b191f7858d29b45cb565269d32d805555f18424f1be7@207.148.28.227:30696",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// atheios test network.
var TestnetBootnodes = []string{
	"enode://c1b2a9293599b6690ae8cd1b38a953e52189f04d18419bb781d2e4d0b4b32951c8d2cf38fa0f6ed3971a223caee892024e99be4f67dc9fffcc818c2d561d6a16@35.197.51.231:30696",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
