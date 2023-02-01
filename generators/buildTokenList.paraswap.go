package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
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
		if helpers.IsIgnoredChain(chainID) {
			continue
		}

		resp, err := http.Get(uri)
		if err != nil {
			logs.Error(err)
			return []TokenListToken{}
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logs.Error(err)
			return []TokenListToken{}
		}
		if (resp.StatusCode < 200) || (resp.StatusCode > 299) {
			logs.Error(`impossible to fetch Paraswap token list`)
			return []TokenListToken{}
		}

		var list TParaswapList
		if err := json.Unmarshal(body, &list); err != nil {
			logs.Error(err)
			return []TokenListToken{}
		}

		tokenAddresses := []common.Address{}
		for _, v := range list.Tokens {
			tokenAddresses = append(tokenAddresses, common.HexToAddress(v.Address))
		}
		tokensInfo := fetchNames(chainID, tokenAddresses)

		for _, v := range list.Tokens {
			name, ok := tokensInfo[common.HexToAddress(v.Address).Hex()]
			if !ok {
				continue
			}
			if (name == `` || v.Symbol == `` || v.Decimals == 0) || helpers.IsIgnoredToken(chainID, common.HexToAddress(v.Address)) {
				continue
			}

			logoURI := v.LogoURI
			if logoURI == `https://cdn.paraswap.io/token/token.png` {
				logoURI = ``
			}
			tokens = append(tokens, TokenListToken{
				ChainID:  int(chainID),
				Decimals: v.Decimals,
				Address:  common.HexToAddress(v.Address).Hex(),
				Name:     name,
				Symbol:   v.Symbol,
				LogoURI:  logoURI,
			})
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
