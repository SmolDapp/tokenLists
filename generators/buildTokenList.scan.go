package main

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gocolly/colly"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
	"github.com/migratooor/tokenLists/generators/common/models"
)

// L1 and L2 use a different code
type TExplorerType string

const (
	L1 TExplorerType = "L1"
	L2 TExplorerType = "L2"
)

type etherscanSASExplorers struct {
	BaseURL string
	Type    TExplorerType
}

var BASE_EXPLORERS_URI = map[uint64]etherscanSASExplorers{
	1: {
		BaseURL: "https://etherscan.io",
		Type:    L1,
	},
	10: {
		BaseURL: "https://optimistic.etherscan.io",
		Type:    L2,
	},
	56: {
		BaseURL: "https://bscscan.com",
		Type:    L2,
	},
	100: {
		BaseURL: "https://gnosisscan.io",
		Type:    L2,
	},
	137: {
		BaseURL: "https://polygonscan.com",
		Type:    L2,
	},
	1101: {
		BaseURL: "https://zkevm.polygonscan.com",
		Type:    L2,
	},
	250: {
		BaseURL: "https://ftmscan.com",
		Type:    L2,
	},
	8453: {
		BaseURL: "https://basescan.org",
		Type:    L2,
	},
	42161: {
		BaseURL: "https://arbiscan.io",
		Type:    L2,
	},
	81457: {
		BaseURL: "https://blastscan.io",
		Type:    L2,
	},
}

func handleScanTokenList(chainID uint64, tokenAddresses []common.Address, imageURI []string) []models.TokenListToken {
	tokenList := helpers.GetTokensFromAddresses(chainID, tokenAddresses)
	tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
	return tokenList
}

func fetchScanTokenListForL2(chainID uint64, currentPage uint8) []models.TokenListToken {
	explorerBaseUri := BASE_EXPLORERS_URI[chainID].BaseURL
	imageURI := []string{}
	tokens := []common.Address{}
	c := colly.NewCollector(
		colly.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36`),
	)

	c.OnHTML("div.media", func(e *colly.HTMLElement) {
		e.ForEach("img.u-xs-avatar", func(i int, h *colly.HTMLElement) {
			src := h.Attr("src")
			imageURI = append(imageURI, explorerBaseUri+src)
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
		c.Visit(explorerBaseUri + `/tokens?p=` + strconv.Itoa(int(currentPage)))
		currentPage++
	}
	return handleScanTokenList(chainID, tokens, imageURI)
}

func fetchScanTokenListForL1(chainID uint64, currentPage uint8) []models.TokenListToken {
	explorerBaseUri := BASE_EXPLORERS_URI[chainID].BaseURL
	imageURI := []string{}
	tokens := []common.Address{}
	c := colly.NewCollector(
		colly.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36`),
	)

	c.OnHTML("a.d-flex.align-items-center.gap-1.link-dark", func(e *colly.HTMLElement) {
		e.ForEach("img.rounded-circle", func(i int, h *colly.HTMLElement) {
			src := h.Attr("src")
			imageURI = append(imageURI, explorerBaseUri+src)
		})
		tokenHref := e.Attr("href")
		tokenAddress := tokenHref[7:]
		tokens = append(tokens, common.HexToAddress(tokenAddress))
	})
	c.OnError(func(r *colly.Response, e error) {
		logs.Error(e)
	})

	for currentPage < 20 {
		c.Visit(explorerBaseUri + `/tokens?p=` + strconv.Itoa(int(currentPage)))
		currentPage++
	}
	return handleScanTokenList(chainID, tokens, imageURI)
}

func fetchScanTokenList(chainID uint64) []models.TokenListToken {
	explorerBaseType := BASE_EXPLORERS_URI[chainID].Type
	if explorerBaseType == L1 {
		return fetchScanTokenListForL1(chainID, 1)
	}
	return fetchScanTokenListForL2(chainID, 1)
}

func buildScanTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`etherscan.json`)
	tokenList.Name = `Etherscan`
	tokenList.LogoURI = `https://etherscan.io/images/brandassets/etherscan-logo-circle.svg`
	tokenList.Keywords = []string{`ethereum`, `etherscan`}
	tokens := []models.TokenListToken{}
	tokens = append(tokens, fetchScanTokenList(1)...)
	tokens = append(tokens, fetchScanTokenList(10)...)
	tokens = append(tokens, fetchScanTokenList(56)...)
	tokens = append(tokens, fetchScanTokenList(100)...)
	tokens = append(tokens, fetchScanTokenList(137)...)
	tokens = append(tokens, fetchScanTokenList(250)...)
	tokens = append(tokens, fetchScanTokenList(1101)...)
	tokens = append(tokens, fetchScanTokenList(8453)...)
	tokens = append(tokens, fetchScanTokenList(42161)...)
	tokens = append(tokens, fetchScanTokenList(81457)...)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `etherscan.json`, helpers.SavingMethodStandard)
}
