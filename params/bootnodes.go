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
	"enode://e41e2b8c2be004cdc1756c8513c4bc8adf1ef50448d6f058e0427d70b6b74e814dc4de25089fec2e833d0b03709ec12162d0bc542a79f4234852678d33e7b8d4@95.179.156.43:30696",
	"enode://3321bd6269248b8a5ef8081e4e3473affc697a7c448d7544ba58195c09293aa8f4899a4322913614554c7ebc5c713246a1fe2bfb2fca947d13bc3d4ef57623bf@99.81.45.102:30696",
	"enode://9154b067f8cfd75866070aafc363ea61fdf7987e110035103d14a75f8e9b815e86f796483b150b7cc5565a30a1cc9f3ea90b49b951477bb453e21b61dfb987c2@167.114.169.117:30696",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// atheios test network.
var TestnetBootnodes = []string{
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
