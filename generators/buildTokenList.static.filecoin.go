package main

import (
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/static"
)

func fetchFilecoinStaticTokenList(chainID uint64) []models.TokenListToken {
	tokens := static.FILECOIN_STATIC
	return handleStaticTokenList(chainID, tokens[chainID])
}

func buildFilecoinStaticTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`filecoin-static.json`)
	tokenList.Name = `Filecoin (Static)`
	tokenList.LogoURI = chains.CHAINS[314].LogoURI
	tokenList.Keywords = []string{`filecoin`}
	tokens := []models.TokenListToken{}
	tokens = append(tokens, fetchFilecoinStaticTokenList(314)...)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `filecoin-static.json`, helpers.SavingMethodStandard)
}
