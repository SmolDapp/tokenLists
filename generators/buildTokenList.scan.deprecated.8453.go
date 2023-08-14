package main

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gocolly/colly"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

func handleScanTokenList_8453(tokenAddresses []common.Address, imageURI []string) []TokenListToken {
	tokenList := []TokenListToken{}

	tokensInfo := retrieveBasicInformations(8453, tokenAddresses)
	for index, address := range tokenAddresses {
		if token, ok := tokensInfo[address.Hex()]; ok {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				imageURI[index],
				8453,
				int(token.Decimals),
			); err == nil {
				tokenList = append(tokenList, newToken)
			}
		}
	}
	tokenList = addEtherToken(8453, tokenList)
	return tokenList
}

func fetchScanTokenList_8453(currentPage uint8) []TokenListToken {
	imageURI := []string{}
	tokens := []common.Address{}
	c := colly.NewCollector(
		colly.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36`),
	)

	c.OnHTML("div.media", func(e *colly.HTMLElement) {
		e.ForEach("img.u-xs-avatar", func(i int, h *colly.HTMLElement) {
			src := h.Attr("src")
			imageURI = append(imageURI, `https://basescan.org`+src)
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
		c.Visit(`https://basescan.org` + `/tokens?p=` + strconv.Itoa(int(currentPage)))
		currentPage++
	}
	return handleScanTokenList_8453(tokens, imageURI)
}

func buildScanTokenList_8453() {
	tokenList := loadTokenListFromJsonFile(`scan-8453.json`)
	tokenList.Name = `Base via BaseScan`
	tokenList.LogoURI = `https://etherscan.io/images/brandassets/etherscan-logo-circle.svg`
	tokenList.Keywords = []string{`base`, `etherscan`, `basescan`}
	tokens := fetchScanTokenList_8453(1)
	saveTokenListInJsonFile(tokenList, tokens, `scan-8453.json`, Standard)
}
