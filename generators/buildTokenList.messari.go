package main

import (
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
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

func fetchMessariTokenList() []TokenListToken {
	limit := 500
	page := 1
	allTokens := []TokenListToken{}

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
				if chainID == 0 || !helpers.IsChainIDSupported(chainID) {
					continue
				}
				if helpers.IsIgnoredToken(chainID, common.HexToAddress(platformData.ContractAddress)) {
					continue
				}
				if newToken, err := SetToken(
					common.HexToAddress(platformData.ContractAddress),
					token.Name,
					token.Symbol,
					logoURI,
					chainID,
					18,
				); err == nil {
					allTokens = append(allTokens, newToken)
				}
			}
		}
		page++
		//20 requests per minutes, so 3 seconds between each request and we should be fine
		time.Sleep(3 * time.Second)
	}

	return fetchTokenList(allTokens)
}

func buildMessariTokenList() {
	tokenList := loadTokenListFromJsonFile(`messari.json`)
	tokenList.Name = "Messari Token List"
	tokenList.LogoURI = "https://messari.io/images/logo_tcr-check.svg"

	tokens := fetchMessariTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `messari.json`, Standard)
}
