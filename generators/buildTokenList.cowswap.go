package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

func buildCowswapTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`cowswap.json`)
	originalTokenList := helpers.FetchJSON[models.TokenListData[models.TokenListToken]](
		`https://raw.githubusercontent.com/cowprotocol/token-lists/main/src/public/CowSwap.json`,
	)
	tokenList.Name = originalTokenList.Name
	tokenList.LogoURI = `https://raw.githubusercontent.com/cowprotocol/cowswap/c5974fb8a45d678029ecb013dab33722e152daaa/src/assets/cow-swap/cow_v2.svg`
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
		if _, ok := tokenList.NextTokensMap[key]; !ok {
			tokenList.NextTokensMap = make(map[string]models.TokenListToken)
		}
		tokenList.NextTokensMap[key] = token
	}

	tokens := helpers.GetTokensFromList(tokenList.Tokens)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `cowswap.json`, helpers.SavingMethodStandard)
}
