package primitives

import (
	"math/big"

	"bitbucket.org/cpchain/chain/commons/log"
	"bitbucket.org/cpchain/chain/configs"
	"github.com/ethereum/go-ethereum/common"
)

//go:generate abigen --sol ../contracts/primitive_contracts.sol --pkg contracts --out ../contracts/primitive_contracts.go

// Definitions of primitive contracts

type GetMaintenance struct {
	Backend RptPrimitiveBackend
}

func (c *GetMaintenance) RequiredGas(input []byte) uint64 {
	return configs.GetMaintenance
}

func (c *GetMaintenance) Run(input []byte) ([]byte, error) {
	const getBalanceInputLength = 52

	if len(input) != getBalanceInputLength {
		log.Warnf("The expected input is %d bytes, but got %d", getBalanceInputLength, len(input))
		return common.LeftPadBytes(new(big.Int).Bytes(), 32), nil
	}

	addr := common.Bytes2Hex(input[0:20])
	number := new(big.Int).SetBytes(input[20:52])
	log.Infof("primitive_maintenance, address %s, block number %d", addr, number)

	// TODO: @AC get cpchain Backend and read balance.
	maintenance, err := c.Backend.Maintenance(common.HexToAddress(addr), number.Uint64())
	if err != nil {
		log.Fatal("NewBasicCollector,error", err)
	}
	ret := new(big.Int).SetInt64(int64(maintenance))
	return common.LeftPadBytes(ret.Bytes(), 32), nil
}
