package main

import (
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
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
	list := helpers.FetchJSON[TokenListData](`https://tokens.coingecko.com/uniswap/all.json`)
	for _, v := range list.Tokens {
		chainIDStr := strconv.FormatInt(int64(v.ChainID), 10)
		logoURIList[chainIDStr+`_`+common.HexToAddress(v.Address).Hex()] = v.LogoURI
	}
	return logoURIList
}

func handleCoingeckoTokenList(tokensPerChainID map[uint64][]common.Address) []TokenListToken {
	logoURIs := fetchCoingeckoLegacyListLogoURI()
	tokensForChainIDSyncMap := initSyncMap(tokensPerChainID)

	// Fetch the basic informations for all the tokens for all the chains
	perChainWG := sync.WaitGroup{}
	perChainWG.Add(len(tokensPerChainID))
	for chainID, list := range tokensPerChainID {
		go func(chainID uint64, list []common.Address) {
			defer perChainWG.Done()
			syncMapRaw, _ := tokensForChainIDSyncMap.Load(chainID)
			syncMap := syncMapRaw.([]TokenListToken)

			tokensInfo := ethereum.FetchBasicInformations(chainID, list)
			for _, address := range list {
				if token, ok := tokensInfo[address.Hex()]; ok {
					if newToken, err := SetToken(
						token.Address,
						token.Name,
						token.Symbol,
						helpers.SafeString(logoURIs[strconv.FormatInt(int64(chainID), 10)+`_`+token.Address.Hex()], ``),
						chainID,
						int(token.Decimals),
					); err == nil {
						syncMap = append(syncMap, newToken)
						tokensForChainIDSyncMap.Store(chainID, syncMap)
					}
				}
			}
		}(chainID, list)
	}
	perChainWG.Wait()

	return extractSyncMap(tokensForChainIDSyncMap)
}

func fetchCoingeckoTokenList() []TokenListToken {
	tokensPerChainID := make(map[uint64][]common.Address)
	list := helpers.FetchJSON[[]TCoingeckoList](`https://api.coingecko.com/api/v3/coins/list?include_platform=true`)

	for _, v := range list {
		if len(v.Platforms) == 0 {
			continue
		}

		for platformName, addressOnPlatform := range v.Platforms {
			chainID := coingeckoMapNetworkNameToChainID(platformName)
			if chainID == 0 || helpers.IsChainIDIgnored(chainID) {
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
