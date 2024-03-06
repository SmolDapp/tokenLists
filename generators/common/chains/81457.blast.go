package chains

import (
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/models"
)

var BLAST = TChain{
	ID:            81457,
	RpcURI:        `https://rpc.blast.io`,
	MaxBlockRange: 100_000_000,
	MaxBatchSize:  math.MaxInt64,
	MulticallContract: TContractData{
		Address: common.HexToAddress(`0xca11bde05977b3631167028862be2a173976ca11`),
		Block:   0,
	},
	Coin: models.TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Ethereum`,
		Symbol:   `ETH`,
		LogoURI:  `https://assets.smold.app/api/token/81457/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`,
		ChainID:  81457,
		Decimals: 18,
	},
	IgnoredTokens: []common.Address{},
	ExtraTokens: []common.Address{
		common.HexToAddress(`0x6d5564584b70240691bd6ff7a834b9fab844e0d4`),
		common.HexToAddress(`0x38aD23b0902D0d86c2F3949BC505194D70B762F5`),
	},
}
