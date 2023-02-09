package main

import (
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

type TMessariContractAddresses struct {
	Platform        string `json:"platform"`
	ContractAddress string `json:"contract_address"`
}
type TMessariTokenData struct {
	ID        string                      `json:"id"`
	Name      string                      `json:"name"`
	Symbol    string                      `json:"symbol"`
	Addresses []TMessariContractAddresses `json:"contract_addresses"`
}
type TMessariList struct {
	Tokens []TMessariTokenData `json:"data,omitempty"`
}

func messariMapNetworkNameToChainID(network string) uint64 {
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

func handleMessariTokenList(tokensPerChainID map[uint64]map[string]TokenListToken) []TokenListToken {
	tokensForChainIDSyncMap := initSyncMap(tokensPerChainID)

	// Fetch the basic informations for all the tokens for all the chains
	perChainWG := sync.WaitGroup{}
	perChainWG.Add(len(tokensPerChainID))
	for chainID, initialList := range tokensPerChainID {
		listOfAddresses := []common.Address{}
		for _, token := range initialList {
			listOfAddresses = append(listOfAddresses, common.HexToAddress(token.Address))
		}

		go func(chainID uint64, list []common.Address, initialList map[string]TokenListToken) {
			defer perChainWG.Done()
			syncMapRaw, _ := tokensForChainIDSyncMap.Load(chainID)
			syncMap := syncMapRaw.([]TokenListToken)

			tokensInfo := ethereum.FetchDecimals(chainID, list)
			for _, address := range list {
				if decimals, ok := tokensInfo[address.Hex()]; ok {
					if initialToken, ok := initialList[address.Hex()]; ok {
						if newToken, err := SetToken(
							address,
							initialToken.Name,
							initialToken.Symbol,
							initialToken.LogoURI,
							initialToken.ChainID,
							int(decimals),
						); err == nil {
							syncMap = append(syncMap, newToken)
							tokensForChainIDSyncMap.Store(chainID, syncMap)
						}
					}
				}
			}
		}(chainID, listOfAddresses, initialList)

	}
	perChainWG.Wait()

	return extractSyncMap(tokensForChainIDSyncMap)
}

func fetchMessariTokenList() []TokenListToken {
	limit := 500
	page := 1
	tokensPerChainID := make(map[uint64]map[string]TokenListToken)

	for {
		uri := `https://data.messari.io/api/v2/assets?fields=name,symbol,contract_addresses,id&sort=id&limit=` + strconv.FormatInt(int64(limit), 10) + `&page=` + strconv.FormatInt(int64(page), 10)
		list := helpers.FetchJSON[TMessariList](uri)

		if list.Tokens == nil || len(list.Tokens) == 0 {
			break
		}
		for _, token := range list.Tokens {
			logoURI := `https://asset-images.messari.io/images/` + token.ID + `/128.png`
			for _, platformData := range token.Addresses {
				chainID := messariMapNetworkNameToChainID(platformData.Platform)
				if chainID == 0 || helpers.IsChainIDIgnored(chainID) {
					continue
				}
				if helpers.IsIgnoredToken(chainID, common.HexToAddress(platformData.ContractAddress)) {
					continue
				}
				if _, ok := tokensPerChainID[chainID]; !ok {
					tokensPerChainID[chainID] = make(map[string]TokenListToken)
				}

				if newToken, err := SetToken(
					common.HexToAddress(platformData.ContractAddress),
					token.Name,
					token.Symbol,
					logoURI,
					chainID,
					18,
				); err == nil {
					tokensPerChainID[chainID][helpers.ToAddress(platformData.ContractAddress)] = newToken
				}
			}
		}
		page++
		//20 requests per minutes, so 3 seconds between each request and we should be fine
		time.Sleep(3 * time.Second)
	}

	return handleMessariTokenList(tokensPerChainID)
}

func buildMessariTokenList() {
	tokenList := loadTokenListFromJsonFile(`messari.json`)
	tokenList.Name = "Messari Token List"
	tokenList.LogoURI = "https://messari.io/images/logo_tcr-check.svg"

	tokens := fetchMessariTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `messari.json`, Standard)
}
