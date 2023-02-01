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

type TCurveTokenData struct {
	Address        string   `json:"address"`
	CoinsAddresses []string `json:"coinsAddresses"`
	LpTokenAddress string   `json:"lpTokenAddress"`
	LpTokenName    string   `json:"name"`
	LpTokenSymbol  string   `json:"symbol"`
}
type TCurveList struct {
	Data struct {
		PoolData []TCurveTokenData `json:"poolData"`
	} `json:"data"`
}

var APIURIForCurve = map[uint64][]string{
	1: {
		`https://api.curve.fi/api/getPools/ethereum/main`,
		`https://api.curve.fi/api/getPools/ethereum/crypto`,
		`https://api.curve.fi/api/getPools/ethereum/factory`,
		`https://api.curve.fi/api/getPools/ethereum/factory-crypto`,
	},
	10: {
		`https://api.curve.fi/api/getPools/optimism/main`,
		`https://api.curve.fi/api/getPools/optimism/crypto`,
		`https://api.curve.fi/api/getPools/optimism/factory`,
	},
	100: {
		`https://api.curve.fi/api/getPools/xdai/main`,
	},
	137: {
		`https://api.curve.fi/api/getPools/polygon/main`,
		`https://api.curve.fi/api/getPools/polygon/crypto`,
		`https://api.curve.fi/api/getPools/polygon/factory`,
	},
	250: {
		`https://api.curve.fi/api/getPools/fantom/main`,
		`https://api.curve.fi/api/getPools/fantom/crypto`,
		`https://api.curve.fi/api/getPools/fantom/factory`,
	},
	42161: {
		`https://api.curve.fi/api/getPools/arbitrum/main`,
		`https://api.curve.fi/api/getPools/arbitrum/crypto`,
		`https://api.curve.fi/api/getPools/arbitrum/factory`,
	},
	43114: {
		`https://api.curve.fi/api/getPools/avalanche/main`,
		`https://api.curve.fi/api/getPools/avalanche/crypto`,
		`https://api.curve.fi/api/getPools/avalanche/factory`,
	},
}

func handleCurveTokenList(listPerChainID map[uint64][]TCurveTokenData) []TokenListToken {
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
		go func(chainID uint64, list []TCurveTokenData) {
			defer perChainWG.Done()

			listOfAddresses := []common.Address{}
			for _, token := range list {
				if !helpers.IsIgnoredToken(chainID, common.HexToAddress(token.Address)) {
					listOfAddresses = append(listOfAddresses, common.HexToAddress(token.Address))
				}
				if !helpers.IsIgnoredToken(chainID, common.HexToAddress(token.LpTokenAddress)) {
					listOfAddresses = append(listOfAddresses, common.HexToAddress(token.LpTokenAddress))
				}
				for _, coinAddress := range token.CoinsAddresses {
					if !helpers.IsIgnoredToken(chainID, common.HexToAddress(coinAddress)) {
						listOfAddresses = append(listOfAddresses, common.HexToAddress(coinAddress))
					}
				}
			}

			tokensInfo := fetchBasicInformations(chainID, listOfAddresses)
			for _, token := range listOfAddresses {
				if tokenInfo, ok := tokensInfo[token.Hex()]; ok {
					if (tokenInfo.Name == `` || tokenInfo.Symbol == `` || tokenInfo.Decimals == 0) || helpers.IsIgnoredToken(chainID, token) {
						continue
					}
					tokensForChainID[chainID] = append(tokensForChainID[chainID], TokenListToken{
						ChainID:  int(chainID),
						Address:  token.Hex(),
						Name:     tokenInfo.Name,
						Symbol:   tokenInfo.Symbol,
						Decimals: int(tokenInfo.Decimals),
						LogoURI:  ``,
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

func fetchCurveTokenList() []TokenListToken {
	listPerChainID := make(map[uint64][]TCurveTokenData)

	for chainID, uris := range APIURIForCurve {
		if helpers.IsIgnoredChain(chainID) {
			continue
		}

		for _, uri := range uris {
			resp, err := http.Get(uri)
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
				logs.Error(`impossible to fetch Curve token list`)
				return []TokenListToken{}
			}

			var list TCurveList
			if err := json.Unmarshal(body, &list); err != nil {
				logs.Error(err)
				return []TokenListToken{}
			}

			listPerChainID[chainID] = append(listPerChainID[chainID], list.Data.PoolData...)
		}
	}

	return handleCurveTokenList(listPerChainID)
}

func buildCurveTokenList() {
	tokenList := loadTokenListFromJsonFile(`curve.json`)
	tokenList.Name = "Curve Token List"
	tokenList.LogoURI = "https://classic.curve.fi/logo.png"

	tokens := fetchCurveTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `curve.json`, Standard)
}
