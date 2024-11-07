package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var FANTOM = TChain{
	ID:            250,
	Name:          `Fantom`,
	Type:          `EVM`,
	LogoURI:       `https://assets.smold.app/chains/250/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://rpc.ftm.tools`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.5,
	MulticallContract: TContractData{
		Address: `0x470ADB45f5a9ac3550bcFFaD9D990Bf7e2e941c9`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Fantom`,
		Symbol:   `FTM`,
		LogoURI:  `https://assets.smold.app/api/token/250/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  250,
		Decimals: 18,
	},
	IgnoredTokens: []string{},
}
