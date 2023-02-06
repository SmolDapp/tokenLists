package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

var TOKENLISTOOOR_LISTS = []string{
	`coingecko`,
	`1inch`,
	`paraswap`,
	`defillama`,
	`yearn`,
	`curve`,
	`cowswap`,
	`uniswap`,
	`sushiswap`,
	`ledger`,
	`portals`,
	`wido`,
}

func contains(arr []TokenListToken, value TokenListToken) bool {
	for _, v := range arr {
		if helpers.ToAddress(v.Address) == helpers.ToAddress(value.Address) {
			return true
		}
	}
	return false
}

func buildTokenListooorList() {
	tokenList := loadTokenListFromJsonFile(`tokenlistooor.json`)
	tokenList.Name = `Tokenlistooor Token List`
	tokenList.LogoURI = `https://raw.githubusercontent.com/Migratooor/tokenLists/main/.github/tokenlistooor.svg`

	for _, name := range TOKENLISTOOOR_LISTS {
		sourceTokenList := loadTokenListFromJsonFile(name + `.json`)
		for _, token := range sourceTokenList.Tokens {
			if contains(tokenList.Tokens, token) {
				continue
			}
			tokenList.Tokens = append(tokenList.Tokens, token)
		}
	}

	saveTokenListInJsonFile(tokenList, tokenList.Tokens, `tokenlistooor.json`, Standard)
}
