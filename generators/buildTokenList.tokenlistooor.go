package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

var TOKENLISTOOOR_LISTS = []string{
	`paraswap`,
	`yearn`,
	`curve`,
	`optimism`,
}

func contains(arr []TokenListToken, value TokenListToken) bool {
	for _, v := range arr {
		if helpers.ToAddress(v.Address) == helpers.ToAddress(value.Address) {
			return true
		}
	}
	return false
}

func buildTokenListooorList_v1() {
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

func buildTokenListooorList() {
	tokenList := loadTokenListFromJsonFile(`tokenlistooor.json`)
	tokenList.Name = `Tokenlistooor Token List`
	tokenList.LogoURI = `https://raw.githubusercontent.com/Migratooor/tokenLists/main/.github/tokenlistooor.svg`
	tokenList.Description = `A curated list of tokens from all the token lists on tokenlistooor.`

	/**************************************************************************
	** Create a map of all tokens from all lists and only add the missing ones
	** in it. Map are WAY faster than arrays.
	**************************************************************************/
	allTokens := make(map[uint64]map[string]TokenListToken)
	allTokensPlain := []TokenListToken{}

	for name, generatorData := range GENERATORS {
		shouldByPassCount := name == `yearn`
		if generatorData.GeneratorType == GeneratorPool {
			continue
		}

		initialCount := 1
		if shouldByPassCount {
			initialCount = 5
		}
		tokenList := loadTokenListFromJsonFile(name + `.json`)
		for _, token := range tokenList.Tokens {
			if _, ok := allTokens[token.ChainID]; !ok {
				allTokens[token.ChainID] = make(map[string]TokenListToken)
			}
			if existingToken, ok := allTokens[token.ChainID][helpers.ToAddress(token.Address)]; ok {
				allTokens[token.ChainID][helpers.ToAddress(token.Address)] = TokenListToken{
					Address:  existingToken.Address,
					Name:     helpers.SafeString(existingToken.Name, token.Name),
					Symbol:   helpers.SafeString(existingToken.Symbol, token.Symbol),
					LogoURI:  helpers.SafeString(existingToken.LogoURI, token.LogoURI),
					Decimals: helpers.SafeInt(existingToken.Decimals, token.Decimals),
					ChainID:  token.ChainID,
					Count:    existingToken.Count + 1,
				}
			} else {
				allTokens[token.ChainID][helpers.ToAddress(token.Address)] = TokenListToken{
					Address:  helpers.ToAddress(token.Address),
					Name:     helpers.SafeString(token.Name, ``),
					Symbol:   helpers.SafeString(token.Symbol, ``),
					LogoURI:  helpers.SafeString(token.LogoURI, ``),
					Decimals: helpers.SafeInt(token.Decimals, 18),
					ChainID:  token.ChainID,
					Count:    initialCount,
				}
			}
		}
	}

	for _, tokens := range allTokens {
		for _, token := range tokens {
			if token.Count >= 5 {
				allTokensPlain = append(allTokensPlain, token)
			}
		}
	}

	saveTokenListInJsonFile(tokenList, allTokensPlain, `tokenlistooor.json`, Standard)
}
