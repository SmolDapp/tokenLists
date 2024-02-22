package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

type TUniContracts struct {
	ContractAddress common.Address
	BlockNumber     *big.Int
	Type            int
}

var UniswapContractsPerChainID = map[uint64][]TUniContracts{
	1: {
		{
			ContractAddress: common.HexToAddress(`0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f`),
			BlockNumber:     big.NewInt(10_000_835),
			Type:            2,
		},
		{
			ContractAddress: common.HexToAddress(`0x1F98431c8aD98523631AE4a59f267346ea31F984`),
			BlockNumber:     big.NewInt(12_369_621),
			Type:            3,
		},
	},
	10: {
		{
			ContractAddress: common.HexToAddress(`0x1F98431c8aD98523631AE4a59f267346ea31F984`),
			BlockNumber:     big.NewInt(100_000),
			Type:            3,
		},
	},
	137: {
		{
			ContractAddress: common.HexToAddress(`0x1F98431c8aD98523631AE4a59f267346ea31F984`),
			BlockNumber:     big.NewInt(22_757_547),
			Type:            3,
		},
	},
	42161: {
		{
			ContractAddress: common.HexToAddress(`0x1F98431c8aD98523631AE4a59f267346ea31F984`),
			BlockNumber:     big.NewInt(165),
			Type:            3,
		},
	},
}

func buildUniswapTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`uniswap.json`)
	originalTokenList := helpers.FetchJSON[models.TokenListData[models.TokenListToken]](`https://tokens.uniswap.org`)
	tokenList.Name = helpers.SafeString(originalTokenList.Name, `Uniswap Token List`)
	tokenList.LogoURI = helpers.SafeString(originalTokenList.LogoURI, `ipfs://QmNa8mQkrNKp1WEEeGjFezDmDeodkWRevGFN8JCV7b4Xir"`)
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
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `uniswap.json`, helpers.SavingMethodStandard)
}
