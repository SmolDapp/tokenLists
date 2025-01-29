package chains

import (
	"math"

	"github.com/migratooor/tokenLists/generators/common/models"
)

var SOPHON = TChain{
	ID:            50104,
	Name:          `Sophon`,
	Type:          `EVM`,
	LogoURI:       `https://assets.smold.app/chains/50104/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://rpc.sophon.xyz`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.0,
	MulticallContract: TContractData{
		Address: `0x5f4867441d2416cA88B1b3fd38f21811680CD2C8`,
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `SOPH`,
		Symbol:   `SOPH`,
		LogoURI:  `https://assets.smold.app/api/token/50104/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  50104,
		Decimals: 18,
	},
	IgnoredTokens: []string{},
}
