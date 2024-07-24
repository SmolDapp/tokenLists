package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/static"
)

func handleStaticTokenList(chainID uint64, tokens []static.TStaticElement) []models.TokenListToken {
	tokenAddresses := []common.Address{}
	for _, token := range tokens {
		tokenAddresses = append(tokenAddresses, token.Address)
	}
	tokenList := helpers.GetTokensFromAddresses(chainID, tokenAddresses)
	tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
	return tokenList
}

func fetchAjnaStaticTokenList(chainID uint64) []models.TokenListToken {
	tokens := static.AJNA_STATIC_TOKENLIST
	return handleStaticTokenList(chainID, tokens[chainID])
}

func buildAjnaStaticTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`ajna-static.json`)
	tokenList.Name = `Ajna (Static)`
	tokenList.LogoURI = `https://www.ajna.finance/static/tokens/ajna.png`
	tokenList.Keywords = []string{`Ajna`}
	tokens := []models.TokenListToken{}
	tokens = append(tokens, fetchAjnaStaticTokenList(1)...)
	tokens = append(tokens, fetchAjnaStaticTokenList(5)...)
	tokens = append(tokens, fetchAjnaStaticTokenList(10)...)
	tokens = append(tokens, fetchAjnaStaticTokenList(137)...)
	tokens = append(tokens, fetchAjnaStaticTokenList(42161)...)
	tokens = append(tokens, fetchAjnaStaticTokenList(1380012617)...)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `ajna-static.json`, helpers.SavingMethodStandard)
}
