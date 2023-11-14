package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

type T1InchTokenData struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	LogoURI  string `json:"logoURI"`
	Decimals int    `json:"decimals"`
}
type T1InchList struct {
	Tokens map[string]T1InchTokenData
}

var APIURIFor1Inch = map[uint64]string{
	1:     `https://api.1inch.io/v5.0/1/tokens`,
	10:    `https://api.1inch.io/v5.0/10/tokens`,
	56:    `https://api.1inch.io/v5.0/56/tokens`,
	100:   `https://api.1inch.io/v5.0/100/tokens`,
	137:   `https://api.1inch.io/v5.0/137/tokens`,
	250:   `https://api.1inch.io/v5.0/250/tokens`,
	42161: `https://api.1inch.io/v5.0/42161/tokens`,
	43114: `https://api.1inch.io/v5.0/43114/tokens`,
}

func fetch1InchTokenList() []TokenListToken {
	tokens := []TokenListToken{}

	for chainID, uri := range APIURIFor1Inch {
		if !helpers.IsChainIDSupported(chainID) {
			continue
		}

		list := helpers.FetchJSON[T1InchList](uri)
		allTokens := []common.Address{}
		for _, token := range list.Tokens {
			allTokens = append(allTokens, common.HexToAddress(token.Address))
		}
		tokensInfo := retrieveBasicInformations(chainID, allTokens)

		for _, existingToken := range list.Tokens {
			if token, ok := tokensInfo[existingToken.Address]; ok {
				if newToken, err := SetToken(
					token.Address,
					helpers.SafeString(token.Name, existingToken.Name),
					helpers.SafeString(token.Symbol, existingToken.Symbol),
					existingToken.LogoURI,
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

func build1InchTokenList() {
	tokenList := loadTokenListFromJsonFile(`1inch.json`)
	tokenList.Name = "1inch Token List"
	tokenList.LogoURI = "https://app.1inch.io/assets/images/logo.png"

	tokens := fetch1InchTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `1inch.json`, Standard)
}
