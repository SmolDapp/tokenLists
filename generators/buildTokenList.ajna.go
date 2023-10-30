package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/static"
)

func handleAjnaTokenList(chainID uint64, tokens []static.TStaticElement) []TokenListToken {
	tokenList := []TokenListToken{}
	tokenAddresses := []common.Address{}
	for _, token := range tokens {
		tokenAddresses = append(tokenAddresses, token.Address)
	}

	tokensInfo := retrieveBasicInformations(chainID, tokenAddresses)
	for index, address := range tokenAddresses {
		if token, ok := tokensInfo[address.Hex()]; ok {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				tokens[index].Icon,
				chainID,
				int(token.Decimals),
			); err == nil {
				tokenList = append(tokenList, newToken)
			}
		}
	}
	return tokenList
}

func fetchAjnaTokenList(chainID uint64) []TokenListToken {
	tokens := static.AJNA_STATIC_TOKENLIST
	return handleAjnaTokenList(chainID, tokens[chainID])
}

func buildAjnaTokenList() {
	tokenList := loadTokenListFromJsonFile(`ajna.json`)
	tokenList.Name = `Ajna`
	tokenList.LogoURI = `https://www.ajna.finance/static/tokens/ajna.png`
	tokenList.Keywords = []string{`Ajna`}
	tokens := []TokenListToken{}
	tokens = append(tokens, fetchAjnaTokenList(1)...)
	tokens = append(tokens, fetchAjnaTokenList(5)...)
	tokens = append(tokens, fetchAjnaTokenList(137)...)
	saveTokenListInJsonFile(tokenList, tokens, `ajna.json`, Standard)
}
