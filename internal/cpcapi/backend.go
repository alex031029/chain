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

// Package cpcapi implements the general Cpchain API functions.
package cpcapi

import (
	"context"
	"math/big"

	"bitbucket.org/cpchain/chain/accounts"
	"bitbucket.org/cpchain/chain/api/rpc"
	"bitbucket.org/cpchain/chain/configs"
	"bitbucket.org/cpchain/chain/core"
	"bitbucket.org/cpchain/chain/core/state"
	"bitbucket.org/cpchain/chain/core/vm"
	"bitbucket.org/cpchain/chain/database"
	"bitbucket.org/cpchain/chain/protocols/cpc/syncer"
	"bitbucket.org/cpchain/chain/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Backend interface provides the common API services (that are provided by
// both full and light clients) with access to necessary functions.
type Backend interface {
	// General Cpchain API
	Downloader() syncer.Syncer
	ProtocolVersion() int
	SuggestPrice(ctx context.Context) (*big.Int, error)
	ChainDb() database.Database
	EventMux() *event.TypeMux
	AccountManager() *accounts.Manager

	// BlockChain API
	SetHead(number uint64)
	HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error)
	BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error)
	StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber, isPrivate bool) (*state.StateDB, *types.Header, error)
	GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error)
	GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error)
	GetPrivateReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	GetEVM(ctx context.Context, msg core.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error)
	SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription
	SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription
	SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription
	Proposers(blockNr rpc.BlockNumber) ([]common.Address, error)
	Validators(blockNr rpc.BlockNumber) ([]common.Address, error)
	ProposerOf(blockNr rpc.BlockNumber) (common.Address, error)

	// TxPool API
	SendTx(ctx context.Context, signedTx *types.Transaction) error
	GetPoolTransactions() (types.Transactions, error)
	GetPoolTransaction(txHash common.Hash) *types.Transaction
	GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error)
	Stats() (pending int, queued int)
	TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions)
	SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription

	ChainConfig() *configs.ChainConfig
	CurrentBlock() *types.Block

	RemoteDB() database.RemoteDatabase // RemoteDB returns remote database instance.

	//RNode API
	RNode() ([]common.Address, uint64)
	CurrentProposerIndex() uint64
	CurrentTerm() uint64
	CurrentView() uint64
	CommitteMember() []common.Address
	CalcRptInfo(address common.Address, addresses []common.Address, blockNum uint64) int64
	ViewLen() uint64
	TermLen() uint64

	BlockReward(blockNr rpc.BlockNumber) *big.Int

	// Private API
	SupportPrivateTx(ctx context.Context) (bool, error)
}

func GetAPIs(apiBackend Backend) []rpc.API {
	nonceLock := new(AddrLocker)
	return []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicCpchainAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicBlockChainAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicTransactionPoolAPI(apiBackend, nonceLock),
			Public:    true,
		}, {
			Namespace: "txpool",
			Version:   "1.0",
			Service:   NewPublicTxPoolAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPublicDebugAPI(apiBackend),
			Public:    true,
		}, {
			Namespace: "debug",
			Version:   "1.0",
			Service:   NewPrivateDebugAPI(apiBackend),
		}, {
			Namespace: "eth",
			Version:   "1.0",
			Service:   NewPublicAccountAPI(apiBackend.AccountManager()),
			Public:    true,
		}, {
			Namespace: "personal",
			Version:   "1.0",
			Service:   NewPrivateAccountAPI(apiBackend, nonceLock),
			Public:    false,
		},
	}
}
