package main

import (
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/common/utils"
)

func buildCowswapTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`cowswap.json`)
	originalTokenList := helpers.FetchJSON[models.TokenListData[models.TokenListToken]](
		`https://raw.githubusercontent.com/cowprotocol/token-lists/main/src/public/CowSwap.json`,
	)
	tokenList.Name = originalTokenList.Name
	tokenList.LogoURI = `https://raw.githubusercontent.com/cowprotocol/cowswap/c5974fb8a45d678029ecb013dab33722e152daaa/src/assets/cow-swap/cow_v2.svg`
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
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `cowswap.json`, helpers.SavingMethodStandard)
}
