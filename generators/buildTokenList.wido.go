package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

type TWidoTokenData struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	LogoURI  string `json:"logoURI"`
	Decimals int    `json:"decimals"`
	ChainID  int    `json:"chainId"`
}
type TWidoList struct {
	Tokens []TWidoTokenData
}

func fetchWidoTokenList() []TokenListToken {
	tokens := []TokenListToken{}
	resp, err := http.Get(`https://api.joinwido.com/tokens`)
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
		logs.Error(`impossible to fetch Wido token list`)
		return []TokenListToken{}
	}

	var list TWidoList
	if err := json.Unmarshal(body, &list); err != nil {
		logs.Error(err)
		return []TokenListToken{}
	}

	for _, v := range list.Tokens {
		if helpers.IsIgnoredChain(uint64(v.ChainID)) {
			continue
		}
		if (v.Name == `` || v.Symbol == `` || v.Decimals == 0) || helpers.IsIgnoredToken(uint64(v.ChainID), common.HexToAddress(v.Address)) {
			continue
		}
		tokens = append(tokens, TokenListToken{
			ChainID:  v.ChainID,
			Decimals: v.Decimals,
			Address:  common.HexToAddress(v.Address).Hex(),
			Name:     v.Name,
			Symbol:   v.Symbol,
			LogoURI:  v.LogoURI,
		})
	}
	return tokens
}

func buildWidoTokenList() {
	tokenList := loadTokenListFromJsonFile(`wido.json`)
	tokenList.Name = "Wido Token List"
	tokenList.LogoURI = "https://uploads-ssl.webflow.com/61d5b57a99635946134f66aa/628b5dbbaed06b1b5b4d520a_logo-small.png"

	tokens := fetchWidoTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `wido.json`, Standard)
}
