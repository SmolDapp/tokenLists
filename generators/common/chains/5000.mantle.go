package chains

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/models"
)

var MANTLE = TChain{
	ID:            5000,
	Name:          `Mantle`,
	LogoURI:       `https://assets.smold.app/chains/5000/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://rpc.mantle.xyz`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xdeaddeaddeaddeaddeaddeaddeaddeaddead0000`,
		Name:     `Mantle`,
		Symbol:   `MNT`,
		LogoURI:  `https://assets.smold.app/api/token/5000/0xdeaddeaddeaddeaddeaddeaddeaddeaddead0000/logo-128.png`,
		Decimals: 18,
		ChainID:  5000,
	},
	IgnoredTokens: []common.Address{},
}
