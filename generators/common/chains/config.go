package chains

import (
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/models"
)

type TContractData struct {
	Address string // Address of the contract
	Block   uint64 // Block number where the contract was deployed
}
type TCoin struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	LogoURI  string `json:"logoURI"`
	ChainID  uint64 `json:"chainId"`
	Decimals int    `json:"decimals"`
}
type TChain struct {
	ID                uint64
	Name              string
	Type              string
	LogoURI           string
	RpcURI            string
	MaxBlockRange     uint64
	MaxBatchSize      uint64
	WeightRatio       float64
	IsTestNet         bool
	MulticallContract TContractData
	Coin              models.TokenListToken
	BlacklistedVaults []string
	ExtraTokens       []string
	IgnoredTokens     []string
}

var DEFAULT_COIN_ADDRESS = common.HexToAddress(`0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`)

// CHAINS is the list of supported chains
var CHAINS = map[uint64]TChain{
	1:                ETHEREUM,
	5:                GOERLI,
	10:               OPTIMISM,
	56:               BINANCE_SMART_CHAIN,
	100:              GNOSIS,
	137:              POLYGON,
	250:              FANTOM,
	314:              FILECOIN,
	324:              ZKSYNC,
	1088:             METIS,
	1101:             POLYGON_ZKEVM,
	5000:             MANTLE,
	8453:             BASE,
	42161:            ARBITRUM,
	42220:            CELO,
	43114:            AVALANCHE,
	50104:            SOPHON,
	59144:            LINEA,
	80094:            BERACHAIN,
	81457:            BLAST,
	534352:           SCROLL,
	7777777:          ZORA,
	1380012617:       RARI,
	1151111081099710: SOLANA,
	57073:            INK,
	747474:           KATANA,
}

var SUPPORTED_CHAIN_IDS = []uint64{}

func init() {
	for k := range CHAINS {
		SUPPORTED_CHAIN_IDS = append(SUPPORTED_CHAIN_IDS, k)
	}
	//Sort the chain IDs
	sort.Slice(SUPPORTED_CHAIN_IDS, func(i, j int) bool { return SUPPORTED_CHAIN_IDS[i] < SUPPORTED_CHAIN_IDS[j] })
}

// IsChainIDSupported returns true if the chainID is supported by our program
func IsChainIDSupported(chainID uint64) bool {
	_, ok := CHAINS[chainID]
	return ok
}

func IsTokenIgnored(chainId uint64, address string) bool {
	if strings.EqualFold(address, "0x0000000000000000000000000000000000000000") {
		return true
	}
	for _, ignoredToken := range CHAINS[chainId].IgnoredTokens {
		if strings.EqualFold(ignoredToken, address) {
			return true
		}
	}
	return false
}
