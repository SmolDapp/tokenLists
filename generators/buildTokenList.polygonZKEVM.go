package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gocolly/colly"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

const POLYGON_ZKEVM_CHAIN_ID = 1101
const POLYGON_ZKEVM_BASE_EXPLORER_URI = `https://zkevm.polygonscan.com/`

func handlePolygonZKEVMTokenList(tokenAddresses []common.Address, imageURI []string) []TokenListToken {
	tokenList := []TokenListToken{}

	tokensInfo := retrieveBasicInformations(POLYGON_ZKEVM_CHAIN_ID, tokenAddresses)
	for index, address := range tokenAddresses {
		if token, ok := tokensInfo[address.Hex()]; ok {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				imageURI[index],
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
	imageURI := []string{}
	tokens := []common.Address{}
	c := colly.NewCollector()

	c.OnHTML("div.media", func(e *colly.HTMLElement) {
		e.ForEach("img.u-xs-avatar", func(i int, h *colly.HTMLElement) {
			src := h.Attr("src")
			imageURI = append(imageURI, POLYGON_ZKEVM_BASE_EXPLORER_URI+src)
		})
		e.ForEach("a.text-primary", func(i int, h *colly.HTMLElement) {
			tokenHref := h.Attr("href")
			tokenAddress := tokenHref[7:]
			tokens = append(tokens, common.HexToAddress(tokenAddress))
		})

	})
	c.OnError(func(r *colly.Response, e error) {
		logs.Error(e)
	})

	c.Visit(POLYGON_ZKEVM_BASE_EXPLORER_URI + `tokens`)
	return handlePolygonZKEVMTokenList(tokens, imageURI)
}

func buildPolygonZKEVMTokenList() {
	tokenList := loadTokenListFromJsonFile(`polygon-zkevm.json`)
	tokenList.Name = `Polygon ZKEVM`
	tokenList.LogoURI = `https://raw.githubusercontent.com/maticnetwork/zkevm-docs/main/static/img/polygon-zkevm-logo.svg`
	tokenList.Keywords = []string{`polygon`, `zkevm`}
	tokens := fetchPolygonZKEVMTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `polygon-zkevm.json`, Standard)
}
