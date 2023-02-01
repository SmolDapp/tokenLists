package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

func fetchYearnTokenList() TokenListData {
	resp, err := http.Get(`https://ydaemon.yearn.finance/tokenlist`)
	if err != nil {
		logs.Error(err)
		return TokenListData{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return TokenListData{}
	}

	var list TokenListData
	if err := json.Unmarshal(body, &list); err != nil {
		logs.Error(err)
		return TokenListData{}
	}
	return list
}

func buildYearnTokenList() {
	tokenList := loadTokenListFromJsonFile(`yearn.json`)
	originalTokenList := fetchYearnTokenList()
	tokenList.Name = `Yearn Token List`
	tokenList.LogoURI = `https://raw.githubusercontent.com/yearn/yearn-assets/master/icons/multichain-tokens/1/0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e/logo.svg`
	tokenList.Keywords = []string{`yearn`, `yfi`, `yvault`, `ytoken`, `ycurve`, `yprotocol`, `vaults`}

	for _, token := range originalTokenList.Tokens {
		if helpers.IsIgnoredChain(uint64(token.ChainID)) {
			continue
		}
		if (token.Name == `` || token.Symbol == `` || token.Decimals == 0) || helpers.IsIgnoredToken(uint64(token.ChainID), common.HexToAddress(token.Address)) {
			continue
		}

		key := GetKey(uint64(token.ChainID), common.HexToAddress(token.Address))
		tokenList.NextTokensMap[key] = token
	}

	saveTokenListInJsonFile(tokenList, tokenList.Tokens, `yearn.json`, Standard)
}
