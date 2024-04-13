package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

func buildConsensysTokenList() {
	tokenList := loadTokenListFromJsonFile(`consensys.json`)
	originalTokenList := helpers.FetchJSON[TokenListData[TokenListToken]](`https://raw.githubusercontent.com/Consensys/linea-token-list/main/json/linea-mainnet-token-shortlist.json`)
	tokenList.Name = originalTokenList.Name
	tokenList.LogoURI = `https://avatars.githubusercontent.com/u/10818037?s=200&v=4`
	tokenList.Keywords = originalTokenList.Keywords
	tokenList = tokenList.Assign(originalTokenList.Tokens)

	tokens := fetchTokenList(tokenList.Tokens)
	saveTokenListInJsonFile(tokenList, tokens, `consensys.json`, Standard)
}
