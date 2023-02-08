package main

import (
	"context"
	"math/big"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/contracts"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

type TUniContracts struct {
	ContractAddress common.Address
	BlockNumber     *big.Int
	Type            int
	ShouldFetch     bool
}

var UniswapContractsPerChainID = map[uint64][]TUniContracts{
	1: {
		{
			ContractAddress: common.HexToAddress(`0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f`),
			BlockNumber:     big.NewInt(10_000_835),
			Type:            2,
			ShouldFetch:     true,
		},
		{
			ContractAddress: common.HexToAddress(`0x1F98431c8aD98523631AE4a59f267346ea31F984`),
			BlockNumber:     big.NewInt(12_369_621),
			Type:            3,
			ShouldFetch:     true,
		},
	},
	10: {
		{
			ContractAddress: common.HexToAddress(`0x1F98431c8aD98523631AE4a59f267346ea31F984`),
			BlockNumber:     big.NewInt(100_000),
			Type:            3,
			ShouldFetch:     true,
		},
	},
	137: {
		{
			ContractAddress: common.HexToAddress(`0x1F98431c8aD98523631AE4a59f267346ea31F984`),
			BlockNumber:     big.NewInt(22_757_547),
			Type:            3,
			ShouldFetch:     true,
		},
	},
	42161: {
		{
			ContractAddress: common.HexToAddress(`0x1F98431c8aD98523631AE4a59f267346ea31F984`),
			BlockNumber:     big.NewInt(165),
			Type:            3,
			ShouldFetch:     true,
		},
	},
}

func handleUniswapPairsTokenList(tokensPerChainID map[uint64][]common.Address, allV2Pairs map[string]string) []TokenListToken {
	tokensForChainIDSyncMap := initSyncMap(tokensPerChainID)

	// Fetch the basic informations for all the tokens for all the chains
	perChainWG := sync.WaitGroup{}
	perChainWG.Add(len(tokensPerChainID))
	for chainID, list := range tokensPerChainID {
		go func(chainID uint64, list []common.Address) {
			defer perChainWG.Done()
			syncMapRaw, _ := tokensForChainIDSyncMap.Load(chainID)
			syncMap := syncMapRaw.([]TokenListToken)

			chainIDStr := strconv.FormatUint(chainID, 10)
			tokensInfo := ethereum.FetchBasicInformations(chainID, list)
			for _, address := range list {
				if token, ok := tokensInfo[address.Hex()]; ok {
					if token.Name == `` || token.Symbol == `` {
						continue
					}

					//We need to check if the token is a pair. If so we will overwrite the name and symbol
					name := token.Name
					symbol := token.Symbol
					if tokensOfThisPair, ok := allV2Pairs[chainIDStr+`_`+address.Hex()]; ok {
						splittedTokens := strings.Split(tokensOfThisPair, `_`)
						if len(splittedTokens) == 2 {
							token1 := splittedTokens[0]
							token2 := splittedTokens[1]
							if tokensInfo[token1] != nil && tokensInfo[token2] == nil {
								token1 := splittedTokens[0]
								token2 := splittedTokens[1]
								name = `Uniswap V2 ` + tokensInfo[token1].Name + ` + ` + tokensInfo[token2].Name
								symbol = `UNI-V2 ` + tokensInfo[token1].Symbol + ` + ` + tokensInfo[token2].Symbol
							}
						}
					}

					if newToken, err := SetToken(
						token.Address,
						name,
						symbol,
						``,
						chainID,
						int(token.Decimals),
					); err == nil {
						tokensForChainIDSyncMap.Store(chainID, append(syncMap, newToken))
					}
				}
			}
		}(chainID, list)
	}
	perChainWG.Wait()

	return extractSyncMap(tokensForChainIDSyncMap)
}

func fetchUniswapPairsTokenList(extra map[string]interface{}) ([]TokenListToken, map[uint64]string) {
	tokensPerChainID := make(map[uint64][]common.Address)

	allTokens := make(map[string]uint64)
	allV2Pairs := make(map[string]string)
	lastBlockSync := make(map[uint64]string)

	/**************************************************************************
	** Looping through all the Uniswap contracts per chainID to read the logs
	** and see the pairs and tokens that are being used.
	** In order to be included, a PAIR must have tokens that are both in at
	** least 10 different pairs.
	**************************************************************************/
	for chainID, uniContract := range UniswapContractsPerChainID {
		if helpers.IsChainIDIgnored(chainID) {
			continue
		}
		tokensPerChainID[chainID] = []common.Address{}

		/**********************************************************************
		** Init the RPC and get the current block number to know where to start
		** and end the logs fetching
		**********************************************************************/
		chainIDStr := strconv.FormatUint(chainID, 10)
		lastBlockSyncForChainID := uint64(0)
		if sync, ok := extra[`lastBlockSyncFor_`+chainIDStr]; ok {
			lastBlockSyncForChainID, _ = strconv.ParseUint(sync.(string), 10, 64)
		}
		client := ethereum.GetRPC(chainID)
		currentBlockNumber, _ := client.BlockNumber(context.Background())
		threshold := uint64(200_000)
		count := 0

		for _, uniContract := range uniContract {
			if !uniContract.ShouldFetch {
				continue
			}

			start := uniContract.BlockNumber.Uint64()
			if (lastBlockSyncForChainID != 0) && (lastBlockSyncForChainID > start) {
				start = lastBlockSyncForChainID
			}
			if uniContract.Type == 2 {
				uniV2Factory, _ := contracts.NewUniV2Factory(uniContract.ContractAddress, client)
				for startBlockToTest := start; startBlockToTest <= currentBlockNumber; startBlockToTest += threshold {
					end := startBlockToTest + threshold
					if end > currentBlockNumber {
						end = currentBlockNumber
					}
					options := &bind.FilterOpts{
						Start:   startBlockToTest,
						End:     &end,
						Context: nil,
					}
					logs.Info(`v2 - start: `, startBlockToTest, ` end: `, end, ` current count: `, count, ` total: `, len(allTokens), ` current block: `, currentBlockNumber, ` chainID: `, chainIDStr)
					if log, err := uniV2Factory.FilterPairCreated(options, nil, nil); err == nil {
						for log.Next() {
							if log.Error() != nil {
								continue
							}
							if _, ok := allTokens[log.Event.Token0.Hex()]; !ok {
								allTokens[log.Event.Token0.Hex()] = 0
							}
							if _, ok := allTokens[log.Event.Token1.Hex()]; !ok {
								allTokens[log.Event.Token1.Hex()] = 0
							}
							if _, ok := allTokens[log.Event.Pair.Hex()]; !ok {
								allTokens[log.Event.Pair.Hex()] = 0
							}
							allTokens[log.Event.Token0.Hex()]++
							allTokens[log.Event.Token1.Hex()]++
							allV2Pairs[chainIDStr+`_`+log.Event.Pair.Hex()] = log.Event.Token0.Hex() + `_` + log.Event.Token1.Hex()
							count += 2
						}
					} else {
						logs.Error("Error fetching all tokens from uniswap factory contract: ", err)
					}
				}
			} else if uniContract.Type == 3 {
				uniV3Factory, _ := contracts.NewUniV3Factory(uniContract.ContractAddress, client)
				for startBlockToTest := start; startBlockToTest <= currentBlockNumber; startBlockToTest += threshold {
					end := startBlockToTest + threshold
					if end > currentBlockNumber {
						end = currentBlockNumber
					}
					options := &bind.FilterOpts{
						Start:   startBlockToTest,
						End:     &end,
						Context: nil,
					}
					logs.Info(`v3 - start: `, startBlockToTest, ` end: `, end, ` current count: `, count, ` total: `, len(allTokens), ` current block: `, currentBlockNumber, ` chainID: `, chainIDStr)
					if log, err := uniV3Factory.FilterPoolCreated(options, nil, nil, nil); err == nil {
						for log.Next() {
							if log.Error() != nil {
								continue
							}
							if _, ok := allTokens[log.Event.Token0.Hex()]; !ok {
								allTokens[log.Event.Token0.Hex()] = 0
							}
							if _, ok := allTokens[log.Event.Token1.Hex()]; !ok {
								allTokens[log.Event.Token1.Hex()] = 0
							}
							allTokens[log.Event.Token0.Hex()]++
							allTokens[log.Event.Token1.Hex()]++
							count += 2
						}
					} else {
						logs.Error("Error fetching all tokens from uniswap factory contract: ", err)
					}
				}
			}
			lastBlockSync[chainID] = strconv.FormatUint(currentBlockNumber, 10)
		}
		/**********************************************************************
		** Adding the pairs that have at least 10 tokens in common
		**********************************************************************/
		for pair, tokensInPair := range allV2Pairs {
			tokens := strings.Split(tokensInPair, `_`)
			if (allTokens[tokens[0]] >= 10) && (allTokens[tokens[1]] >= 10) {
				chainAndPair := strings.Split(pair, `_`)
				allTokens[chainAndPair[1]] = 10
			}
		}

		/**********************************************************************
		** Transforming the output to the format that we need for the handle
		** function
		**********************************************************************/
		for address := range allTokens {
			if helpers.IsIgnoredToken(chainID, common.HexToAddress(address)) {
				continue
			}
			tokensPerChainID[chainID] = append(tokensPerChainID[chainID], common.HexToAddress(address))
		}
	}

	return handleUniswapPairsTokenList(tokensPerChainID, allV2Pairs), lastBlockSync
}

func buildUniswapPairsTokenList() {
	tokenList := loadTokenListFromJsonFile(`uniswap-pairs.json`)
	tokenList.Name = "Uniswap Token Pairs"
	tokenList.LogoURI = "ipfs://QmNa8mQkrNKp1WEEeGjFezDmDeodkWRevGFN8JCV7b4Xir"

	tokens, lastBlockSync := fetchUniswapPairsTokenList(tokenList.Extra)
	if tokenList.Extra == nil {
		tokenList.Extra = make(map[string]interface{})
	}
	for chainID, blockNumber := range lastBlockSync {
		chainIDStr := strconv.FormatUint(chainID, 10)
		tokenList.Extra[`lastBlockSyncFor_`+chainIDStr] = blockNumber
	}

	saveTokenListInJsonFile(tokenList, tokens, `uniswap-pairs.json`, Append)
}
