package main

import (
	"math"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

type TSmolAssetsList struct {
	Version struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	}
	Tokens []string `json:"tokens"`
}

func buildTokenListooorList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`tokenlistooor.json`)
	tokenList.Name = `Tokenlistooor Token List`
	tokenList.LogoURI = `https://raw.githubusercontent.com/smoldapp/tokenLists/main/.github/tokenlistooor.svg`
	tokenList.Description = `A curated list of tokens from all the token lists on tokenlistooor.`

	/**************************************************************************
	** Create a map of all tokens from all lists and only add the missing ones
	** in it. Map are WAY faster than arrays fir our use case
	**************************************************************************/
	allTokens := make(map[uint64]map[string]models.TokenListToken)
	allTokensPlain := []models.TokenListToken{}
	listsPerChain := make(map[uint64][]string)

	for _, chain := range chains.CHAINS {
		allTokensPlain = append(allTokensPlain, chain.Coin)
	}

	/**************************************************************************
	** We want to know which tokens to add to the aggregated tokenlistooor list
	** and to do that we need to know in how many lists they are present.
	** This is chain sensitive: we need a token to be available in at least
	** 50% of the lists for a given chain to be added to the aggregated list.
	**************************************************************************/
	for name, generatorData := range GENERATORS {
		if name == `tokenlistooor` {
			continue
		}
		if generatorData.GeneratorType == GeneratorPool {
			continue
		}

		initialCount := 1
		shouldByPassCount := name == `yearn` || name == `smolAssets`
		if shouldByPassCount {
			initialCount = math.MaxInt32
		}
		tokenList := helpers.LoadTokenListFromJsonFile(name + `.json`)
		for _, token := range tokenList.Tokens {
			if !chains.IsChainIDSupported(token.ChainID) {
				continue
			}
			if _, ok := listsPerChain[token.ChainID]; !ok {
				listsPerChain[token.ChainID] = []string{}
			}

			if !helpers.Includes(listsPerChain[token.ChainID], name) {
				listsPerChain[token.ChainID] = append(listsPerChain[token.ChainID], name)
			}

			if _, ok := allTokens[token.ChainID]; !ok {
				allTokens[token.ChainID] = make(map[string]models.TokenListToken)
			}

			if existingToken, ok := allTokens[token.ChainID][helpers.ToAddress(token.Address)]; ok {
				newOccurence := existingToken.Occurrence
				if newOccurence != math.MaxInt32 {
					foundInExtraTokens := false
					for _, extraToken := range chains.CHAINS[token.ChainID].ExtraTokens {
						if common.HexToAddress(token.Address) == extraToken {
							foundInExtraTokens = true
							break
						}
					}
					if foundInExtraTokens {
						newOccurence = math.MaxInt32
					} else if strings.HasSuffix(name, `-static`) {
						newOccurence = math.MaxInt32
					} else {
						newOccurence += 1
					}
				}

				allTokens[token.ChainID][helpers.ToAddress(token.Address)] = models.TokenListToken{
					Address:    existingToken.Address,
					Name:       helpers.SafeString(existingToken.Name, token.Name),
					Symbol:     helpers.SafeString(existingToken.Symbol, token.Symbol),
					LogoURI:    helpers.SafeString(existingToken.LogoURI, token.LogoURI),
					Decimals:   helpers.SafeInt(existingToken.Decimals, token.Decimals),
					ChainID:    token.ChainID,
					Occurrence: newOccurence,
				}
			} else {
				tokenInitialOccurence := initialCount
				for _, extraToken := range chains.CHAINS[token.ChainID].ExtraTokens {
					if common.HexToAddress(token.Address) == extraToken {
						tokenInitialOccurence = math.MaxInt32
					}
				}
				if strings.HasSuffix(name, `-static`) {
					tokenInitialOccurence = math.MaxInt32
				}

				allTokens[token.ChainID][helpers.ToAddress(token.Address)] = models.TokenListToken{
					Address:    helpers.ToAddress(token.Address),
					Name:       helpers.SafeString(token.Name, ``),
					Symbol:     helpers.SafeString(token.Symbol, ``),
					LogoURI:    helpers.SafeString(token.LogoURI, ``),
					Decimals:   helpers.SafeInt(token.Decimals, 18),
					ChainID:    token.ChainID,
					Occurrence: tokenInitialOccurence,
				}
			}
		}
	}

	for chainID, tokens := range allTokens {
		for _, token := range tokens {
			if _, ok := listsPerChain[chainID]; !ok {
				continue
			}
			chainCount := len(listsPerChain[uint64(chainID)])
			if token.Occurrence >= int(math.Ceil(float64(chainCount)*0.5)) {
				allTokensPlain = append(allTokensPlain, token)
			}
		}
	}

	tokens := helpers.GetTokensFromList(allTokensPlain)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `tokenlistooor.json`, helpers.SavingMethodStandard)
}
