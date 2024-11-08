package main

import (
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

type TJupiterFMToken struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	LogoURI  string `json:"logoURI"`
	Decimals int    `json:"decimals"`
}

var JupiterTokenEndpoint = `https://tokens.jup.ag/tokens?tags=verified`

func fetchJupiterTokenList() []models.TokenListToken {
	tokenLists := []models.TokenListToken{}

	tokenAddresses := []string{}
	tokenIcons := map[string]string{}
	fetchingURL := JupiterTokenEndpoint
	list := helpers.FetchJSON[[]TJupiterFMToken](fetchingURL)

	for _, token := range list {
		tokenAddresses = append(tokenAddresses, token.Address)
		tokenIcons[token.Address] = token.LogoURI
	}
	tokenLists = helpers.GetTokensFromAddressesWithIcons(chains.SOLANA.ID, tokenAddresses, tokenIcons)
	return tokenLists
}

func buildJupiterTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`jupiter.json`)
	tokenList.Name = "Jupiter Verified Token List"
	tokenList.LogoURI = "https://station.jup.ag/img/jupiter-logo.svg"

	tokens := fetchJupiterTokenList()
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `jupiter.json`, helpers.SavingMethodStandard)
}
