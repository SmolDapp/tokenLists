package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/contracts"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
	"github.com/migratooor/tokenLists/generators/common/models"
)

func handleAjnaTokenList(chainID uint64, tokenAddresses []common.Address) []models.TokenListToken {
	tokenList := helpers.GetTokensFromAddresses(chainID, tokenAddresses)
	tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
	return tokenList
}

func fetchAjnaTokenList(chainID uint64, sugarAddress common.Address) []models.TokenListToken {
	client := ethereum.GetRPC(chainID)
	ajnaPoolFactory, err := contracts.NewAjnaPoolFactoryCaller(sugarAddress, client)
	if err != nil {
		logs.Error(err)
		return []models.TokenListToken{}
	}
	/**************************************************************************
	** We first fetch all the pools deployed on Ajna. This will allow us to
	** retrieve the collateral and quote tokens for each pool in a second
	** step.
	**************************************************************************/
	allPools, err := ajnaPoolFactory.GetDeployedPoolsList(nil)
	if err != nil {
		logs.Error(err)
		return []models.TokenListToken{}
	}

	/**************************************************************************
	** We have all the pools, now we need to retrive what's going on in each
	** pool. We need to retrieve the collateralAddress and the
	** quoteTokenAddress
	**************************************************************************/
	addressesMap := make(map[common.Address]bool)
	for _, pool := range allPools {
		ajnaPool, err := contracts.NewAjnaPoolCaller(pool, client)
		if err != nil {
			logs.Error(err)
			continue
		}
		collateralAddress, errCollateral := ajnaPool.CollateralAddress(nil)
		if errCollateral == nil {
			addressesMap[collateralAddress] = true
		}
		quoteTokenAddress, errQuoteToken := ajnaPool.QuoteTokenAddress(nil)
		if errQuoteToken == nil {
			addressesMap[quoteTokenAddress] = true
		}
	}

	/**************************************************************************
	** We are supposed to get everything, so we can now remove the duplicates
	** and build the token list.
	**************************************************************************/
	addressesSlice := []common.Address{}
	for address := range addressesMap {
		addressesSlice = append(addressesSlice, address)
	}
	return handleAjnaTokenList(chainID, addressesSlice)
}

func buildAjnaTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`ajna.json`)
	tokenList.Name = `Ajna`
	tokenList.LogoURI = `https://www.ajna.finance/static/tokens/ajna.png`
	tokenList.Keywords = []string{`Ajna`}
	tokens := []models.TokenListToken{}
	tokens = append(tokens, fetchAjnaTokenList(1, common.HexToAddress(`0x6146DD43C5622bB6D12A5240ab9CF4de14eDC625`))...)
	tokens = append(tokens, fetchAjnaTokenList(5, common.HexToAddress(`0xDB61f8aD0B3ed0c5522b8FE71b80023fe9188e9e`))...)
	tokens = append(tokens, fetchAjnaTokenList(10, common.HexToAddress(`0x609C4e8804fafC07c96bE81A8a98d0AdCf2b7Dfa`))...)
	tokens = append(tokens, fetchAjnaTokenList(100, common.HexToAddress(`0x87578E357358163FCAb1711c62AcDB5BBFa1C9ef`))...)
	tokens = append(tokens, fetchAjnaTokenList(137, common.HexToAddress(`0x1f172F881eBa06Aa7a991651780527C173783Cf6`))...)
	tokens = append(tokens, fetchAjnaTokenList(8453, common.HexToAddress(`0x214f62B5836D83f3D6c4f71F174209097B1A779C`))...)
	tokens = append(tokens, fetchAjnaTokenList(42161, common.HexToAddress(`0xA3A1e968Bd6C578205E11256c8e6929f21742aAF`))...)

	helpers.SaveTokenListInJsonFile(tokenList, tokens, `ajna.json`, helpers.SavingMethodStandard)
}
