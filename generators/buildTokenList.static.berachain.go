package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/static"
)

func fetchBerachainStaticTokenList(chainID uint64) []models.TokenListToken {
	tokens := static.BERACHAIN_STATIC_TOKENLIST
	return handleStaticTokenList(chainID, tokens[chainID])
}

func buildBerachainStaticTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`berachain-static.json`)
	tokenList.Name = `Berachain (Static)`
	tokenList.Description = `An EVM-identical L1 aligning security and liquidity powered by Proof Of Liquidity`
	tokenList.LogoURI = `https://assets.smold.app/chains/80094/logo-128.png`
	tokenList.Keywords = []string{`Berachain`}
	tokens := []models.TokenListToken{}
	tokens = append(tokens, fetchBerachainStaticTokenList(80094)...)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `berachain-static.json`, helpers.SavingMethodStandard)
}
