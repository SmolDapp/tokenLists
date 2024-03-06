package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

var ROUTESCAN_URI = map[uint64]string{
	81457: `https://cdn.routescan.io/api/evm/all/erc20`,
}

func handleRouteScanTokenList(chainID uint64, tokenAddresses []common.Address, logos map[common.Address]string) []models.TokenListToken {
	tokenList := helpers.GetTokensFromAddresses(chainID, tokenAddresses)
	tokenList = append(tokenList, chains.CHAINS[chainID].Coin)

	for i, token := range tokenList {
		if logo, ok := logos[common.HexToAddress(token.Address)]; ok {
			if logo != `` {
				tokenList[i].LogoURI = logo
			}
		}
	}

	return tokenList
}

func fetchRouteScanTokenList(chainID uint64) []models.TokenListToken {
	type TRoutescanAPIResponse struct {
		Items []struct {
			Address string `json:"address"`
			Detail  struct {
				Icon      string `json:"icon"`
				Type      string `json:"type"`
				Blacklist bool   `json:"blacklist"`
			}
			Transfers struct {
				Last24h uint64 `json:"last24h"`
			}
			Holders uint64 `json:"holders"`
		} `json:"items"`
	}

	explorerBaseURI := ROUTESCAN_URI[chainID]
	nextPageURI := `?count=false&includedChainIds=81457&limit=1000&sort=marketCap%2Cdesc`
	tokens := []common.Address{}
	logos := map[common.Address]string{}
	response := helpers.FetchJSON[TRoutescanAPIResponse](explorerBaseURI + nextPageURI)
	for _, token := range response.Items {
		if token.Detail.Type == `ERC-721` || token.Detail.Type == `ERC-1155` {
			continue
		}
		if token.Detail.Blacklist {
			continue
		}
		if token.Holders == 0 {
			continue
		}
		if token.Transfers.Last24h == 0 {
			continue
		}
		tokens = append(tokens, common.HexToAddress(token.Address))
		logos[common.HexToAddress(token.Address)] = token.Detail.Icon
	}

	return handleRouteScanTokenList(chainID, tokens, logos)
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
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `routescan.json`, helpers.SavingMethodStandard)
}
