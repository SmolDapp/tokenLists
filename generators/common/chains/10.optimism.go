package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var OPTIMISM = TChain{
	ID:            10,
	Name:          `Optimism`,
	Type:          `EVM`,
	LogoURI:       `https://assets.smold.app/chains/10/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://mainnet.optimism.io`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.5,
	MulticallContract: TContractData{
		Address: `0xca11bde05977b3631167028862be2a173976ca11`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Ethereum`,
		Symbol:   `ETH`,
		LogoURI:  `https://assets.smold.app/api/token/10/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  1,
		Decimals: 18,
	},
	IgnoredTokens: []string{},
}
