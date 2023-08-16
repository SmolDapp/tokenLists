package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

type TBebopTokenData struct {
	Name         string `json:"name"`
	Symbol       string `json:"ticker"`
	Address      string `json:"contractAddress"`
	ChainID      uint64 `json:"chainId"`
	Decimals     int    `json:"decimals"`
	Availability struct {
		IsAvailable bool `json:"isAvailable"`
		CanBuy      bool `json:"canBuy"`
		CanSell     bool `json:"canSell"`
	} `json:"availability"`
}
type TBebopList struct {
	Tokens map[string]TBebopTokenData
}

func bebopMapNetworkChainIDToName(chainID uint64) string {
	switch chainID {
	case 1:
		return `ethereum`
	case 137:
		return `polygon`
	case 42161:
		return `arbitrum`
	}
	return `ethereum`
}

func fetchbebopTokenList() []TokenListToken {
	supportedChainID := []uint64{1, 137, 42161}
	tokens := []TokenListToken{}

	tokenList := helpers.FetchJSON[TokenListData](`https://api.bebop.xyz/token_list`)
	tokenMap := map[string]TokenListToken{}
	for _, token := range tokenList.Tokens {
		tokenMap[token.Address] = token
	}

	for _, chainID := range supportedChainID {
		list := helpers.FetchJSON[TBebopList](`https://api.bebop.xyz/` + bebopMapNetworkChainIDToName(chainID) + `/v1/token-info`)
		for _, token := range list.Tokens {
			if !token.Availability.IsAvailable {
				continue
			}
			if !token.Availability.CanBuy && !token.Availability.CanSell {
				continue
			}
			logoURI := ``
			if tokenFromList, ok := tokenMap[token.Address]; ok {
				logoURI = tokenFromList.LogoURI
			}

			if newToken, err := SetToken(
				common.HexToAddress(token.Address),
				token.Name,
				token.Symbol,
				logoURI,
				token.ChainID,
				int(token.Decimals),
			); err == nil {
				tokens = append(tokens, newToken)
			}
		}
	}
	return tokens
}

func buildBebopTokenList() {
	tokenList := loadTokenListFromJsonFile(`bebop.json`)
	tokenList.Name = "Bebop"
	tokenList.LogoURI = "https://bebop-public-images.s3.eu-west-2.amazonaws.com/bebop-logo.png"

	tokens := fetchbebopTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `bebop.json`, Standard)
}
