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

import (
	"math/big"

	"github.com/atheioschain/go-atheios/common"
)

var (
	TestNetGenesisHash = common.HexToHash("0x21a7b9d55e6325b499101fc8e68eb3e2633e758f23d9eb59e08627225568bf72") // Testnet genesis hash to enforce below configs on
	MainNetGenesisHash = common.HexToHash("0x1c76cc56b1eaf7392b6fa7f57cc9e9aa7ccbac39225c0ed911e61a80116f1581") // Mainnet genesis hash to enforce below configs on

	TestNetHomesteadBlock = big.NewInt(0) // Testnet homestead block
	MainNetHomesteadBlock = big.NewInt(0) // Mainnet homestead block

	TestNetHomesteadGasRepriceBlock = big.NewInt(0) // Testnet gas reprice block
	MainNetHomesteadGasRepriceBlock = big.NewInt(0) // Mainnet gas reprice block

	TestNetHomesteadGasRepriceHash = common.HexToHash("0x21a7b9d55e6325b499101fc8e68eb3e2633e758f23d9eb59e08627225568bf72") // Testnet gas reprice block hash (used by fast sync)
	MainNetHomesteadGasRepriceHash = common.HexToHash("0x1c76cc56b1eaf7392b6fa7f57cc9e9aa7ccbac39225c0ed911e61a80116f1581") // Mainnet gas reprice block hash (used by fast sync)

	TestNetSpuriousDragon = big.NewInt(10)
	MainNetSpuriousDragon = big.NewInt(10)

	TestNetChainID = big.NewInt(1619) // Test net default chain ID
	MainNetChainID = big.NewInt(1618) // main net default chain ID
)
