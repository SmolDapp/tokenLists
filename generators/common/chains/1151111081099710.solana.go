package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var SOLANA = TChain{
	ID:                1151111081099710,
	Name:              `Solana`,
	Type:              `SVM`,
	LogoURI:           `https://assets.smold.app/chains/1151111081099710/logo-128.png`,
	IsTestNet:         false,
	RpcURI:            `https://api.mainnet-beta.solana.com`,
	MaxBlockRange:     100_000_000,
	MaxBatchSize:      math.MaxInt64,
	WeightRatio:       0.0,
	MulticallContract: TContractData{},
	Coin: models.TokenListToken{
		Address:  `11111111111111111111111111111111`,
		Name:     `Solana`,
		Symbol:   `SOL`,
		LogoURI:  `https://assets.smold.app/api/token/1151111081099710/11111111111111111111111111111111/logo-128.png`,
		ChainID:  1151111081099710,
		Decimals: 9,
	},
	IgnoredTokens: []string{},
}
