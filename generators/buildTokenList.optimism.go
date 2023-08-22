package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

func buildOptimismTokenList() {
	tokenList := loadTokenListFromJsonFile(`optimism.json`)
	originalTokenList := helpers.FetchJSON[TokenListData[TokenListToken]](`https://raw.githubusercontent.com/ethereum-optimism/ethereum-optimism.github.io/master/optimism.tokenlist.json`)
	tokenList.Name = helpers.SafeString(originalTokenList.Name, `Optimism Token List`)
	tokenList.LogoURI = helpers.SafeString(originalTokenList.LogoURI, `https://ethereum-optimism.github.io/optimism.svg`)
	tokenList.Keywords = originalTokenList.Keywords
	tokenList = tokenList.Assign(originalTokenList.Tokens)

	saveTokenListInJsonFile(tokenList, tokenList.Tokens, `optimism.json`, Standard)
}
