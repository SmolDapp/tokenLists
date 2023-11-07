package helpers

import (
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

const DEFAULT_SMOL_NOT_FOUND = `https://assets.smold.app/not-found.png`
const DEFAULT_PARASWAP_NOT_FOUND = `https://cdn.paraswap.io/token/token.png`
const DEFAULT_ETHERSCAN_NOT_FOUND = `https://etherscan.io/images/main/empty-token.png`

var ExistingTokenLogoURI = make(map[uint64]map[string]string)
var smoldAssetsPerChain = make(map[uint64][]string)
var shouldLogAssetError = false

type TSmolAssetsList struct {
	Version struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	}
	Tokens []string `json:"tokens"`
}

func init() {
	for chainID := range SUPPORTED_CHAIN_IDS {
		smoldAssetsPerChain[chainID] = FetchJSON[TSmolAssetsList](`https://assets.smold.app/tokens/` + strconv.FormatUint(chainID, 10) + `/list.json`).Tokens
	}
	cmdArgs := os.Args
	for _, arg := range cmdArgs {
		if arg == "--log-assets-error" {
			shouldLogAssetError = true
		}
	}
}

func GetSmolAssetsPerChain(chainID uint64) []string {
	return smoldAssetsPerChain[chainID]
}

func UseIcon(chainID uint64, tokenName string, tokenAddress common.Address, fallback string) string {
	smolAssets := GetSmolAssetsPerChain(chainID)
	if IncludesAddress(smolAssets, tokenAddress) {
		return `https://assets.smold.app/api/token/` + strconv.FormatUint(chainID, 10) + `/` + tokenAddress.Hex() + `/logo-128.png`
	}

	if shouldLogAssetError {
		logs.Info(`Missing icon for token ` + tokenName + ` (` + tokenAddress.Hex() + `) on chain ` + strconv.FormatUint(chainID, 10))
	}
	if strings.Contains(fallback, `assets.coingecko.com`) && strings.Contains(fallback, `/thumb/`) {
		fallback = strings.Replace(fallback, `/thumb/`, `/large/`, 1)
	}
	if strings.Contains(fallback, `assets.coingecko.com`) && strings.Contains(fallback, `/small/`) {
		fallback = strings.Replace(fallback, `/small/`, `/large/`, 1)
	}

	if _, ok := ExistingTokenLogoURI[chainID]; ok {
		if existingLogo, ok := ExistingTokenLogoURI[chainID][tokenAddress.Hex()]; ok {
			fallback = existingLogo
		}
	}
	if strings.Contains(fallback, DEFAULT_PARASWAP_NOT_FOUND) || strings.Contains(fallback, DEFAULT_ETHERSCAN_NOT_FOUND) {
		return DEFAULT_SMOL_NOT_FOUND
	}

	return fallback
}
