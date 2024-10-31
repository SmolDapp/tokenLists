package main

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
	"github.com/migratooor/tokenLists/generators/common/models"
)

type TLifiTokenData struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	LogoURI  string `json:"logoURI"`
	ChainID  uint64 `json:"chainId"`
	Decimals int    `json:"decimals"`
}
type TLifiList struct {
	Tokens map[string]TLifiTokenData
}

var lifiBaseURL = `https://raw.githubusercontent.com/lifinance/customized-token-list/refs/heads/main/tokens/`
var GithubURLForLifi = map[uint64]string{
	1:     lifiBaseURL + `ETH.json`,
	10:    lifiBaseURL + `OPT.json`,
	56:    lifiBaseURL + `BSC.json`,
	100:   lifiBaseURL + `DAI.json`,
	122:   lifiBaseURL + `FUS.json`,
	137:   lifiBaseURL + `POL.json`,
	250:   lifiBaseURL + `FTM.json`,
	324:   lifiBaseURL + `ERA.json`,
	1101:  lifiBaseURL + `PZE.json`,
	1284:  lifiBaseURL + `MOO.json`,
	1285:  lifiBaseURL + `MOR.json`,
	1329:  lifiBaseURL + `SEI.json`,
	8453:  lifiBaseURL + `BAS.json`,
	13371: lifiBaseURL + `IMX.json`,
	42161: lifiBaseURL + `ARB.json`,
	42220: lifiBaseURL + `CEL.json`,
	43114: lifiBaseURL + `AVA.json`,
	59144: lifiBaseURL + `LNA.json`,
}

func fetchLifiTokenList() []models.TokenListToken {
	tokenLists := []models.TokenListToken{}

	for chainID, uri := range GithubURLForLifi {
		if !chains.IsChainIDSupported(chainID) {
			continue
		}

		logs.Info(`Fetching Lifi token list for chainID: ` + strconv.FormatUint(chainID, 10))
		list := helpers.FetchJSON[[]TLifiTokenData](uri)
		tokenAddresses := []common.Address{}
		for _, token := range list {
			tokenAddresses = append(tokenAddresses, common.HexToAddress(token.Address))
		}
		tokenList := helpers.GetTokensFromAddresses(chainID, tokenAddresses)
		tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
		tokenLists = append(tokenLists, tokenList...)
	}

	return tokenLists
}

func buildLifiTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`lifi.json`)
	tokenList.Name = "Lifi Token List"
	tokenList.LogoURI = "https://avatars.githubusercontent.com/u/85288935?s=200&v=4"

	tokens := fetchLifiTokenList()
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `lifi.json`, helpers.SavingMethodStandard)
}
