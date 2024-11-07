package main

import (
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/common/utils"
)

func handleSmolAssetsTokenList(chainID uint64, tokenAddresses []string) []models.TokenListToken {
	tokenList := helpers.GetTokensFromAddresses(chainID, tokenAddresses)
	if len(tokenList) == 0 {
		return tokenList
	}
	tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
	return tokenList
}

func fetchSmolAssetsTokenList(chainID uint64) []models.TokenListToken {
	smolAssets := helpers.GetSmolAssetsPerChain(chainID)
	allTokensToAdd := []string{}
	for _, token := range smolAssets {
		allTokensToAdd = append(allTokensToAdd, utils.ToAddress(token))
	}
	return handleSmolAssetsTokenList(chainID, allTokensToAdd)
}

func buildSmolAssetsTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`smolAssets.json`)
	tokenList.Name = `SmolAssets`
	tokenList.Description = `A list of tokens supported by Smoldapp Token Assets repository`
	tokenList.LogoURI = `https://raw.githubusercontent.com/smoldapp/tokenLists/main/.github/tokenlistooor.svg`
	tokenList.Keywords = []string{`smol`, `tokenAssets`}
	tokens := []models.TokenListToken{}
	for _, chainID := range chains.SUPPORTED_CHAIN_IDS {
		tokens = append(tokens, fetchSmolAssetsTokenList(chainID)...)
	}
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `smolAssets.json`, helpers.SavingMethodStandard)
}
