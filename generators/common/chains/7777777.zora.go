package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var ZORA = TChain{
	ID:            7777777,
	Name:          `Zora`,
	Type:          `EVM`,
	LogoURI:       `https://assets.smold.app/chains/7777777/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://rpc.zora.energy`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.05,
	MulticallContract: TContractData{
		Address: `0xca11bde05977b3631167028862be2a173976ca11`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Ethereum`,
		Symbol:   `ETH`,
		LogoURI:  `https://assets.smold.app/api/token/7777777/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  7777777,
		Decimals: 18,
	},
	IgnoredTokens: []string{},
}
