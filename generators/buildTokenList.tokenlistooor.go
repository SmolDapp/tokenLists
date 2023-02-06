package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

var TOKENLISTOOOR_LISTS = []string{
	`paraswap`,
	`yearn`,
	`curve`,
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

	/**************************************************************************
	** Create a map of all tokens from all lists and only add the missing ones
	** in it. Map are WAY faster than arrays.
	**************************************************************************/
	tokenListMap := make(map[string]TokenListToken)
	for _, name := range TOKENLISTOOOR_LISTS {
		sourceTokenList := loadTokenListFromJsonFile(name + `.json`)
		for _, token := range sourceTokenList.Tokens {
			if data, ok := tokenListMap[helpers.ToAddress(token.Address)]; ok {
				data.LogoURI = helpers.SafeString(data.LogoURI, token.LogoURI)
				tokenListMap[helpers.ToAddress(token.Address)] = data
				continue
			}
			tokenListMap[helpers.ToAddress(token.Address)] = token
		}
	}

	/**************************************************************************
	** Transform the map into an array to be able to save it correctly in the
	** JSON file.
	**************************************************************************/
	tokens := []TokenListToken{}
	for _, token := range tokenListMap {
		tokens = append(tokens, token)
	}
	saveTokenListInJsonFile(tokenList, tokens, `tokenlistooor.json`, Standard)
}
