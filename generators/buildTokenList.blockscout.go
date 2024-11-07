package main

import (
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

var BLOCKSCOUTV5_URI = map[uint64]string{
	42220: `https://explorer.celo.org/mainnet/`,
}

var BLOCKSCOUTV6_URI = map[uint64]string{
	1:       `https://eth.blockscout.com`,
	5:       `https://eth-goerli.blockscout.com`,
	10:      `https://optimism.blockscout.com`,
	100:     `https://gnosis.blockscout.com`,
	137:     `https://polygon.blockscout.com`,
	314:     `https://filecoin.blockscout.com`, //No market cap
	1101:    `https://zkevm.blockscout.com`,
	1088:    `https://andromeda-explorer.metis.io`, //No market cap
	5000:    `https://explorer.mantle.xyz`,
	8453:    `https://base.blockscout.com`,
	7777777: `https://explorer.zora.energy`, //No market cap
}

var ALLOW_EMPTY_MARKET_CAP = map[uint64]bool{
	314:     false,
	1088:    true,
	7777777: true,
}

func handleBlockScoutTokenList(chainID uint64, tokenAddresses []string) []models.TokenListToken {
	tokenList := helpers.GetTokensFromAddresses(chainID, tokenAddresses)
	tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
	return tokenList
}

func fetchBlockScoutV5TokenList(chainID uint64) []models.TokenListToken {
	type TBlockScoutAPIResponse struct {
		Items    []string `json:"items"`
		NextPage string   `json:"next_page_path"`
	}

	explorerBaseURI := BLOCKSCOUTV5_URI[chainID]
	nextPageURI := `/tokens?type=JSON`
	tokens := []string{}

	for i := 0; i < 20; i++ {
		response := helpers.FetchJSON[TBlockScoutAPIResponse](explorerBaseURI + nextPageURI)
		for _, token := range response.Items {
			dataIdentifierHash := strings.Split(token, "data-identifier-hash=\"")[1]
			dataIdentifierHash = strings.Split(dataIdentifierHash, "\"")[0]
			tokens = append(tokens, common.HexToAddress(dataIdentifierHash).Hex())
		}
		nextPageURI = response.NextPage + `&type=JSON`
	}

	return handleBlockScoutTokenList(chainID, tokens)
}

func fetchBlockScoutV6TokenList(chainID uint64) []models.TokenListToken {
	type TBlockScoutAPIResponse struct {
		Items []struct {
			Address   string `json:"address"`
			IconURI   string `json:"icon_url"`
			Type      string `json:"type"`
			MarketCap string `json:"circulating_market_cap,omitempty"`
		} `json:"items"`
		NextPage struct {
			ContractAddressHash string `json:"contract_address_hash"`
			FiatValue           string `json:"fiat_value"`
			HolderCount         int    `json:"holder_count"`
			IsNameNull          bool   `json:"is_name_null"`
			ItemsCount          int    `json:"items_count"`
			MarketCap           string `json:"market_cap"`
			Name                string `json:"name"`
		} `json:"next_page_params"`
	}

	explorerBaseURI := BLOCKSCOUTV6_URI[chainID]
	nextPageURI := `/api/v2/tokens`
	tokens := []string{}

	for i := 0; i < 40; i++ {
		response := helpers.FetchJSON[TBlockScoutAPIResponse](explorerBaseURI + nextPageURI)
		for _, token := range response.Items {
			if token.Type == `ERC-721` || token.Type == `ERC-1155` {
				continue
			}
			if token.MarketCap == `0` || token.MarketCap == `` {
				if !ALLOW_EMPTY_MARKET_CAP[chainID] {
					continue
				}
			}
			tokens = append(tokens, common.HexToAddress(token.Address).Hex())
		}
		nextPageURI = `/api/v2/tokens?contract_address_hash=` + response.NextPage.ContractAddressHash + `&fiat_value=` + response.NextPage.FiatValue + `&holder_count=` + strconv.Itoa(response.NextPage.HolderCount) + `&is_name_null=` + strconv.FormatBool(response.NextPage.IsNameNull) + `&items_count=` + strconv.Itoa(response.NextPage.ItemsCount) + `&market_cap=` + response.NextPage.MarketCap + `&name=` + response.NextPage.Name
		nextPageURI = strings.ReplaceAll(nextPageURI, ` `, `%20`)
	}

	return handleBlockScoutTokenList(chainID, tokens)
}

func buildBlockScoutTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`blockscout.json`)
	tokenList.Name = `Blockscout`
	tokenList.LogoURI = `https://2383309224-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-Lq1XoWGmy8zggj_u2fM%2Ficon%2FyFkt6mPJJvjKiSBBOppe%2FBS_logo_slack.png?alt=media`
	tokenList.Keywords = []string{`explorer`, `blockscout`}

	tokens := []models.TokenListToken{}
	for chainID := range BLOCKSCOUTV5_URI {
		tokens = append(tokens, fetchBlockScoutV5TokenList(chainID)...)
	}
	for chainID := range BLOCKSCOUTV6_URI {
		tokens = append(tokens, fetchBlockScoutV6TokenList(chainID)...)
	}
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `blockscout.json`, helpers.SavingMethodStandard)
}
