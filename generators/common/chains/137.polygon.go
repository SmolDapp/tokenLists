package chains

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/models"
)

var POLYGON = TChain{
	ID:            137,
	RpcURI:        `https://polygon.llamarpc.com`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Matic`,
		Symbol:   `MATIC`,
		LogoURI:  `https://assets.smold.app/api/token/137/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  137,
		Decimals: 18,
	},
	IgnoredTokens: []common.Address{
		common.HexToAddress(`0xec6432B90e7fD4d9f872cc5C781f05B617DB861E`),
	},
}
