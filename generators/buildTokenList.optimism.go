package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

func buildOptimismTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`optimism.json`)
	originalTokenList := helpers.FetchJSON[models.TokenListData[models.TokenListToken]](`https://raw.githubusercontent.com/ethereum-optimism/ethereum-optimism.github.io/master/optimism.tokenlist.json`)
	tokenList.Name = helpers.SafeString(originalTokenList.Name, `Optimism Token List`)
	tokenList.LogoURI = helpers.SafeString(originalTokenList.LogoURI, `https://ethereum-optimism.github.io/optimism.svg`)
	tokenList.Keywords = originalTokenList.Keywords

	tokenList = models.TokenListData[models.TokenListToken]{}
	for _, token := range originalTokenList.Tokens {
		if !chains.IsChainIDSupported(token.ChainID) {
			continue
		}
		if (token.Name == `` || token.Symbol == `` || token.Decimals == 0) || chains.IsTokenIgnored(token.ChainID, common.HexToAddress(token.Address)) {
			continue
		}
		key := helpers.GetKey(token.ChainID, common.HexToAddress(token.Address))
		tokenList.NextTokensMap[key] = token
	}

	tokens := helpers.GetTokensFromList(tokenList.Tokens)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `optimism.json`, helpers.SavingMethodStandard)
}
