package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

type T1InchTokenData struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	LogoURI  string `json:"logoURI"`
	Decimals int    `json:"decimals"`
}
type T1InchList struct {
	Tokens map[string]T1InchTokenData
}

var APIURIFor1Inch = map[uint64]string{
	1:     `https://api.1inch.dev/token/v1.2/1/token-list`,
	10:    `https://api.1inch.dev/token/v1.2/10/token-list`,
	56:    `https://api.1inch.dev/token/v1.2/56/token-list`,
	100:   `https://api.1inch.dev/token/v1.2/100/token-list`,
	137:   `https://api.1inch.dev/token/v1.2/137/token-list`,
	250:   `https://api.1inch.dev/token/v1.2/250/token-list`,
	42161: `https://api.1inch.dev/token/v1.2/42161/token-list`,
	43114: `https://api.1inch.dev/token/v1.2/43114/token-list`,
}

func fetch1InchTokenList() []models.TokenListToken {
	tokenLists := []models.TokenListToken{}

	for chainID, uri := range APIURIFor1Inch {
		if !chains.IsChainIDSupported(chainID) {
			continue
		}

		list := helpers.FetchJSON[models.TokenListData[models.TokenListToken]](uri)
		tokenAddresses := []common.Address{}
		for _, token := range list.Tokens {
			tokenAddresses = append(tokenAddresses, common.HexToAddress(token.Address))
		}
		tokenList := helpers.GetTokensFromAddresses(chainID, tokenAddresses)
		tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
		tokenLists = append(tokenLists, tokenList...)
	}

	return tokenLists
}

func build1InchTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`1inch.json`)
	tokenList.Name = "1inch Token List"
	tokenList.LogoURI = "https://app.1inch.io/assets/images/logo.png"

	tokens := fetch1InchTokenList()
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `1inch.json`, helpers.SavingMethodStandard)
}
