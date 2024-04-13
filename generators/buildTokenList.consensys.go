package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

func buildConsensysTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`consensys.json`)
	originalTokenList := helpers.FetchJSON[models.TokenListData[models.TokenListToken]](
		`https://raw.githubusercontent.com/Consensys/linea-token-list/main/json/linea-mainnet-token-shortlist.json`,
	)
	tokenList.Name = originalTokenList.Name
	tokenList.LogoURI = `https://avatars.githubusercontent.com/u/10818037?s=200&v=4`
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
			if token, ok := tokensInfo[existingToken.Hex()]; ok {
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
		if (token.Name == `` || token.Symbol == `` || token.Decimals == 0) || chains.IsTokenIgnored(token.ChainID, common.HexToAddress(token.Address)) {
			continue
		}

		key := helpers.GetKey(token.ChainID, common.HexToAddress(token.Address))
		tokenList.NextTokensMap[key] = token
	}

	tokens := helpers.GetTokensFromList(tokenList.Tokens)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `consensys.json`, helpers.SavingMethodStandard)
}
