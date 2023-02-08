package main

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
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

func handleDefillamaTokenList(listPerChainID map[uint64][]TDefillamaList) []TokenListToken {
	tokensForChainIDSyncMap := initSyncMap(listPerChainID)

	// Fetch the basic informations for all the tokens for all the chains
	perChainWG := sync.WaitGroup{}
	perChainWG.Add(len(listPerChainID))
	for chainID, list := range listPerChainID {
		go func(chainID uint64, list []TDefillamaList) {
			defer perChainWG.Done()
			syncMapRaw, _ := tokensForChainIDSyncMap.Load(chainID)
			syncMap := syncMapRaw.([]TokenListToken)

			listOfAddresses := []common.Address{}
			for _, token := range list {
				listOfAddresses = append(listOfAddresses, token.Address)
			}

			decimalsInfo := ethereum.FetchDecimals(chainID, listOfAddresses)
			for _, token := range list {
				if decimals, ok := decimalsInfo[token.Address.Hex()]; ok {
					if newToken, err := SetToken(
						token.Address,
						token.Name,
						token.Symbol,
						token.LogoURI,
						chainID,
						int(decimals),
					); err == nil {
						tokensForChainIDSyncMap.Store(chainID, append(syncMap, newToken))
					}
				}
			}
		}(chainID, list)
	}
	perChainWG.Wait()

	return extractSyncMap(tokensForChainIDSyncMap)
}

func fetchDefillamaTokenList() []TokenListToken {
	list := helpers.FetchJSON[[]TDefillamaList](`https://defillama-datasets.llama.fi/tokenlist/all.json`)
	listPerChainID := make(map[uint64][]TDefillamaList)
	for _, v := range list {
		if len(v.Platforms) == 0 {
			continue
		}
		for platformName, addressOnPlatform := range v.Platforms {
			chainID := defillamaMapNetworkNameToChainID(platformName)
			if chainID == 0 || helpers.IsChainIDIgnored(chainID) {
				continue
			}
			if helpers.IsIgnoredToken(chainID, common.HexToAddress(addressOnPlatform)) {
				continue
			}
			if _, ok := listPerChainID[chainID]; !ok {
				listPerChainID[chainID] = []TDefillamaList{}
			}
			v.Address = common.HexToAddress(addressOnPlatform)
			listPerChainID[chainID] = append(listPerChainID[chainID], v)
		}
	}
	return handleDefillamaTokenList(listPerChainID)
}

func buildDefillamaTokenList() {
	tokenList := loadTokenListFromJsonFile(`defillama.json`)
	tokenList.Name = "DefiLlama"
	tokenList.LogoURI = "https://wiki.defillama.com/w/resources/assets/wiki.png?88de1"

	tokens := fetchDefillamaTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `defillama.json`, Standard)
}
