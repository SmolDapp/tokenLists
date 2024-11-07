package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

func handleLedgerTokenList(tokensPerChainID map[uint64][]string) []models.TokenListToken {
	tokenLists := []models.TokenListToken{}

	for chainID, list := range tokensPerChainID {
		tokenList := helpers.GetTokensFromAddresses(chainID, list)
		tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
		tokenLists = append(tokenLists, tokenList...)
	}

	return tokenLists
}

func fetchLedgerTokenList() []models.TokenListToken {
	tokensPerChainID := map[uint64][]string{}
	tokensPerChainID[1] = []string{}
	tokensPerChainID[56] = []string{}
	tokensPerChainID[137] = []string{}

	resp, err := http.Get(`https://raw.githubusercontent.com/LedgerHQ/ledger-live/develop/apps/ledger-live-desktop/cryptoassets.md`)
	if err != nil {
		return []models.TokenListToken{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []models.TokenListToken{}
	}

	lines := strings.Split(string(body), "\n")
	index := 0
	for i, line := range lines {
		if strings.Contains(line, "parent currency") {
			index = i
			break
		}
	}

	//Slice the lines from the parent currency line and get all the addresses
	lines = lines[index+1:]
	for _, line := range lines {
		if strings.HasPrefix(line, "| Ethereum |") {
			address := strings.Split(line, "|")[3]
			address = strings.TrimSpace(address)
			if !chains.IsTokenIgnored(1, address) {
				tokensPerChainID[1] = append(tokensPerChainID[1], address)
			}
		}
		if strings.HasPrefix(line, "| Binance Smart Chain |") {
			address := strings.Split(line, "|")[3]
			address = strings.TrimSpace(address)
			if !chains.IsTokenIgnored(56, address) {
				tokensPerChainID[56] = append(tokensPerChainID[56], address)
			}
		}
		if strings.HasPrefix(line, "| Polygon |") {
			address := strings.Split(line, "|")[3]
			address = strings.TrimSpace(address)
			if !chains.IsTokenIgnored(137, address) {
				tokensPerChainID[137] = append(tokensPerChainID[137], address)
			}
		}
	}
	return handleLedgerTokenList(tokensPerChainID)
}

func buildLedgersTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`ledger.json`)
	tokenList.Name = "Ledger"
	tokenList.Keywords = []string{"Ledger"}
	tokenList.LogoURI = "https://www.ledger.com/wp-content/uploads/2021/11/Ledger_favicon.png"

	tokens := fetchLedgerTokenList()
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `ledger.json`, helpers.SavingMethodStandard)
}
