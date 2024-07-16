package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
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
	LogoURI           string                           `json:"logoURI"`
	Keywords          []string                         `json:"keywords"`
	Tokens            []models.TokenListToken          `json:"tokens"`
	PreviousTokensMap map[string]models.TokenListToken `json:"-"`
	NextTokensMap     map[string]models.TokenListToken `json:"-"`
	Metadata          map[string]interface{}           `json:"metadata,omitempty"`
}

func bebopMapNetworkChainIDToName(chainID uint64) string {
	switch chainID {
	case 1:
		return `ethereum`
	case 10:
		return `optimism`
	case 56:
		return `bsc`
	case 137:
		return `polygon`
	case 324:
		return `zksync`
	case 8453:
		return `base`
	case 42161:
		return `arbitrum`
	case 81457:
		return `blast`
	}
	return `ethereum`
}

func fetchBebopTokenList() []models.TokenListToken {
	supportedChainID := []uint64{1, 10, 56, 137, 8453, 42161, 81457}
	tokens := []models.TokenListToken{}

	type TBebopTokenListToken struct {
		models.TokenListToken
		Tags       []string `json:"tags"`
		Extensions struct {
			Color           string `json:"color"`
			DisplayDecimals int    `json:"displayDecimals"`
		} `json:"extensions"`
	}

	for _, chainID := range supportedChainID {
		list := helpers.FetchJSON[models.TokenListData[TBebopTokenListToken]](`https://api.bebop.xyz/pmm/` + bebopMapNetworkChainIDToName(chainID) + `/v3/tokenlist?active_only=true`)

		tokenMap := map[string]TBebopTokenListToken{}
		tokenList := []common.Address{}
		for _, token := range list.Tokens {
			tokenMap[token.Address] = token
			tokenList = append(tokenList, common.HexToAddress(token.Address))
		}

		tokensInfo := helpers.RetrieveBasicInformations(chainID, tokenList)
		for _, existingToken := range list.Tokens {
			if token, ok := tokensInfo[common.HexToAddress(existingToken.Address).Hex()]; ok {
				if newToken, err := helpers.SetToken(
					token.Address,
					helpers.SafeString(token.Name, existingToken.Name),
					helpers.SafeString(token.Symbol, existingToken.Symbol),
					helpers.SafeString(existingToken.LogoURI, ""),
					chainID,
					int(token.Decimals),
				); err == nil {
					newToken.Metadata = map[string]any{
						`tags`:            existingToken.Tags,
						`color`:           existingToken.Extensions.Color,
						`displayDecimals`: existingToken.Extensions.DisplayDecimals,
					}
					tokens = append(tokens, newToken)
				}
			}
		}
	}
	return tokens
}

func buildBebopTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`bebop.json`)
	tokenList.Name = "Bebop"
	tokenList.LogoURI = "https://bebop-public-images.s3.eu-west-2.amazonaws.com/bebop-logo.png"

	tokens := fetchBebopTokenList()
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `bebop.json`, helpers.SavingMethodStandard)
}
