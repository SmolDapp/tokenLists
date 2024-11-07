package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var GNOSIS = TChain{
	ID:            100,
	Name:          `Gnosis`,
	Type:          `EVM`,
	LogoURI:       `https://assets.smold.app/chains/100/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://rpc.gnosis.gateway.fm`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.3,
	MulticallContract: TContractData{
		Address: `0xca11bde05977b3631167028862be2a173976ca11`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `xDai`,
		Symbol:   `xDAI`,
		LogoURI:  `https://assets.smold.app/api/token/100/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  100,
		Decimals: 18,
	},
	IgnoredTokens: []string{},
}
