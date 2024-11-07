package main

import (
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

type TSolanaFMToken struct {
	TokenHash string `json:"tokenHash"`
	Data      struct {
		Mint          string   `json:"mint"`
		TokenName     string   `json:"tokenName"`
		Symbol        string   `json:"symbol"`
		Decimals      int      `json:"decimals"`
		Description   string   `json:"description"`
		Logo          string   `json:"logo"`
		Tags          []string `json:"tags"`
		Verified      string   `json:"verified"`
		Network       []string `json:"network"`
		MetadataToken string   `json:"metadataToken"`
	} `json:"data"`
}
type TSolanaFMEndpoint struct {
	Result struct {
		Data       []TSolanaFMToken `json:"data"`
		Pagination struct {
			Next string `json:"next"`
		} `json:"pagination"`
	} `json:"result"`
}

var solanaFMTokenEndpoint = `https://api.solana.fm/v0/tokens`

func fetchSolanaFMTokenList() []models.TokenListToken {
	tokenLists := []models.TokenListToken{}

	next := ``
	round := 0
	tokenAddresses := []string{}
	for {
		fetchingURL := solanaFMTokenEndpoint
		if next != `` {
			fetchingURL += `?from=` + next
		}
		list := helpers.FetchJSON[TSolanaFMEndpoint](fetchingURL)

		for _, token := range list.Result.Data {
			if token.Data.Verified != `true` {
				continue
			}
			if !helpers.Includes(token.Data.Network, `mainnet`) {
				continue
			}
			tokenAddresses = append(tokenAddresses, token.TokenHash)
		}
		next = list.Result.Pagination.Next
		if next == `` {
			break
		}
		round++
	}
	tokenLists = helpers.GetTokensFromAddresses(chains.SOLANA.ID, tokenAddresses)
	return tokenLists
}

func buildSolanaFMTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`solanaFM.json`)
	tokenList.Name = "SolanaFM Token List"
	tokenList.LogoURI = "https://avatars.githubusercontent.com/u/53240107?s=200&v=4"

	tokens := fetchSolanaFMTokenList()
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `solanaFM.json`, helpers.SavingMethodStandard)
}
