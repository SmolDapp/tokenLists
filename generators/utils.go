package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/utils"
)

func loadAllTokenLogoURI() map[uint64]map[string]string {
	allTokenLogoURI := make(map[uint64]map[string]string)
	for name := range GENERATORS {
		tokenList := helpers.LoadTokenListFromJsonFile(name + `.json`)
		for _, token := range tokenList.Tokens {
			if _, ok := allTokenLogoURI[token.ChainID]; !ok {
				allTokenLogoURI[token.ChainID] = make(map[string]string)
			}
			currentIcon := allTokenLogoURI[token.ChainID][utils.ToAddress(token.Address)]
			if currentIcon == helpers.DEFAULT_SMOL_NOT_FOUND ||
				currentIcon == helpers.DEFAULT_PARASWAP_NOT_FOUND ||
				currentIcon == helpers.DEFAULT_ETHERSCAN_NOT_FOUND ||
				currentIcon == `` {
				baseIcon := helpers.UseIcon(token.ChainID, token.Name+` - `+token.Symbol, token.Address, token.LogoURI)
				allTokenLogoURI[token.ChainID][utils.ToAddress(token.Address)] = baseIcon
			}

		}
	}
	helpers.ExistingTokenLogoURI = allTokenLogoURI
	return allTokenLogoURI
}
