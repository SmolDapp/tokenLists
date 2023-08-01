package main

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gocolly/colly"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

const ETHEREUM_CHAIN_ID = 1
const ETHEREUM_BASE_EXPLORER_URI = `https://etherscan.io`

func handleEthereumTokenList(tokenAddresses []common.Address, imageURI []string) []TokenListToken {
	tokenList := []TokenListToken{}

	tokensInfo := retrieveBasicInformations(ETHEREUM_CHAIN_ID, tokenAddresses)
	for index, address := range tokenAddresses {
		if token, ok := tokensInfo[address.Hex()]; ok {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				imageURI[index],
				ETHEREUM_CHAIN_ID,
				int(token.Decimals),
			); err == nil {
				tokenList = append(tokenList, newToken)
			}
		}
	}

	return tokenList
}

func fetchEthereumTokenList(currentPage uint8) []TokenListToken {
	imageURI := []string{}
	tokens := []common.Address{}
	c := colly.NewCollector(
		colly.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36`),
	)

	c.OnHTML("a.d-flex.align-items-center.gap-1.link-dark", func(e *colly.HTMLElement) {
		e.ForEach("img.rounded-circle", func(i int, h *colly.HTMLElement) {
			src := h.Attr("src")
			imageURI = append(imageURI, ETHEREUM_BASE_EXPLORER_URI+src)
		})
		tokenHref := e.Attr("href")
		tokenAddress := tokenHref[7:]
		tokens = append(tokens, common.HexToAddress(tokenAddress))
	})
	c.OnError(func(r *colly.Response, e error) {
		logs.Error(e)
	})

	for currentPage < 10 {
		c.Visit(ETHEREUM_BASE_EXPLORER_URI + `/tokens?p=` + strconv.Itoa(int(currentPage)))
		currentPage++
	}
	return handleEthereumTokenList(tokens, imageURI)
}

func buildEthereumTokenList() {
	tokenList := loadTokenListFromJsonFile(`ethereum-etherscan.json`)
	tokenList.Name = `Ethereum via Etherscan`
	tokenList.LogoURI = `https://2163491710-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-McrExXKKJBLJqymbFhO%2Fuploads%2Fgit-blob-29966808954100660b18274b582c1afb2fe45fed%2Fetherscan-logo-circle.png?alt=media`
	tokenList.Keywords = []string{`ethereum`, `etherscan`}
	tokens := fetchEthereumTokenList(1)
	saveTokenListInJsonFile(tokenList, tokens, `ethereum-etherscan.json`, Standard)
}
