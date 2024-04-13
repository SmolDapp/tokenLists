package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

func handleZkSyncTokenList(chainID uint64, tokenAddresses []common.Address) []models.TokenListToken {
	tokenList := helpers.GetTokensFromAddresses(chainID, tokenAddresses)
	tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
	return tokenList
}

func fetchZkSyncTokenList() []models.TokenListToken {
	type TZkSyncAPIResponse struct {
		Items []struct {
			Address string `json:"l2address"`
			IconURI string `json:"iconURL"`
		} `json:"items"`
		Meta struct {
			TotalItems int `json:"totalItems"`
		}
		Links struct {
			Next string `json:"next"`
		} `json:"links"`
	}

	baseAPIEndpoint := `https://block-explorer-api.mainnet.zksync.io/`
	nextPageURI := `tokens?page=1&limit=100&minLiquidity=0`
	tokens := []common.Address{}
	for i := 0; i < 40; i++ {
		response := helpers.FetchJSON[TZkSyncAPIResponse](baseAPIEndpoint + nextPageURI)
		for _, token := range response.Items {
			tokens = append(tokens, common.HexToAddress(token.Address))
		}
		nextPageURI = response.Links.Next
		if nextPageURI == `` {
			break
		}
		if len(tokens) >= response.Meta.TotalItems {
			break
		}
	}
	return handleZkSyncTokenList(324, tokens)
}

func buildZkSyncTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`zksync.json`)
	tokenList.Name = `zkSync`
	tokenList.LogoURI = `https://assets.smold.app/api/chain/324/logo-128.png`
	tokenList.Keywords = []string{`zksync`, `explorer`}
	tokens := []models.TokenListToken{}
	tokens = append(tokens, fetchZkSyncTokenList()...)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `zksync.json`, helpers.SavingMethodStandard)
}
