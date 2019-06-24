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
// the main Ubiq network.
var MainnetBootnodes = []string{
	// Atheios Go Bootnodes
	"enode://eec3280e5d827c857fc66a68ac6d72f6a94cce6483f880abc2e1846fcb77ce64d13bf6a643ad6cdc2c71a7c1d379860b095fab9877ed7e25c0c9addbc8831b9f@213.249.173.253:30696",
	"enode://e41e2b8c2be004cdc1756c8513c4bc8adf1ef50448d6f058e0427d70b6b74e814dc4de25089fec2e833d0b03709ec12162d0bc542a79f4234852678d33e7b8d4@95.179.156.43:30696",
	"enode://3321bd6269248b8a5ef8081e4e3473affc697a7c448d7544ba58195c09293aa8f4899a4322913614554c7ebc5c713246a1fe2bfb2fca947d13bc3d4ef57623bf@99.81.45.102:30696",
	"enode://9154b067f8cfd75866070aafc363ea61fdf7987e110035103d14a75f8e9b815e86f796483b150b7cc5565a30a1cc9f3ea90b49b951477bb453e21b61dfb987c2@167.114.169.117:30696",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ubiq test network.
var TestnetBootnodes = []string{
	"enode://81b11410a96e0ea6ecc927f7714a2c256c13e200bc73d087ad120e5a3fc3e1e098760c6fae3dcd7a3c393e49c205e636bacfc10adf6581672f6d3a66e2442248@45.77.7.41:30388",
	"enode://0595ec507bb779873703f516072b37d07f3305271da3d9585ada3b1734535635eac50cffd8c9a413b87a77ede5f49af391a08ca9348d027fda74c38f1ea5ec91@108.61.188.12:30388",
	"enode://8cb060312b4667ed6a0f61dd6cc0cd5d39e70c17429cd5e8ca480fcd7caf72f1b9c92884ce1f8e06e84a7ed1580ba302df0e95ec2ce99f727297bd2787ed8149@45.76.90.144:30388",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
