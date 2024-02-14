package main

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
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

func handleCurveTokenList(listPerChainID map[uint64][]TCurveTokenData) []models.TokenListToken {
	tokensForChainIDSyncMap := helpers.InitSyncMap(listPerChainID)

	// Fetch the basic informations for all the tokens for all the chains
	perChainWG := sync.WaitGroup{}
	perChainWG.Add(len(listPerChainID))
	for chainID, list := range listPerChainID {
		go func(chainID uint64, list []TCurveTokenData) {
			defer perChainWG.Done()
			syncMapRaw, _ := tokensForChainIDSyncMap.Load(chainID)
			syncMap := syncMapRaw.([]models.TokenListToken)

			listOfAddresses := []common.Address{}
			for _, token := range list {
				if !chains.IsTokenIgnored(chainID, common.HexToAddress(token.Address)) {
					listOfAddresses = append(listOfAddresses, common.HexToAddress(token.Address))
				}
				if !chains.IsTokenIgnored(chainID, common.HexToAddress(token.LpTokenAddress)) {
					listOfAddresses = append(listOfAddresses, common.HexToAddress(token.LpTokenAddress))
				}
				for _, coinAddress := range token.CoinsAddresses {
					if !chains.IsTokenIgnored(chainID, common.HexToAddress(coinAddress)) {
						listOfAddresses = append(listOfAddresses, common.HexToAddress(coinAddress))
					}
				}
			}

			tokensInfo := helpers.RetrieveBasicInformations(chainID, listOfAddresses)
			for _, address := range listOfAddresses {
				if token, ok := tokensInfo[address.Hex()]; ok {
					if newToken, err := helpers.SetToken(
						token.Address,
						token.Name,
						token.Symbol,
						``,
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

func fetchCurveTokenList() []models.TokenListToken {
	listPerChainID := make(map[uint64][]TCurveTokenData)

	for chainID, uris := range APIURIForCurve {
		if !chains.IsChainIDSupported(chainID) {
			continue
		}

		for _, uri := range uris {
			list := helpers.FetchJSON[TCurveList](uri)
			listPerChainID[chainID] = append(listPerChainID[chainID], list.Data.PoolData...)
		}
	}

	return handleCurveTokenList(listPerChainID)
}

func buildCurveTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`curve.json`)
	tokenList.Name = "Curve Token List"
	tokenList.LogoURI = "https://classic.curve.fi/logo.png"

	tokens := fetchCurveTokenList()
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `curve.json`, helpers.SavingMethodStandard)
}
