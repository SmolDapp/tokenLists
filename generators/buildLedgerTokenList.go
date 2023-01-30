package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

var TOKENS_BY_LEDGER = map[uint64][]common.Address{
	1:   {},
	56:  {},
	137: {},
}

func fetchLedgerList() {
	resp, err := http.Get(`https://raw.githubusercontent.com/LedgerHQ/ledger-live/develop/apps/ledger-live-desktop/cryptoassets.md`)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
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
			TOKENS_BY_LEDGER[1] = append(TOKENS_BY_LEDGER[1], common.HexToAddress(address))
		}
		if strings.HasPrefix(line, "| Binance Smart Chain |") {
			address := strings.Split(line, "|")[3]
			TOKENS_BY_LEDGER[56] = append(TOKENS_BY_LEDGER[56], common.HexToAddress(address))
		}
		if strings.HasPrefix(line, "| Polygon |") {
			address := strings.Split(line, "|")[3]
			TOKENS_BY_LEDGER[56] = append(TOKENS_BY_LEDGER[56], common.HexToAddress(address))
		}
	}
}

func buildLedgersTokenList() {
	fetchLedgerList()

	var tokenList TokenListData
	tokenList.Name = "Ledger"
	tokenList.Keywords = []string{"ledger", "BNB", "ETH", "POLYGON"}
	tokenList.LogoURI = "https://www.ledger.com/wp-content/uploads/2021/11/Ledger_favicon.png"
	tokenList.Timestamp = time.Now().Format(time.RFC3339)
	tokenList.Version.Major = 1
	tokenList.Version.Minor = 0
	tokenList.Version.Patch = 0

	for chainID, list := range TOKENS_BY_LEDGER {
		tokensInfo := fetchBasicInformations(chainID, list)
		for _, address := range list {
			if token, ok := tokensInfo[address.Hex()]; ok {
				tokenList.Tokens = append(tokenList.Tokens, TokenListToken{
					ChainID:  int(chainID),
					Address:  token.Address.Hex(),
					Name:     token.Name,
					Symbol:   token.Symbol,
					Decimals: int(token.Decimals),
					LogoURI:  ``,
				})
			}
		}
	}

	jsonData, err := json.MarshalIndent(tokenList, "", "  ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile(helpers.BASE_PATH+`/lists/ledger.json`, jsonData, 0644)
	if err != nil {
		logs.Error(err)
		return
	}
}
