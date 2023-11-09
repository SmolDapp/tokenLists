package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

type TWidoTokenData struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	LogoURI  string `json:"logoURI"`
	ChainID  uint64 `json:"chainId"`
	Decimals int    `json:"decimals"`
}
type TWidoList struct {
	Tokens []TWidoTokenData
}

func fetchWidoTokenList() []TokenListToken {
	tokens := []TokenListToken{}
	list := helpers.FetchJSON[TWidoList](`https://api.joinwido.com/tokens`)

	for _, token := range list.Tokens {
		if newToken, err := SetToken(
			common.HexToAddress(token.Address),
			token.Name,
			token.Symbol,
			token.LogoURI,
			token.ChainID,
			int(token.Decimals),
		); err == nil {
			tokens = append(tokens, newToken)
		}
	}
	return fetchTokenList(tokens)
}

func buildWidoTokenList() {
	tokenList := loadTokenListFromJsonFile(`wido.json`)
	tokenList.Name = "Wido Token List"
	tokenList.LogoURI = "https://uploads-ssl.webflow.com/61d5b57a99635946134f66aa/628b5dbbaed06b1b5b4d520a_logo-small.png"

	tokens := fetchWidoTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `wido.json`, Standard)
}
