package main

import (
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/common/utils"
)

type TParaswapTokenData struct {
	Symbol   string `json:"symbol"`
	Address  string `json:"address"`
	LogoURI  string `json:"img"`
	Decimals int    `json:"decimals"`
}
type TParaswapList struct {
	Tokens []TParaswapTokenData
}

var APIURIForParaswap = map[uint64]string{
	1:     `https://apiv5.paraswap.io/tokens/1`,
	10:    `https://apiv5.paraswap.io/tokens/10`,
	56:    `https://apiv5.paraswap.io/tokens/56`,
	137:   `https://apiv5.paraswap.io/tokens/137`,
	8453:  `https://apiv5.paraswap.io/tokens/8453`,
	42161: `https://apiv5.paraswap.io/tokens/42161`,
	43114: `https://apiv5.paraswap.io/tokens/43114`,
}

func fetchParaswapTokenList() []models.TokenListToken {
	tokens := []models.TokenListToken{}

	for chainID, uri := range APIURIForParaswap {
		if !chains.IsChainIDSupported(chainID) {
			continue
		}

		tokenAddresses := []string{}
		list := helpers.FetchJSON[TParaswapList](uri)
		for _, v := range list.Tokens {
			tokenAddresses = append(tokenAddresses, v.Address)
		}

		tokensInfo := helpers.RetrieveBasicInformations(chainID, tokenAddresses)
		for _, existingToken := range list.Tokens {
			if token, ok := tokensInfo[utils.ToAddress(existingToken.Address)]; ok {
				logoURI := existingToken.LogoURI
				if logoURI == `https://cdn.paraswap.io/token/token.png` {
					logoURI = ``
				}

				if newToken, err := helpers.SetToken(
					token.Address,
					token.Name,
					token.Symbol,
					helpers.SafeString(logoURI, ``),
					chainID,
					int(token.Decimals),
				); err == nil {
					tokens = append(tokens, newToken)
				}
			}
		}
	}

	return tokens
}

func buildParaswapTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`paraswap.json`)
	tokenList.Name = "Paraswap Token List"
	tokenList.LogoURI = "https://app.paraswap.io/psp_logo.svg"

	tokens := fetchParaswapTokenList()
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `paraswap.json`, helpers.SavingMethodStandard)
}
