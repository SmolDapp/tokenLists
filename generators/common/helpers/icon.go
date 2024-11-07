package helpers

import (
	"os"
	"strconv"
	"strings"

	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/logs"
	"github.com/migratooor/tokenLists/generators/common/utils"
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
	basePath := `https://raw.githubusercontent.com/SmolDapp/tokenAssets/main/tokens/`
	for _, chainID := range chains.SUPPORTED_CHAIN_IDS {
		smoldAssetsPerChain[chainID] = FetchJSON[TSmolAssetsList](basePath + strconv.FormatUint(chainID, 10) + `/list.json`).Tokens
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

func UseIcon(chainID uint64, tokenName string, tokenAddress string, fallback string) string {
	tokenAddress = utils.ToAddress(tokenAddress)
	smolAssets := GetSmolAssetsPerChain(chainID)

	if IncludesAddress(smolAssets, tokenAddress) {
		return `https://assets.smold.app/api/token/` + strconv.FormatUint(chainID, 10) + `/` + tokenAddress + `/logo-128.png`
	}

	if shouldLogAssetError {
		logs.Info(`Missing icon for token ` + tokenName + ` (` + tokenAddress + `) on chain ` + strconv.FormatUint(chainID, 10))
	}
	if strings.Contains(fallback, `assets.coingecko.com`) && strings.Contains(fallback, `/thumb/`) {
		fallback = strings.Replace(fallback, `/thumb/`, `/large/`, 1)
	}
	if strings.Contains(fallback, `assets.coingecko.com`) && strings.Contains(fallback, `/small/`) {
		fallback = strings.Replace(fallback, `/small/`, `/large/`, 1)
	}

	if _, ok := ExistingTokenLogoURI[chainID]; ok {
		if existingLogo, ok := ExistingTokenLogoURI[chainID][tokenAddress]; ok {
			fallback = existingLogo
		}
	}
	if strings.Contains(fallback, DEFAULT_PARASWAP_NOT_FOUND) || strings.Contains(fallback, DEFAULT_ETHERSCAN_NOT_FOUND) {
		return DEFAULT_SMOL_NOT_FOUND
	}

	return fallback
}
