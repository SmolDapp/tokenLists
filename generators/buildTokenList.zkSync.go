package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/static"
)

func handleZkSyncTokenList(chainID uint64, tokens []static.TStaticElement) []models.TokenListToken {
	tokenAddresses := []common.Address{}
	for _, token := range tokens {
		tokenAddresses = append(tokenAddresses, token.Address)
	}
	tokenList := helpers.GetTokensFromAddresses(chainID, tokenAddresses)
	tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
	return tokenList
}

func fetchZkSyncTokenList() []models.TokenListToken {
	tokens := static.ZKSYNC_STATIC_TOKENLIST
	return handleZkSyncTokenList(324, tokens[324])
}

func buildZkSyncTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`zksync.json`)
	tokenList.Name = `zkSync`
	tokenList.LogoURI = `https://assets.smold.app/api/chain/324/logo-128.png`
	tokenList.Keywords = []string{`zksync`}
	tokens := []models.TokenListToken{}
	tokens = append(tokens, fetchZkSyncTokenList()...)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `zksync.json`, helpers.SavingMethodStandard)
}
