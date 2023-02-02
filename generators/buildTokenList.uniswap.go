package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

func buildUniswapTokenList() {
	tokenList := loadTokenListFromJsonFile(`uniswap.json`)
	originalTokenList := helpers.FetchJSON[TokenListData](`https://tokens.uniswap.org`)
	tokenList.Name = originalTokenList.Name
	tokenList.LogoURI = originalTokenList.LogoURI
	tokenList.Keywords = originalTokenList.Keywords
	tokenList = tokenList.Assign(originalTokenList.Tokens)

	saveTokenListInJsonFile(tokenList, tokenList.Tokens, `uniswap.json`, Standard)
}
