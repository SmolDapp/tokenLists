package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

func buildCowswapTokenList() {
	tokenList := loadTokenListFromJsonFile(`cowswap.json`)
	originalTokenList := helpers.FetchJSON[TokenListData](`https://raw.githubusercontent.com/cowprotocol/token-lists/main/src/public/CowSwap.json`)
	tokenList.Name = originalTokenList.Name
	tokenList.LogoURI = `https://raw.githubusercontent.com/cowprotocol/cowswap/c5974fb8a45d678029ecb013dab33722e152daaa/src/assets/cow-swap/cow_v2.svg`
	tokenList.Keywords = originalTokenList.Keywords
	tokenList = tokenList.Assign(originalTokenList.Tokens)

	saveTokenListInJsonFile(tokenList, tokenList.Tokens, `cowswap.json`, Standard)
}
