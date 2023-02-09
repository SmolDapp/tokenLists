package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

func handleLedgerTokenList(tokensPerChainID map[uint64][]common.Address) []TokenListToken {
	tokenList := []TokenListToken{}

	for chainID, list := range tokensPerChainID {
		tokensInfo := retrieveBasicInformations(chainID, list)
		for _, address := range list {
			if token, ok := tokensInfo[address.Hex()]; ok {
				if newToken, err := SetToken(
					token.Address,
					token.Name,
					token.Symbol,
					``,
					chainID,
					int(token.Decimals),
				); err == nil {
					tokenList = append(tokenList, newToken)
				}
			}
		}
	}

	return tokenList
}

func fetchLedgerTokenList() []TokenListToken {
	tokensPerChainID := map[uint64][]common.Address{}
	tokensPerChainID[1] = []common.Address{}
	tokensPerChainID[56] = []common.Address{}
	tokensPerChainID[137] = []common.Address{}

	resp, err := http.Get(`https://raw.githubusercontent.com/LedgerHQ/ledger-live/develop/apps/ledger-live-desktop/cryptoassets.md`)
	if err != nil {
		return []TokenListToken{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []TokenListToken{}
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
			if !helpers.IsIgnoredToken(1, common.HexToAddress(address)) {
				tokensPerChainID[1] = append(tokensPerChainID[1], common.HexToAddress(address))
			}
		}
		if strings.HasPrefix(line, "| Binance Smart Chain |") {
			address := strings.Split(line, "|")[3]
			address = strings.TrimSpace(address)
			if !helpers.IsIgnoredToken(56, common.HexToAddress(address)) {
				tokensPerChainID[56] = append(tokensPerChainID[56], common.HexToAddress(address))
			}
		}
		if strings.HasPrefix(line, "| Polygon |") {
			address := strings.Split(line, "|")[3]
			address = strings.TrimSpace(address)
			if !helpers.IsIgnoredToken(137, common.HexToAddress(address)) {
				tokensPerChainID[137] = append(tokensPerChainID[137], common.HexToAddress(address))
			}
		}
	}
	return handleLedgerTokenList(tokensPerChainID)
}

func buildLedgersTokenList() {
	tokenList := loadTokenListFromJsonFile(`ledger.json`)
	tokenList.Name = "Ledger"
	tokenList.Keywords = []string{"Ledger"}
	tokenList.LogoURI = "https://www.ledger.com/wp-content/uploads/2021/11/Ledger_favicon.png"

	tokens := fetchLedgerTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `ledger.json`, Standard)
}
