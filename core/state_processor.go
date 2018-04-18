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

package core

import (
	"math/big"

	"github.com/atheioschain/go-atheios/core/state"
	"github.com/atheioschain/go-atheios/core/types"
	"github.com/atheioschain/go-atheios/core/vm"
	"github.com/atheioschain/go-atheios/crypto"
	"github.com/atheioschain/go-atheios/logger"
	"github.com/atheioschain/go-atheios/logger/glog"
	"github.com/atheioschain/go-atheios/params"
)

var (
	big2  = big.NewInt(2)
	big32 = big.NewInt(32)
)

// StateProcessor is a basic Processor, which takes care of transitioning
// state from one point to another.
//
// StateProcessor implements Processor.
type StateProcessor struct {
	config *params.ChainConfig
	bc     *BlockChain
}

// NewStateProcessor initialises a new StateProcessor.
func NewStateProcessor(config *params.ChainConfig, bc *BlockChain) *StateProcessor {
	return &StateProcessor{
		config: config,
		bc:     bc,
	}
}

// Process processes the state changes according to the Ethereum rules by running
// the transaction messages using the statedb and applying any rewards to both
// the processor (coinbase) and any included uncles.
//
// Process returns the receipts and logs accumulated during the process and
// returns the amount of gas that was used in the process. If any of the
// transactions failed to execute due to insufficient gas it will return an error.
func (p *StateProcessor) Process(block *types.Block, statedb *state.StateDB, cfg vm.Config) (types.Receipts, []*types.Log, *big.Int, error) {
	var (
		receipts     types.Receipts
		totalUsedGas = big.NewInt(0)
		err          error
		header       = block.Header()
		allLogs      []*types.Log
		gp           = new(GasPool).AddGas(block.GasLimit())
	)
	// Iterate over and process the individual transactions
	for i, tx := range block.Transactions() {
		//fmt.Println("tx:", i)
		statedb.StartRecord(tx.Hash(), block.Hash(), i)
		receipt, _, err := ApplyTransaction(p.config, p.bc, gp, statedb, header, tx, totalUsedGas, cfg)
		if err != nil {
			return nil, nil, nil, err
		}
		receipts = append(receipts, receipt)
		allLogs = append(allLogs, receipt.Logs...)
	}
	AccumulateRewards(statedb, header, block.Uncles())

	return receipts, allLogs, totalUsedGas, err
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. It returns the receipt
// for the transaction, gas used and an error if the transaction failed,
// indicating the block was invalid.
func ApplyTransaction(config *params.ChainConfig, bc *BlockChain, gp *GasPool, statedb *state.StateDB, header *types.Header, tx *types.Transaction, usedGas *big.Int, cfg vm.Config) (*types.Receipt, *big.Int, error) {
	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		return nil, nil, err
	}
	// Create a new context to be used in the EVM environment
	context := NewEVMContext(msg, header, bc)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(context, statedb, config, cfg)
	// Apply the transaction to the current state (included in the env)
	_, gas, err := ApplyMessage(vmenv, msg, gp)
	if err != nil {
		return nil, nil, err
	}

	// Update the state with pending changes
	usedGas.Add(usedGas, gas)
	// Create a new receipt for the transaction, storing the intermediate root and gas used by the tx
	// based on the eip phase, we're passing wether the root touch-delete accounts.
	receipt := types.NewReceipt(statedb.IntermediateRoot(config.IsEIP158(header.Number)).Bytes(), usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = new(big.Int).Set(gas)
	// if the transaction created a contract, store the creation address in the receipt.
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(vmenv.Context.Origin, tx.Nonce())
	}

	// Set the receipt logs and create a bloom for filtering
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

	glog.V(logger.Debug).Infoln(receipt)

	return receipt, gas, err
}

// AccumulateRewards credits the coinbase of the given block with the
// mining reward. The total reward consists of the static block reward
// and rewards for included uncles. The coinbase of each uncle block is
// also rewarded.
func AccumulateRewards(statedb *state.StateDB, header *types.Header, uncles []*types.Header) {
	reward := new(big.Int).Set(BlockReward)
	rewardDev := new(big.Int).Set(DevReward)

	// Epoch 1 - Beyond Block 716727
	if header.Number.Cmp(big.NewInt(716727)) > 0 {
		reward = new(big.Int).Mul(big.NewInt(10), big.NewInt(1e+18))
		rewardDev = new(big.Int).Mul(big.NewInt(2), big.NewInt(1e+17))
	}
	// Epoch 2 - Beyond Block 1433454
	if header.Number.Cmp(big.NewInt(1433454)) > 0 {
		reward = new(big.Int).Mul(big.NewInt(9), big.NewInt(1e+18))
		rewardDev = new(big.Int).Mul(big.NewInt(3), big.NewInt(1e+17))
	}
	// Epoch 3 - Beyond Block 2866908
	if header.Number.Cmp(big.NewInt(2866908)) > 0 {
		reward = new(big.Int).Mul(big.NewInt(8), big.NewInt(1e+18))
		rewardDev = new(big.Int).Mul(big.NewInt(2), big.NewInt(1e+17))
	}
	// Epoch 4 - Beyond Block 4300362
	if header.Number.Cmp(big.NewInt(4300362)) > 0 {
		reward = new(big.Int).Mul(big.NewInt(7), big.NewInt(1e+18))
		rewardDev = new(big.Int).Mul(big.NewInt(18), big.NewInt(1e+16))
	}
	// Epoch 5 - Beyond Block 5733816
	if header.Number.Cmp(big.NewInt(5733816)) > 0 {
		reward = new(big.Int).Mul(big.NewInt(6), big.NewInt(1e+18))
		rewardDev = new(big.Int).Mul(big.NewInt(15), big.NewInt(1e+16))
	}
	// Epoch 6 - Beyond Block 7167270
	if header.Number.Cmp(big.NewInt(7167270)) > 0 {
		reward = new(big.Int).Mul(big.NewInt(5), big.NewInt(1e+18))
		rewardDev = new(big.Int).Mul(big.NewInt(1), big.NewInt(1e+17))
	}
	// Epoch 7 - Beyond Block 8600724
	if header.Number.Cmp(big.NewInt(8600724)) > 0 {
		reward = new(big.Int).Mul(big.NewInt(4), big.NewInt(1e+18))
		rewardDev = new(big.Int).Mul(big.NewInt(8), big.NewInt(1e+16))
	}
	// Epoch 8 - Beyond Block 10034178
	if header.Number.Cmp(big.NewInt(10034178)) > 0 {
		reward = new(big.Int).Mul(big.NewInt(3), big.NewInt(1e+18))
		rewardDev = new(big.Int).Mul(big.NewInt(5), big.NewInt(1e+16))
	}
	// Epoch 9 - Beyond Block 11467632
	if header.Number.Cmp(big.NewInt(11467632)) > 0 {
		reward = new(big.Int).Mul(big.NewInt(2), big.NewInt(1e+18))
		rewardDev = new(big.Int).Mul(big.NewInt(3), big.NewInt(1e+16))
	}
	// Epoch 10 - Beyond Block 14334540
	if header.Number.Cmp(big.NewInt(14334540)) > 0 {
		reward = new(big.Int).Mul(big.NewInt(1), big.NewInt(1e+18))
		rewardDev = new(big.Int).Mul(big.NewInt(1), big.NewInt(1e+16))
	}

	r := new(big.Int)
	for _, uncle := range uncles {
		r.Add(uncle.Number, big2)
		r.Sub(r, header.Number)
		r.Mul(r, BlockReward)
		r.Div(r, big2)

		if header.Number.Cmp(big.NewInt(10)) < 0 {
			statedb.AddBalance(uncle.Coinbase, r)
			r.Div(BlockReward, big32)
			if r.Cmp(big.NewInt(0)) < 0 {
				r = big.NewInt(0)
			}
		} else {
			if r.Cmp(big.NewInt(0)) < 0 {
				r = big.NewInt(0)
			}
			statedb.AddBalance(uncle.Coinbase, r)
			r.Div(BlockReward, big32)
		}

		reward.Add(reward, r)
	}
	statedb.AddBalance(header.Coinbase, reward)
	statedb.AddBalance(devFund, rewardDev)
}
