package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var INK = TChain{
	ID:            57073,
	Name:          `Ink`,
	Type:          `EVM`,
	LogoURI:       `https://assets.smold.app/chains/57073/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://rpc-gel.inkonchain.com`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.0,
	MulticallContract: TContractData{
		Address: `0xca11bde05977b3631167028862be2a173976ca11`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Ethereum`,
		Symbol:   `ETH`,
		LogoURI:  `https://assets.smold.app/api/token/57073/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  57073,
		Decimals: 18,
	},
	IgnoredTokens: []string{},
}
