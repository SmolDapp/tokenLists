package chains

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/models"
)

var FILECOIN = TChain{
	ID:            314,
	Name:          `Filecoin`,
	LogoURI:       `https://assets.smold.app/chains/314/logo-128.png`,
	IsTestNet:     false,
	RpcURI:        `https://rpc.ankr.com/filecoin`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0xcA11bde05977b3631167028862bE2a173976CA11`),
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Filecoin`,
		Symbol:   `FIL`,
		LogoURI:  `https://assets.smold.app/api/token/314/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  314,
		Decimals: 18,
	},
	IgnoredTokens: []common.Address{},
}
