package chains

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/models"
)

var ZKSYNC = TChain{
	ID:            324,
	Name:          `ZKSync`,
	LogoURI:       `https://assets.smold.app/chains/324/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://mainnet.era.zksync.io`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.1,
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0xF9cda624FBC7e059355ce98a31693d299FACd963`),
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Ethereum`,
		Symbol:   `ETH`,
		LogoURI:  `https://assets.smold.app/api/token/324/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  324,
		Decimals: 18,
	},
	IgnoredTokens: []common.Address{},
}
