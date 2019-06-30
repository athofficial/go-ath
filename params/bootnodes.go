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
	"enode://44bc9f5fb3664e399e5afa4ec710085f974f5ae3042ae86c6053851a19cc715c6e3b246429d517c663923eb77ebc665027e88537b16dae14e2af15bf3cee4dee@8.9.5.171:30696",
	"enode://9154b067f8cfd75866070aafc363ea61fdf7987e110035103d14a75f8e9b815e86f796483b150b7cc5565a30a1cc9f3ea90b49b951477bb453e21b61dfb987c2@167.114.169.117:30696",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ubiq test network.
var TestnetBootnodes = []string{
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
