// Copyright 2014 The go-ethereum Authors
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

package core

import (
	"math/big"

	"github.com/kek-mex/go-ath/common"
)

var BlockReward *big.Int = new(big.Int).Mul(big.NewInt(12), big.NewInt(1e+18))
var DevReward *big.Int = new(big.Int).Mul(big.NewInt(1), big.NewInt(1e+17))

var devFund = common.HexToAddress("0x3e5c79bc6742ff23a884b8db576bd401b3e7ff59")
