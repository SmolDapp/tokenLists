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

type TSushiContracts struct {
	ContractAddress common.Address
	BlockNumber     *big.Int
}

var SushiswapContractsPerChainID = map[uint64][]TSushiContracts{
	1: {
		{
			ContractAddress: common.HexToAddress(`0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac`),
			BlockNumber:     big.NewInt(10_794_229),
		},
	},
	56: {
		{
			ContractAddress: common.HexToAddress(`0xc35DADB65012eC5796536bD9864eD8773aBc74C4`),
			BlockNumber:     big.NewInt(5_205_069),
		},
	},
	100: {
		{
			ContractAddress: common.HexToAddress(`0xc35DADB65012eC5796536bD9864eD8773aBc74C4`),
			BlockNumber:     big.NewInt(14_735_904),
		},
	},
	137: {
		{
			ContractAddress: common.HexToAddress(`0xc35DADB65012eC5796536bD9864eD8773aBc74C4`),
			BlockNumber:     big.NewInt(11_333_218),
		},
	},
	250: {
		{
			ContractAddress: common.HexToAddress(`0xc35DADB65012eC5796536bD9864eD8773aBc74C4`),
			BlockNumber:     big.NewInt(2_457_879),
		},
	},
	42161: {
		{
			ContractAddress: common.HexToAddress(`0xc35DADB65012eC5796536bD9864eD8773aBc74C4`),
			BlockNumber:     big.NewInt(70),
		},
	},
	43114: {
		{
			ContractAddress: common.HexToAddress(`0xc35DADB65012eC5796536bD9864eD8773aBc74C4`),
			BlockNumber:     big.NewInt(506_190),
		},
	},
}

func handleSushiswapPairsTokenList(tokensPerChainID map[uint64][]common.Address, allV2Pairs map[string]string) []TokenListToken {
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
								name = `Sushiswap V2 ` + tokensInfo[token1].Name + ` + ` + tokensInfo[token2].Name
								symbol = `SUSHI-V2 ` + tokensInfo[token1].Symbol + ` + ` + tokensInfo[token2].Symbol
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

func fetchSushiswapPairsTokenList(extra map[string]interface{}) ([]TokenListToken, map[uint64]string) {
	tokensPerChainID := make(map[uint64][]common.Address)

	allTokens := make(map[string]uint64)
	allV2Pairs := make(map[string]string)
	lastBlockSync := make(map[uint64]string)

	/**************************************************************************
	** Looping through all the Sushiswap contracts per chainID to read the logs
	** and see the pairs and tokens that are being used.
	** In order to be included, a PAIR must have tokens that are both in at
	** least 10 different pairs.
	**************************************************************************/
	for chainID, sushiContract := range SushiswapContractsPerChainID {
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
		if chainID == 56 {
			threshold = uint64(5_000)
		}
		count := 0

		for _, sushiContract := range sushiContract {
			start := sushiContract.BlockNumber.Uint64()
			if (lastBlockSyncForChainID != 0) && (lastBlockSyncForChainID > start) {
				start = lastBlockSyncForChainID
			}
			sushiV2Factory, _ := contracts.NewSushiV2Factory(sushiContract.ContractAddress, client)
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
				if log, err := sushiV2Factory.FilterPairCreated(options, nil, nil); err == nil {
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
					logs.Error("Error fetching all tokens from sushiswap factory contract: ", err)
					startBlockToTest -= threshold
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
		for address, count := range allTokens {
			if count < 10 {
				continue
			}
			if helpers.IsIgnoredToken(chainID, common.HexToAddress(address)) {
				continue
			}
			tokensPerChainID[chainID] = append(tokensPerChainID[chainID], common.HexToAddress(address))
		}
	}

	return handleSushiswapPairsTokenList(tokensPerChainID, allV2Pairs), lastBlockSync
}

func buildSushiswapPairsTokenList() {
	tokenList := loadTokenListFromJsonFile(`sushiswap-pairs.json`)
	tokenList.Name = "SushiSwap Token Pairs"
	tokenList.LogoURI = "https://raw.githubusercontent.com/sushiswap/art/master/sushi/logo-256x256.png"

	tokens, lastBlockSync := fetchSushiswapPairsTokenList(tokenList.Extra)
	if tokenList.Extra == nil {
		tokenList.Extra = make(map[string]interface{})
	}
	for chainID, blockNumber := range lastBlockSync {
		chainIDStr := strconv.FormatUint(chainID, 10)
		tokenList.Extra[`lastBlockSyncFor_`+chainIDStr] = blockNumber
	}

	saveTokenListInJsonFile(tokenList, tokens, `sushiswap-pairs.json`, Append)
}
