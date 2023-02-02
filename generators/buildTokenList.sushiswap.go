package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

func buildSushiswapTokenList() {
	tokenList := loadTokenListFromJsonFile(`sushiswap.json`)
	originalTokenList := helpers.FetchJSON[TokenListData](`https://token-list.sushi.com/`)
	tokenList.Name = originalTokenList.Name
	tokenList.LogoURI = originalTokenList.LogoURI
	tokenList.Keywords = originalTokenList.Keywords
	tokenList = tokenList.Assign(originalTokenList.Tokens)

	saveTokenListInJsonFile(tokenList, tokenList.Tokens, `sushiswap.json`, Standard)
}
