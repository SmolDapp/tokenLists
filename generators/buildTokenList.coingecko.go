package main

import (
	"strconv"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/common/utils"
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
	list := helpers.FetchJSON[models.TokenListData[models.TokenListToken]](`https://tokens.coingecko.com/uniswap/all.json`)
	for _, v := range list.Tokens {
		chainIDStr := strconv.FormatInt(int64(v.ChainID), 10)
		logoURIList[chainIDStr+`_`+common.HexToAddress(v.Address).Hex()] = v.LogoURI
	}
	return logoURIList
}

func handleCoingeckoTokenList(tokensPerChainID map[uint64][]string) []models.TokenListToken {
	logoURIs := fetchCoingeckoLegacyListLogoURI()
	tokensForChainIDSyncMap := helpers.InitSyncMap(tokensPerChainID)

	// Fetch the basic informations for all the tokens for all the chains
	perChainWG := sync.WaitGroup{}
	perChainWG.Add(len(tokensPerChainID))
	for chainID, list := range tokensPerChainID {
		go func(chainID uint64, list []string) {
			defer perChainWG.Done()
			syncMapRaw, _ := tokensForChainIDSyncMap.Load(chainID)
			syncMap := syncMapRaw.([]models.TokenListToken)

			tokensInfo := helpers.RetrieveBasicInformations(chainID, list)
			for _, address := range list {
				if token, ok := tokensInfo[utils.ToAddress(address)]; ok {
					if newToken, err := helpers.SetToken(
						token.Address,
						token.Name,
						token.Symbol,
						helpers.SafeString(logoURIs[strconv.FormatInt(int64(chainID), 10)+`_`+common.HexToAddress(token.Address).Hex()], ``),
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

	return helpers.ExtractSyncMap(tokensForChainIDSyncMap)
}

func fetchCoingeckoTokenList() []models.TokenListToken {
	tokensPerChainID := make(map[uint64][]string)
	list := helpers.FetchJSON[[]TCoingeckoList](`https://api.coingecko.com/api/v3/coins/list?include_platform=true`)

	for _, v := range list {
		if len(v.Platforms) == 0 {
			continue
		}

		for platformName, addressOnPlatform := range v.Platforms {
			chainID := coingeckoMapNetworkNameToChainID(platformName)
			if !chains.IsChainIDSupported(chainID) {
				continue
			}
			if chains.IsTokenIgnored(chainID, common.HexToAddress(addressOnPlatform).Hex()) {
				continue
			}
			if _, ok := tokensPerChainID[chainID]; !ok {
				tokensPerChainID[chainID] = []string{}
			}
			tokensPerChainID[chainID] = append(tokensPerChainID[chainID], common.HexToAddress(addressOnPlatform).Hex())
		}
	}
	return handleCoingeckoTokenList(tokensPerChainID)
}

func buildCoingeckoTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`coingecko.json`)
	tokenList.Name = "CoinGecko"
	tokenList.Keywords = []string{"coingecko", "defi"}
	tokenList.LogoURI = "https://static.coingecko.com/s/about/gecko-1b23cd303298d7474345b1938c21fdb20c71f4f399eefa8637ad243b8ac5dbf5.png"

	tokens := fetchCoingeckoTokenList()
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `coingecko.json`, helpers.SavingMethodStandard)
}
