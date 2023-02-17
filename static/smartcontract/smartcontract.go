package smartcontract

import (
	"github.com/blocklords/gosds/static/smartcontract/key"
)

type Smartcontract struct {
	// Body abi.ABI
	NetworkId               string `json:"network_id"`
	Address                 string `json:"address"`
	AbiHash                 string `json:"abi_hash"`
	Txid                    string `json:"transaction_id"`
	Deployer                string `json:"deployer"`
	PreDeployBlockNumber    uint64 `json:"pre_deploy_block_number"`
	PreDeployBlockTimestamp uint64 `json:"pre_deploy_block_timestamp"`
	exists                  bool
}

func (c *Smartcontract) Key() key.Key {
	return key.New(c.NetworkId, c.Address)
}

func (c *Smartcontract) SetExists(exists bool) {
	c.exists = exists
}