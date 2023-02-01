package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

type TCoingeckoList struct {
	ID        string            `json:"id"`
	Symbol    string            `json:"symbol"`
	Name      string            `json:"name"`
	Platforms map[string]string `json:"platforms"`
}

func coingeckoMapNetworkNameToChainID(network string) uint64 {
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

func fetchCoingeckoLegacyListLogoURI() map[string]string {
	logoURIList := make(map[string]string)
	resp, err := http.Get(`https://tokens.coingecko.com/uniswap/all.json`)
	if err != nil {
		return logoURIList
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return logoURIList
	}

	var list TokenListData
	if err := json.Unmarshal(body, &list); err != nil {
		return logoURIList
	}

	for _, v := range list.Tokens {
		chainIDStr := strconv.FormatInt(int64(v.ChainID), 10)
		logoURIList[chainIDStr+`_`+common.HexToAddress(v.Address).Hex()] = v.LogoURI
	}
	return logoURIList
}

func handleCoingeckoTokenList(tokensPerChainID map[uint64][]common.Address) []TokenListToken {
	logoURIs := fetchCoingeckoLegacyListLogoURI()
	tokensForChainID := make(map[uint64][]TokenListToken)

	// Initialize the map before the goroutines to avoid `fatal error: concurrent map writes`
	for chainID := range tokensPerChainID {
		if _, ok := tokensForChainID[chainID]; !ok {
			tokensForChainID[chainID] = []TokenListToken{}
		}
	}

	// Fetch the basic informations for all the tokens for all the chains
	perChainWG := sync.WaitGroup{}
	perChainWG.Add(len(tokensPerChainID))
	for chainID, list := range tokensPerChainID {
		go func(chainID uint64, list []common.Address) {
			defer perChainWG.Done()
			tokensInfo := fetchBasicInformations(chainID, list)
			for _, address := range list {
				if token, ok := tokensInfo[address.Hex()]; ok {
					if (token.Name == `` || token.Symbol == `` || token.Decimals == 0) || helpers.IsIgnoredToken(chainID, token.Address) {
						continue
					}

					logoURI := ``
					if v, ok := logoURIs[strconv.FormatInt(int64(chainID), 10)+`_`+token.Address.Hex()]; ok {
						logoURI = v
					}

					tokensForChainID[chainID] = append(tokensForChainID[chainID], TokenListToken{
						ChainID:  int(chainID),
						Address:  token.Address.Hex(),
						Name:     token.Name,
						Symbol:   token.Symbol,
						Decimals: int(token.Decimals),
						LogoURI:  logoURI,
					})
				}
			}
		}(chainID, list)
	}
	perChainWG.Wait()

	tokenList := []TokenListToken{}
	// Merge all the tokens for all the chains
	for chainID := range tokensPerChainID {
		tokenList = append(tokenList, tokensForChainID[chainID]...)
	}
	return tokenList
}

func fetchCoingeckoTokenList() []TokenListToken {
	tokensPerChainID := make(map[uint64][]common.Address)
	resp, err := http.Get(`https://api.coingecko.com/api/v3/coins/list?include_platform=true`)
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
		logs.Error(`impossible to fetch Coingecko token list`)
		return []TokenListToken{}
	}

	list := []TCoingeckoList{}
	if err := json.Unmarshal(body, &list); err != nil {
		logs.Error(err)
		return []TokenListToken{}
	}

	for _, v := range list {
		if len(v.Platforms) == 0 {
			continue
		}

		for platformName, addressOnPlatform := range v.Platforms {
			chainID := coingeckoMapNetworkNameToChainID(platformName)
			if chainID == 0 || helpers.IsIgnoredChain(chainID) {
				continue
			}
			if helpers.IsIgnoredToken(chainID, common.HexToAddress(addressOnPlatform)) {
				continue
			}
			if _, ok := tokensPerChainID[chainID]; !ok {
				tokensPerChainID[chainID] = []common.Address{}
			}
			tokensPerChainID[chainID] = append(tokensPerChainID[chainID], common.HexToAddress(addressOnPlatform))
		}
	}
	return handleCoingeckoTokenList(tokensPerChainID)
}

func buildCoingeckoTokenList() {
	tokenList := loadTokenListFromJsonFile(`coingecko.json`)
	tokenList.Name = "CoinGecko"
	tokenList.Keywords = []string{"coingecko", "defi"}
	tokenList.LogoURI = "https://www.coingecko.com/assets/thumbnail-007177f3eca19695592f0b8b0eabbdae282b54154e1be912285c9034ea6cbaf2.png"

	tokens := fetchCoingeckoTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `coingecko.json`, Standard)
}
