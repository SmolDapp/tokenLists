package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

func buildUniswapTokenList() {
	tokenList := loadTokenListFromJsonFile(`uniswap.json`)
	originalTokenList := helpers.FetchJSON[TokenListData](`https://tokens.uniswap.org`)
	tokenList.Name = helpers.SafeString(originalTokenList.Name, `Uniswap Token List`)
	tokenList.LogoURI = helpers.SafeString(originalTokenList.LogoURI, `ipfs://QmNa8mQkrNKp1WEEeGjFezDmDeodkWRevGFN8JCV7b4Xir"`)
	tokenList.Keywords = originalTokenList.Keywords
	tokenList = tokenList.Assign(originalTokenList.Tokens)

	saveTokenListInJsonFile(tokenList, tokenList.Tokens, `uniswap.json`, Standard)
}
