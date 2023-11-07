package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

func buildYearnExtendedTokenList() {
	tokenList := loadTokenListFromJsonFile(`yearn-extended.json`)
	originalTokenList := helpers.FetchJSON[TokenListData[TokenListToken]](`https://ydaemon.yearn.fi/tokenlist`)
	tokenList.Name = `Yearn Extended Token List`
	tokenList.LogoURI = `https://assets.smold.app/api/token/1/0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e/logo.svg`
	tokenList.Keywords = []string{`yearn`, `yfi`, `yvault`, `ytoken`, `ycurve`, `yprotocol`, `vaults`}
	tokenList = tokenList.Assign(originalTokenList.Tokens)

	saveTokenListInJsonFile(tokenList, tokenList.Tokens, `yearn-extended.json`, Standard)
}
