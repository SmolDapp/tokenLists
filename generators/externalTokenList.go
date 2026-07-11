package main

import (
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/common/utils"
)

func buildExternalTokenList(tokenList models.TokenListData[models.TokenListToken], originalTokens []models.TokenListToken, fileName string) {
	/**************************************************************************
	* Parse the original token list and create a new one with the same tokens
	* with actual onchain data
	**************************************************************************/
	var newTokenList []models.TokenListToken
	grouped := helpers.GroupByChainID(originalTokens)
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
	helpers.SaveTokenListInJsonFile(tokenList, newTokenList, fileName, helpers.SavingMethodStandard)
}
