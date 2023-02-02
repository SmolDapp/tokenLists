package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

func buildCowswapTokenList() {
	tokenList := loadTokenListFromJsonFile(`cowswap.json`)
	originalTokenList := helpers.FetchJSON[TokenListData](`https://raw.githubusercontent.com/cowprotocol/token-lists/main/src/public/CowSwap.json`)
	tokenList.Name = originalTokenList.Name
	tokenList.LogoURI = originalTokenList.LogoURI
	tokenList.Keywords = originalTokenList.Keywords
	tokenList = tokenList.Assign(originalTokenList.Tokens)

	saveTokenListInJsonFile(tokenList, tokenList.Tokens, `cowswap.json`, Standard)
}
