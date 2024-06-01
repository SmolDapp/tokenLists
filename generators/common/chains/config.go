package chains

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/models"
)

type TContractData struct {
	Address common.Address // Address of the contract
	Block   uint64         // Block number where the contract was deployed
}
type TCoin struct {
	Address  common.Address `json:"address"`
	Name     string         `json:"name"`
	Symbol   string         `json:"symbol"`
	LogoURI  string         `json:"logoURI"`
	ChainID  uint64         `json:"chainId"`
	Decimals int            `json:"decimals"`
}
type TChain struct {
	ID                uint64
	RpcURI            string
	MaxBlockRange     uint64
	MaxBatchSize      uint64
	MulticallContract TContractData
	Coin              models.TokenListToken
	BlacklistedVaults []common.Address
	ExtraTokens       []common.Address
	IgnoredTokens     []common.Address
}

var DEFAULT_COIN_ADDRESS = common.HexToAddress(`0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`)

// CHAINS is the list of supported chains
var CHAINS = map[uint64]TChain{
	1:       ETHEREUM,
	5:       GOERLI,
	10:      OPTIMISM,
	56:      BINANCE_SMART_CHAIN,
	100:     GNOSIS,
	137:     POLYGON,
	250:     FANTOM,
	314:     FILECOIN,
	324:     ZKSYNC,
	1088:    METIS,
	1101:    POLYGON_ZKEVM,
	5000:    MANTLE,
	8453:    BASE,
	42161:   ARBITRUM,
	42220:   CELO,
	43114:   AVALANCHE,
	59144:   LINEA,
	81457:   BLAST,
	534352:  SCROLL,
	7777777: ZORA,
}

var SUPPORTED_CHAIN_IDS = []uint64{}

func init() {
	for k := range CHAINS {
		SUPPORTED_CHAIN_IDS = append(SUPPORTED_CHAIN_IDS, k)
	}
}

// IsChainIDSupported returns true if the chainID is supported by our program
func IsChainIDSupported(chainID uint64) bool {
	_, ok := CHAINS[chainID]
	return ok
}

func IsTokenIgnored(chainId uint64, address common.Address) bool {
	if address.Hex() == common.HexToAddress("0x0000000000000000000000000000000000000000").Hex() {
		return true
	}
	for _, ignoredToken := range CHAINS[chainId].IgnoredTokens {
		if ignoredToken.Hex() == address.Hex() {
			return true
		}
	}
	return false
}
