package chains

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/models"
)

var METIS = TChain{
	ID:            1088,
	Name:          `Metis`,
	LogoURI:       `https://assets.smold.app/chains/1088/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://metis-mainnet.public.blastapi.io`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xdeaddeaddeaddeaddeaddeaddeaddeaddead0000`,
		Name:     `Metis`,
		Symbol:   `METIS`,
		LogoURI:  `https://assets.smold.app/api/token/1088/0xdeaddeaddeaddeaddeaddeaddeaddeaddead0000/logo-128.png`,
		Decimals: 18,
		ChainID:  1088,
	},
	IgnoredTokens: []common.Address{},
}
