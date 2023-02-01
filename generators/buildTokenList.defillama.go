package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
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
	tokensForChainID := make(map[uint64][]TokenListToken)

	// Initialize the map before the goroutines to avoid `fatal error: concurrent map writes`
	for chainID := range listPerChainID {
		if _, ok := tokensForChainID[chainID]; !ok {
			tokensForChainID[chainID] = []TokenListToken{}
		}
	}

	// Fetch the basic informations for all the tokens for all the chains
	perChainWG := sync.WaitGroup{}
	perChainWG.Add(len(listPerChainID))
	for chainID, list := range listPerChainID {
		go func(chainID uint64, list []TDefillamaList) {
			defer perChainWG.Done()

			listOfAddresses := []common.Address{}
			for _, token := range list {
				listOfAddresses = append(listOfAddresses, token.Address)
			}

			decimalsInfo := fetchDecimals(chainID, listOfAddresses)
			for _, token := range list {
				if decimal, ok := decimalsInfo[token.Address.Hex()]; ok {
					if (token.Name == `` || token.Symbol == `` || decimal == 0) || helpers.IsIgnoredToken(chainID, token.Address) {
						continue
					}
					tokensForChainID[chainID] = append(tokensForChainID[chainID], TokenListToken{
						ChainID:  int(chainID),
						Address:  token.Address.Hex(),
						Name:     token.Name,
						Symbol:   token.Symbol,
						Decimals: int(decimal),
						LogoURI:  token.LogoURI,
					})
				}
			}
		}(chainID, list)
	}
	perChainWG.Wait()

	tokenList := []TokenListToken{}
	// Merge all the tokens for all the chains
	for chainID := range listPerChainID {
		tokenList = append(tokenList, tokensForChainID[chainID]...)
	}
	return tokenList
}

func fetchDefillamaTokenList() []TokenListToken {
	resp, err := http.Get(`https://defillama-datasets.llama.fi/tokenlist/all.json`)
	if err != nil {
		logs.Error(err)
		return []TokenListToken{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return []TokenListToken{}
	}
	if (resp.StatusCode < 200) || (resp.StatusCode > 299) {
		logs.Error(`impossible to fetch DefiLlama token list`)
		return []TokenListToken{}
	}

	list := []TDefillamaList{}
	if err := json.Unmarshal(body, &list); err != nil {
		logs.Error(err)
		return []TokenListToken{}
	}

	listPerChainID := make(map[uint64][]TDefillamaList)
	for _, v := range list {
		if len(v.Platforms) == 0 {
			continue
		}
		for platformName, addressOnPlatform := range v.Platforms {
			chainID := defillamaMapNetworkNameToChainID(platformName)
			if chainID == 0 || helpers.IsIgnoredChain(chainID) {
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
