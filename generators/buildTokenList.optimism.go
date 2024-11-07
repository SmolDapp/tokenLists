package main

import (
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/common/utils"
)

func buildOptimismTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`optimism.json`)
	originalTokenList := helpers.FetchJSON[models.TokenListData[models.TokenListToken]](`https://raw.githubusercontent.com/ethereum-optimism/ethereum-optimism.github.io/master/optimism.tokenlist.json`)
	tokenList.Name = helpers.SafeString(originalTokenList.Name, `Optimism Token List`)
	tokenList.LogoURI = helpers.SafeString(originalTokenList.LogoURI, `https://ethereum-optimism.github.io/optimism.svg`)
	tokenList.Keywords = originalTokenList.Keywords

	/**************************************************************************
	* Parse the original token list and create a new one with the same tokens
	* with actual onchain data
	**************************************************************************/
	var newTokenList []models.TokenListToken
	grouped := helpers.GroupByChainID(originalTokenList.Tokens)
	for chainID, tokensForChain := range grouped {
		if !chains.IsChainIDSupported(chainID) {
			continue
		}

		tokensInfo := helpers.RetrieveBasicInformations(chainID, tokensForChain)
		for _, existingToken := range tokensForChain {
			if token, ok := tokensInfo[utils.ToAddress(existingToken)]; ok {
				if newToken, err := helpers.SetToken(
					token.Address,
					token.Name,
					token.Symbol,
					``,
					chainID,
					int(token.Decimals),
				); err == nil {
					newTokenList = append(newTokenList, newToken)
				}
			}
		}
	}

	/**************************************************************************
	* Ensure the data availability for the new token list is correct and
	* save it in a json file
	**************************************************************************/
	for _, token := range newTokenList {
		if (token.Name == `` || token.Symbol == `` || token.Decimals == 0) || chains.IsTokenIgnored(token.ChainID, token.Address) {
			continue
		}

		key := helpers.GetKey(token.ChainID, token.Address)
		tokenList.NextTokensMap[key] = token
	}

	tokens := helpers.GetTokensFromList(tokenList.Tokens)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `optimism.json`, helpers.SavingMethodStandard)
}
