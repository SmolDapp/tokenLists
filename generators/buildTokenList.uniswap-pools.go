package main

import (
	"context"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/contracts"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
	"github.com/migratooor/tokenLists/generators/common/models"
)

var UNI_POOL_THRESHOLD_FOR_CHAINID = map[uint64]int{
	1:     10,
	10:    3,
	56:    3,
	100:   3,
	137:   3,
	250:   3,
	42161: 3,
	43114: 3,
}

func handleUniswapPoolsTokenList(tokensPerChainID map[uint64][]common.Address, allPoolsPerChainID map[uint64]map[string]string) []models.TokenListToken {
	tokensForChainIDSyncMap := helpers.InitSyncMap(tokensPerChainID)

	// Fetch the basic informations for all the tokens for all the chains
	perChainWG := sync.WaitGroup{}
	perChainWG.Add(len(tokensPerChainID))
	for chainID, list := range tokensPerChainID {
		go func(chainID uint64, list []common.Address) {
			defer perChainWG.Done()
			syncMapRaw, _ := tokensForChainIDSyncMap.Load(chainID)
			syncMap := syncMapRaw.([]models.TokenListToken)

			/**************************************************************************
			** The pairs have no name or symbol to recognize them. We need to fetch the
			** underlying tokens and use their name and symbol to build the pair name.
			** The first step is to fetch the data for all the underlying tokens.
			**************************************************************************/
			underlyingTokenInfo := helpers.RetrieveBasicInformations(chainID, list)

			/**************************************************************************
			** Once we have the data for all the underlying tokens, we can loop over
			** the pairs and build the name and symbol.
			**************************************************************************/
			for pool, tokens := range allPoolsPerChainID[chainID] {
				splittedTokens := strings.Split(tokens, `_`)
				if len(splittedTokens) != 2 {
					continue
				}

				token1, ok1 := underlyingTokenInfo[helpers.ToAddress(splittedTokens[0])]
				token2, ok2 := underlyingTokenInfo[helpers.ToAddress(splittedTokens[1])]
				if !ok1 || !ok2 {
					continue
				}

				if newToken, err := helpers.SetToken(
					common.HexToAddress(pool),
					`Uniswap V2 `+token1.Name+` + `+token2.Name,
					`UNI-V2 `+token1.Symbol+` + `+token2.Symbol,
					``,
					chainID,
					18,
				); err == nil {
					syncMap = append(syncMap, newToken)
					tokensForChainIDSyncMap.Store(chainID, syncMap)
				}
			}
		}(chainID, list)
	}
	perChainWG.Wait()

	return helpers.ExtractSyncMap(tokensForChainIDSyncMap)
}

func fetchUniswapPoolsTokenList(extra map[string]interface{}) ([]models.TokenListToken, map[uint64]string) {
	tokensPerChainID := make(map[uint64][]common.Address)
	poolsPerChainID := make(map[uint64]map[string]string)
	allTokens := make(map[string]int)
	allPools := make(map[string]string)
	lastBlockSync := make(map[uint64]string)

	/**************************************************************************
	** Looping through all the Uniswap contracts per chainID to read the logs
	** and see the pairs and tokens that are being used.
	** In order to be included, a PAIR must have tokens that are both in at
	** least 3 different pairs.
	**************************************************************************/
	for chainID, uniContract := range UniswapContractsPerChainID {
		if !chains.IsChainIDSupported(chainID) {
			continue
		}
		tokensPerChainID[chainID] = []common.Address{}
		poolsPerChainID[chainID] = make(map[string]string)

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
		threshold := uint64(100_000)
		if chainID == 56 {
			threshold = uint64(5_000)
		}

		/**********************************************************************
		** For each registered Uniswap contract, we will fetch the logs to
		** get the pairs and tokens
		**********************************************************************/
		for _, uniContract := range uniContract {
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
					logs.Info(`v2 - start: `, startBlockToTest, ` end: `, end, ` total: `, len(allTokens), ` current block: `, currentBlockNumber, ` chainID: `, chainIDStr)
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
							allTokens[log.Event.Token0.Hex()]++
							allTokens[log.Event.Token1.Hex()]++
							allPools[log.Event.Pair.Hex()] = log.Event.Token0.Hex() + `_` + log.Event.Token1.Hex()
						}
					} else {
						logs.Error("Error fetching all tokens from uniswap factory contract: ", err)
						startBlockToTest -= threshold
					}
				}
			}
			lastBlockSync[chainID] = strconv.FormatUint(currentBlockNumber, 10)
		}
		/**********************************************************************
		** Adding the pairs that have at least UNI_POOL_THRESHOLD tokens in
		** common
		**********************************************************************/
		chainThreshold := UNI_POOL_THRESHOLD_FOR_CHAINID[chainID]
		for pool, tokensInPool := range allPools {
			tokens := strings.Split(tokensInPool, `_`)
			if (allTokens[tokens[0]] >= chainThreshold) && (allTokens[tokens[1]] >= chainThreshold) {
				if chains.IsTokenIgnored(chainID, common.HexToAddress(tokens[0])) ||
					chains.IsTokenIgnored(chainID, common.HexToAddress(tokens[1])) {
					continue
				}
				tokensPerChainID[chainID] = append(tokensPerChainID[chainID], common.HexToAddress(tokens[0]))
				tokensPerChainID[chainID] = append(tokensPerChainID[chainID], common.HexToAddress(tokens[1]))
				poolsPerChainID[chainID][pool] = tokensInPool
			}
		}
	}

	return handleUniswapPoolsTokenList(tokensPerChainID, poolsPerChainID), lastBlockSync
}

func buildUniswapPoolsTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`uniswap-pools.json`)
	tokenList.Name = "Uniswap Token Pools"
	tokenList.LogoURI = "ipfs://QmNa8mQkrNKp1WEEeGjFezDmDeodkWRevGFN8JCV7b4Xir"

	tokens, lastBlockSync := fetchUniswapPoolsTokenList(tokenList.Metadata)
	if tokenList.Metadata == nil {
		tokenList.Metadata = make(map[string]interface{})
	}
	for chainID, blockNumber := range lastBlockSync {
		chainIDStr := strconv.FormatUint(chainID, 10)
		tokenList.Metadata[`lastBlockSyncFor_`+chainIDStr] = blockNumber
	}

	helpers.SaveTokenListInJsonFile(tokenList, tokens, `uniswap-pools.json`, helpers.SavingMethodAppend)
}
