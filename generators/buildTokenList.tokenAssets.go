package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

func handleSmolAssetsTokenList(chainID uint64, tokenAddresses []common.Address) []models.TokenListToken {
	tokenList := helpers.GetTokensFromAddresses(chainID, tokenAddresses)
	tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
	return tokenList
}

func fetchSmolAssetsTokenList(chainID uint64) []models.TokenListToken {
	smolAssets := helpers.GetSmolAssetsPerChain(chainID)
	allTokensToAdd := []common.Address{}
	for _, token := range smolAssets {
		allTokensToAdd = append(allTokensToAdd, common.HexToAddress(token))
	}
	return handleSmolAssetsTokenList(chainID, allTokensToAdd)
}

func buildSmolAssetsTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`smolAssets.json`)
	tokenList.Name = `SmolAssets`
	tokenList.Description = `A list of tokens supported by Smoldapp Token Assets repository`
	tokenList.LogoURI = `https://raw.githubusercontent.com/Migratooor/tokenLists/main/.github/tokenlistooor.svg`
	tokenList.Keywords = []string{`smol`, `tokenAssets`}
	tokens := []models.TokenListToken{}
	for _, chainID := range chains.SUPPORTED_CHAIN_IDS {
		tokens = append(tokens, fetchSmolAssetsTokenList(chainID)...)
	}
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `smolAssets.json`, helpers.SavingMethodStandard)
}
