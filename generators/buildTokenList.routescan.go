package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
	"github.com/migratooor/tokenLists/generators/common/models"
)

var ROUTESCAN_URI = map[uint64]string{
	81457: `https://cdn.routescan.io/api/evm/all/erc20`,
}

func handleRouteScanTokenList(chainID uint64, tokenAddresses []common.Address) []models.TokenListToken {
	tokenList := helpers.GetTokensFromAddresses(chainID, tokenAddresses)
	tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
	return tokenList
}

func fetchRouteScanTokenList(chainID uint64) []models.TokenListToken {
	type TRoutescanAPIResponse struct {
		Items []struct {
			Address string `json:"address"`
			Details struct {
				Icon      string `json:"icon"`
				Type      string `json:"type"`
				Blacklist bool   `json:"blacklist"`
			}
			Transfers struct {
				last24h uint64 `json:"last24h"`
			}
			Holders uint64 `json:"holders"`
		} `json:"items"`
	}

	explorerBaseURI := ROUTESCAN_URI[chainID]
	nextPageURI := `?count=false&includedChainIds=81457&limit=1000&sort=marketCap%2Cdesc`
	tokens := []common.Address{}
	response := helpers.FetchJSON[TRoutescanAPIResponse](explorerBaseURI + nextPageURI)
	for _, token := range response.Items {
		if token.Details.Type == `ERC-721` || token.Details.Type == `ERC-1155` {
			continue
		}
		if token.Details.Blacklist {
			continue
		}
		if token.Holders == 0 {
			continue
		}
		if token.Transfers.last24h == 0 {
			continue
		}
		tokens = append(tokens, common.HexToAddress(token.Address))
	}

	return handleRouteScanTokenList(chainID, tokens)
}

func buildRouteScanTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`routescan.json`)
	tokenList.Name = `RouteScan`
	tokenList.LogoURI = `https://cms-cdn.avascan.com/cms2/routescan.432df9c80dd7.svg`
	tokenList.Keywords = []string{`explorer`, `routescan`}

	tokens := []models.TokenListToken{}
	for chainID := range ROUTESCAN_URI {
		tokens = append(tokens, fetchRouteScanTokenList(chainID)...)
	}
	logs.Pretty(tokens)
	// helpers.SaveTokenListInJsonFile(tokenList, tokens, `routescan.json`, helpers.SavingMethodStandard)
}
