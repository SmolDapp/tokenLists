package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var BERACHAIN = TChain{
	ID:            80094,
	Name:          `Berachain`,
	Type:          `EVM`,
	LogoURI:       `https://assets.smold.app/chains/80094/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://rpc.berachain.com`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0,
	MulticallContract: TContractData{
		Address: `0xca11bde05977b3631167028862be2a173976ca11`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Bera`,
		Symbol:   `BERA`,
		LogoURI:  `https://assets.smold.app/api/token/80094/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  80094,
		Decimals: 18,
	},
	IgnoredTokens: []string{},
	ExtraTokens:   []string{},
}
