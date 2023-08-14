package main

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gocolly/colly"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

func handleEthereumTokenList(tokenAddresses []common.Address, imageURI []string) []TokenListToken {
	tokenList := []TokenListToken{}

	tokensInfo := retrieveBasicInformations(1, tokenAddresses)
	for index, address := range tokenAddresses {
		if token, ok := tokensInfo[address.Hex()]; ok {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				imageURI[index],
				1,
				int(token.Decimals),
			); err == nil {
				tokenList = append(tokenList, newToken)
			}
		}
	}
	tokenList = addEtherToken(1, tokenList)
	return tokenList
}

func fetchScanTokenList_1(currentPage uint8) []TokenListToken {
	imageURI := []string{}
	tokens := []common.Address{}
	c := colly.NewCollector(
		colly.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36`),
	)

	c.OnHTML("a.d-flex.align-items-center.gap-1.link-dark", func(e *colly.HTMLElement) {
		e.ForEach("img.rounded-circle", func(i int, h *colly.HTMLElement) {
			src := h.Attr("src")
			imageURI = append(imageURI, `https://etherscan.io`+src)
		})
		tokenHref := e.Attr("href")
		tokenAddress := tokenHref[7:]
		tokens = append(tokens, common.HexToAddress(tokenAddress))
	})
	c.OnError(func(r *colly.Response, e error) {
		logs.Error(e)
	})

	for currentPage < 20 {
		c.Visit(`https://etherscan.io` + `/tokens?p=` + strconv.Itoa(int(currentPage)))
		currentPage++
	}
	return handleEthereumTokenList(tokens, imageURI)
}

func buildScanTokenList_1() {
	tokenList := loadTokenListFromJsonFile(`ethereum-etherscan.json`)
	tokenList.Name = `Ethereum via Etherscan`
	tokenList.LogoURI = `https://etherscan.io/images/brandassets/etherscan-logo-circle.svg`
	tokenList.Keywords = []string{`ethereum`, `etherscan`}
	tokens := fetchScanTokenList_1(1)
	saveTokenListInJsonFile(tokenList, tokens, `ethereum-etherscan.json`, Standard)
}
