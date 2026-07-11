package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

func buildOptimismTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`optimism.json`)
	originalTokenList := helpers.FetchJSON[models.TokenListData[models.TokenListToken]](`https://raw.githubusercontent.com/ethereum-optimism/ethereum-optimism.github.io/master/optimism.tokenlist.json`)
	tokenList.Name = helpers.SafeString(originalTokenList.Name, `Optimism Token List`)
	tokenList.LogoURI = helpers.SafeString(originalTokenList.LogoURI, `https://ethereum-optimism.github.io/optimism.svg`)
	tokenList.Keywords = originalTokenList.Keywords
	buildExternalTokenList(tokenList, originalTokenList.Tokens, `optimism.json`)
}
