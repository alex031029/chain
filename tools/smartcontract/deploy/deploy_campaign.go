// Copyright 2018 The cpchain authors
// This file is part of the cpchain library.
//
// The cpchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The cpchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the cpchain library. If not, see <http://www.gnu.org/licenses/>.

package deploy

import (
	"math/big"

	"bitbucket.org/cpchain/chain/commons/log"
	campaign "bitbucket.org/cpchain/chain/contracts/dpor/campaign"
	"bitbucket.org/cpchain/chain/tools/smartcontract/config"
	"github.com/ethereum/go-ethereum/common"
)

func DeployCampaign(acAddr common.Address, rewardAddr common.Address, password string, nonce uint64) common.Address {
	client, err, privateKey, _, fromAddress := config.Connect(password)
	printBalance(client, fromAddress)

	// Launch contract deploy transaction.
	auth := newTransactor(privateKey, new(big.Int).SetUint64(nonce))
	contractAddress, tx, _, err := campaign.DeployCampaign(auth, client, acAddr, rewardAddr)
	if err != nil {
		log.Fatal(err.Error())
	}
	printTx(tx, err, client, contractAddress)
	return contractAddress
}

func UpdateCampaignParameters(password string, campaignContractAddr common.Address, nonce1 uint64, nonce2 uint64) {
	client, err, privateKey, _, _ := config.Connect(password)
	// get chain config
	cfg, err := client.ChainConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	auth := newTransactor(privateKey, new(big.Int).SetUint64(nonce1))
	campaign, _ := campaign.NewCampaign(campaignContractAddr, client)
	tx, err := campaign.UpdateTermLen(auth, new(big.Int).SetUint64(cfg.Dpor.TermLen))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info("updated term len", "txhash", tx.Hash().Hex(), "termLen", cfg.Dpor.TermLen)
}
