package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

type TDefillamaList struct {
	Name      string            `json:"name"`
	Symbol    string            `json:"symbol"`
	LogoURI   string            `json:"logoURI"`
	Address   common.Address    `json:"address,omitempty"`
	Platforms map[string]string `json:"platforms"`
}

func fetchDefillamaTokenList() []models.TokenListToken {
	list := helpers.FetchJSON[[]TDefillamaList](`https://defillama-datasets.llama.fi/tokenlist/all.json`)
	listPerChainID := []models.TokenListToken{}
	for _, v := range list {
		if len(v.Platforms) == 0 {
			continue
		}
		for platformName, addressOnPlatform := range v.Platforms {
			chainID := coingeckoMapNetworkNameToChainID(platformName)
			if !chains.IsChainIDSupported(chainID) {
				continue
			}
			if chains.IsTokenIgnored(chainID, addressOnPlatform) {
				continue
			}
			listPerChainID = append(listPerChainID, models.TokenListToken{
				Address: common.HexToAddress(addressOnPlatform).Hex(),
				Name:    v.Name,
				Symbol:  v.Symbol,
				LogoURI: v.LogoURI,
				ChainID: chainID,
			})
		}
	}
	return helpers.GetTokensFromList(listPerChainID)
}

func buildDefillamaTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`defillama.json`)
	tokenList.Name = "DefiLlama"
	tokenList.LogoURI = "https://wiki.defillama.com/w/resources/assets/wiki.png?88de1"

	tokens := fetchDefillamaTokenList()
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `defillama.json`, helpers.SavingMethodStandard)
}
