package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var BASE = TChain{
	ID:            8453,
	Name:          `Base`,
	Type:          `EVM`,
	LogoURI:       `https://assets.smold.app/chains/8453/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://mainnet.base.org/`,
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
		LogoURI:  `https://assets.smold.app/api/token/8453/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  8453,
		Decimals: 18,
	},
	IgnoredTokens: []string{},
}
