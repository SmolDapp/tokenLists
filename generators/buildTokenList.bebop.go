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

type TBebopTokenListData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
	Version     struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	} `json:"version"`
	LogoURI           string                    `json:"logoURI"`
	Keywords          []string                  `json:"keywords"`
	Tokens            []TokenListToken          `json:"tokens"`
	PreviousTokensMap map[string]TokenListToken `json:"-"`
	NextTokensMap     map[string]TokenListToken `json:"-"`
	Metadata          map[string]interface{}    `json:"metadata,omitempty"`
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

	type TBebopTokenListToken struct {
		TokenListToken
		Tags       []string `json:"tags"`
		Extensions struct {
			Color           string `json:"color"`
			DisplayDecimals int    `json:"displayDecimals"`
		} `json:"extensions"`
	}

	tokenList := helpers.FetchJSON[TokenListData[TBebopTokenListToken]](`https://api.bebop.xyz/token_list`)
	tokenMap := map[string]TBebopTokenListToken{}
	for _, token := range tokenList.Tokens {
		tokenMap[token.Address] = token
	}

	for _, chainID := range supportedChainID {
		list := helpers.FetchJSON[TBebopList](`https://api.bebop.xyz/` + bebopMapNetworkChainIDToName(chainID) + `/v2/token-info`)

		allTokens := []common.Address{}
		for _, token := range list.Tokens {
			if !token.Availability.IsAvailable {
				continue
			}
			if !token.Availability.CanBuy && !token.Availability.CanSell {
				continue
			}
			allTokens = append(allTokens, common.HexToAddress(token.Address))
		}

		tokensInfo := retrieveBasicInformations(chainID, allTokens)
		for _, existingToken := range list.Tokens {
			if !existingToken.Availability.IsAvailable {
				continue
			}
			if !existingToken.Availability.CanBuy && !existingToken.Availability.CanSell {
				continue
			}
			logoURI := ``
			if tokenFromList, ok := tokenMap[existingToken.Address]; ok {
				logoURI = tokenFromList.LogoURI
			}

			if token, ok := tokensInfo[existingToken.Address]; ok {
				if newToken, err := SetToken(
					token.Address,
					helpers.SafeString(token.Name, existingToken.Name),
					helpers.SafeString(token.Symbol, existingToken.Symbol),
					logoURI,
					chainID,
					int(token.Decimals),
				); err == nil {
					if tokenFromList, ok := tokenMap[existingToken.Address]; ok {
						newToken.Metadata = map[string]any{
							`tags`:            tokenFromList.Tags,
							`color`:           tokenFromList.Extensions.Color,
							`displayDecimals`: tokenFromList.Extensions.DisplayDecimals,
						}
					}
					tokens = append(tokens, newToken)
				}
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
