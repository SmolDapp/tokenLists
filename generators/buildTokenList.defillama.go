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

func defillamaMapNetworkNameToChainID(network string) uint64 {
	switch network {
	case `ethereum`:
		return 1
	case `optimistic-ethereum`:
		return 10
	case `binance-smart-chain`:
		return 56
	case `xdai`:
		return 100
	case `polygon-pos`:
		return 137
	case `fantom`:
		return 250
	case `arbitrum-one`:
		return 42161
	case `avalanche`:
		return 43114
	}
	return 0
}

func fetchDefillamaTokenList() []models.TokenListToken {
	list := helpers.FetchJSON[[]TDefillamaList](`https://defillama-datasets.llama.fi/tokenlist/all.json`)
	listPerChainID := []models.TokenListToken{}
	for _, v := range list {
		if len(v.Platforms) == 0 {
			continue
		}
		for platformName, addressOnPlatform := range v.Platforms {
			chainID := defillamaMapNetworkNameToChainID(platformName)
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
