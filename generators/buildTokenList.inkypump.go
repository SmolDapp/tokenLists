package main

import (
	"strconv"

	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/common/utils"
)

func handleInkypumpTokenList(
	chainID uint64,
	tokenAddresses []string,
	tokenIcons map[string]string,
) []models.TokenListToken {
	tokenList := helpers.GetTokensFromAddressesWithIcons(chainID, tokenAddresses, tokenIcons)
	tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
	return tokenList
}

func fetchInkypumpTokenList() []models.TokenListToken {
	type TInkypumpAPIResponse struct {
		Tokens []struct {
			Address         string  `json:"address"`
			IconURI         string  `json:"image_url"`
			Name            string  `json:"name"`
			Symbol          string  `json:"ticker"`
			FundingProgress float64 `json:"funding_progress"`
		} `json:"tokens"`
		TotalPages int `json:"totalPages"`
		TotalCount int `json:"totalCount"`
	}

	baseAPIEndpoint := `https://inkypump.com/api/tokens`
	tokenAddresses := []string{}
	tokenIcons := make(map[string]string)
	for i := 0; i < 40; i++ {
		nextPageURI := `?page=` + strconv.Itoa(i) + `&sortBy=mcap-high`
		response := helpers.FetchJSON[TInkypumpAPIResponse](baseAPIEndpoint + nextPageURI)
		for _, token := range response.Tokens {
			if token.FundingProgress >= 1.0 {
				tokenAddresses = append(tokenAddresses, token.Address)
				tokenIcons[utils.ToAddress(token.Address)] = token.IconURI
			}
		}
		if i >= response.TotalPages {
			break
		}
	}
	return handleInkypumpTokenList(57073, tokenAddresses, tokenIcons)
}

func buildInkypumpTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`inkypump.json`)
	tokenList.Name = `Inkypump`
	tokenList.LogoURI = `https://assets.smold.app/api/chain/57073/logo-128.png`
	tokenList.Keywords = []string{`inkypump`, `launcher`}
	tokens := []models.TokenListToken{}
	tokens = append(tokens, fetchInkypumpTokenList()...)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `inkypump.json`, helpers.SavingMethodStandard)
}
