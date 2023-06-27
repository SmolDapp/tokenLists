package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gocolly/colly"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

const POLYGON_ZKEVM_CHAIN_ID = 1101

func handlePolygonZKEVMTokenList(tokenAddresses []common.Address) []TokenListToken {
	tokenList := []TokenListToken{}

	tokensInfo := retrieveBasicInformations(POLYGON_ZKEVM_CHAIN_ID, tokenAddresses)
	for _, address := range tokenAddresses {
		if token, ok := tokensInfo[address.Hex()]; ok {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				``,
				POLYGON_ZKEVM_CHAIN_ID,
				int(token.Decimals),
			); err == nil {
				tokenList = append(tokenList, newToken)
			}
		}
	}

	return tokenList
}

func fetchPolygonZKEVMTokenList() []TokenListToken {
	tokens := []common.Address{}
	c := colly.NewCollector()

	c.OnHTML("div.media-body", func(e *colly.HTMLElement) {
		e.ForEach("a.text-primary", func(i int, h *colly.HTMLElement) {
			tokenHref := h.Attr("href")
			tokenAddress := tokenHref[7:]
			tokens = append(tokens, common.HexToAddress(tokenAddress))
		})

	})
	c.OnError(func(r *colly.Response, e error) {
		logs.Error(e)
	})

	c.Visit(`https://zkevm.polygonscan.com/tokens`)
	return handlePolygonZKEVMTokenList(tokens)
}

func buildPolygonZKEVMTokenList() {
	tokenList := loadTokenListFromJsonFile(`polygonZKEVM.json`)
	tokenList.Name = `Polygon ZKEVM`
	tokenList.LogoURI = `https://raw.githubusercontent.com/maticnetwork/zkevm-docs/main/static/img/polygon-zkevm-logo.svg`
	tokenList.Keywords = []string{`polygon`, `zkevm`}
	tokens := fetchPolygonZKEVMTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `polygonZKEVM.json`, Standard)
}
