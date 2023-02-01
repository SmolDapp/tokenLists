package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

type T1InchTokenData struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
	Address  string `json:"address"`
	LogoURI  string `json:"logoURI"`
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
			logs.Error(`impossible to fetch 1Inch token list`)
			return []TokenListToken{}
		}

		var list T1InchList
		if err := json.Unmarshal(body, &list); err != nil {
			logs.Error(err)
			return []TokenListToken{}
		}

		for _, v := range list.Tokens {
			if (v.Name == `` || v.Symbol == `` || v.Decimals == 0) || helpers.IsIgnoredToken(chainID, common.HexToAddress(v.Address)) {
				continue
			}

			tokens = append(tokens, TokenListToken{
				ChainID:  int(chainID),
				Decimals: v.Decimals,
				Address:  common.HexToAddress(v.Address).Hex(),
				Name:     v.Name,
				Symbol:   v.Symbol,
				LogoURI:  v.LogoURI,
			})
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
