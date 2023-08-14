package main

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gocolly/colly"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

func handleTokenList1101(tokenAddresses []common.Address, imageURI []string) []TokenListToken {
	tokenList := []TokenListToken{}

	tokensInfo := retrieveBasicInformations(1101, tokenAddresses)
	for index, address := range tokenAddresses {
		if token, ok := tokensInfo[address.Hex()]; ok {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				imageURI[index],
				1101,
				int(token.Decimals),
			); err == nil {
				tokenList = append(tokenList, newToken)
			}
		}
	}
	tokenList = addEtherToken(1101, tokenList)
	return tokenList
}

func fetchScanTokenList_1101(currentPage uint64) []TokenListToken {
	imageURI := []string{}
	tokens := []common.Address{}
	c := colly.NewCollector()

	c.OnHTML("div.media", func(e *colly.HTMLElement) {
		e.ForEach("img.u-xs-avatar", func(i int, h *colly.HTMLElement) {
			src := h.Attr("src")
			imageURI = append(imageURI, `https://zkevm.polygonscan.com/`+src)
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

	for currentPage < 20 {
		c.Visit(`https://zkevm.polygonscan.com/` + `tokens?p=` + strconv.Itoa(int(currentPage)))
		currentPage++
	}
	return handleTokenList1101(tokens, imageURI)
}

func buildScanTokenList_1101() {
	tokenList := loadTokenListFromJsonFile(`polygon-zkevm.json`)
	tokenList.Name = `Polygon ZKEVM`
	tokenList.LogoURI = `https://raw.githubusercontent.com/maticnetwork/zkevm-docs/main/static/img/polygon-zkevm-logo.svg`
	tokenList.Keywords = []string{`polygon`, `zkevm`}
	tokens := fetchScanTokenList_1101(1)
	saveTokenListInJsonFile(tokenList, tokens, `polygon-zkevm.json`, Standard)
}
