package chains

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/models"
)

var CELO = TChain{
	ID:            42220,
	Name:          `Celo`,
	LogoURI:       `https://assets.smold.app/chains/42220/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://1rpc.io/celo`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	WeightRatio:   0.5,
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xdeaddeaddeaddeaddeaddeaddeaddeaddead0000`,
		Name:     `CELO`,
		Symbol:   `CELO`,
		LogoURI:  `https://assets.smold.app/api/token/42220/0xdeaddeaddeaddeaddeaddeaddeaddeaddead0000/logo-128.png`,
		Decimals: 18,
		ChainID:  42220,
	},
	IgnoredTokens: []common.Address{},
}
