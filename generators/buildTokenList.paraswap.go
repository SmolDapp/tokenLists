package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
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
	56:    `https://apiv5.paraswap.io/tokens/56`,
	137:   `https://apiv5.paraswap.io/tokens/137`,
	43114: `https://apiv5.paraswap.io/tokens/43114`,
}

func fetchParaswapTokenList() []TokenListToken {
	tokens := []TokenListToken{}

	for chainID, uri := range APIURIForParaswap {
		if helpers.IsChainIDIgnored(chainID) {
			continue
		}

		tokenAddresses := []common.Address{}
		list := helpers.FetchJSON[TParaswapList](uri)
		for _, v := range list.Tokens {
			tokenAddresses = append(tokenAddresses, common.HexToAddress(v.Address))
		}
		tokensInfo := ethereum.FetchNames(chainID, tokenAddresses)

		for _, token := range list.Tokens {
			logoURI := token.LogoURI
			if logoURI == `https://cdn.paraswap.io/token/token.png` {
				logoURI = ``
			}

			if newToken, err := SetToken(
				common.HexToAddress(token.Address),
				helpers.SafeString(tokensInfo[common.HexToAddress(token.Address).Hex()], ``),
				token.Symbol,
				helpers.SafeString(token.LogoURI, ``),
				chainID,
				token.Decimals,
			); err == nil {
				tokens = append(tokens, newToken)
			}
		}
	}

	return tokens
}

func buildParaswapTokenList() {
	tokenList := loadTokenListFromJsonFile(`paraswap.json`)
	tokenList.Name = "Paraswap Token List"
	tokenList.LogoURI = "https://app.paraswap.io/psp_logo.svg"

	tokens := fetchParaswapTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `paraswap.json`, Standard)
}
