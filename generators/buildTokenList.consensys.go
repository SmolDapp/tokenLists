package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

func buildConsensysTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`consensys.json`)
	originalTokenList := helpers.FetchJSON[models.TokenListData[models.TokenListToken]](
		`https://raw.githubusercontent.com/Consensys/linea-token-list/main/json/linea-mainnet-token-shortlist.json`,
	)
	tokenList.Name = originalTokenList.Name
	tokenList.LogoURI = `https://avatars.githubusercontent.com/u/10818037?s=200&v=4`
	tokenList.Keywords = originalTokenList.Keywords
	buildExternalTokenList(tokenList, originalTokenList.Tokens, `consensys.json`)
}
