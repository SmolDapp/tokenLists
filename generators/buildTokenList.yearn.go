package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

func buildYearnTokenList() {
	tokenList := loadTokenListFromJsonFile(`yearn.json`)
	originalTokenList := helpers.FetchJSON[TokenListData](`https://ydaemon.yearn.finance/tokenlist`)
	tokenList.Name = `Yearn Token List`
	tokenList.LogoURI = `https://raw.githubusercontent.com/yearn/yearn-assets/master/icons/multichain-tokens/1/0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e/logo.svg`
	tokenList.Keywords = []string{`yearn`, `yfi`, `yvault`, `ytoken`, `ycurve`, `yprotocol`, `vaults`}
	tokenList = tokenList.Assign(originalTokenList.Tokens)

	saveTokenListInJsonFile(tokenList, tokenList.Tokens, `yearn.json`, Standard)
}
