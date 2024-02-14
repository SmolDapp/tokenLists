package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/contracts"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
	"github.com/migratooor/tokenLists/generators/common/models"
)

func handleVeloTokenList(chainID uint64, tokenAddresses []common.Address) []models.TokenListToken {
	tokenList := helpers.GetTokensFromAddresses(chainID, tokenAddresses)
	tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
	return tokenList
}

func fetchVeloLikeTokenList(chainID uint64, sugarAddress common.Address) []models.TokenListToken {
	client := ethereum.GetRPC(chainID)
	veloSugar, err := contracts.NewVeloSugarV2Caller(sugarAddress, client)
	if err != nil {
		logs.Error(err)
		return []models.TokenListToken{}
	}
	allTokens, err := veloSugar.All(nil, big.NewInt(10_000), big.NewInt(0), common.Address{})
	if err != nil {
		logs.Error(err)
		return []models.TokenListToken{}
	}
	addressesMap := make(map[common.Address]bool)
	for _, token := range allTokens {
		addressesMap[token.Token0] = true
		addressesMap[token.Token1] = true
		addressesMap[token.EmissionsToken] = true
		addressesMap[token.Lp] = true
	}
	// Used to remove the duplicates
	addressesSlice := []common.Address{}
	for address := range addressesMap {
		addressesSlice = append(addressesSlice, address)
	}
	return handleVeloTokenList(chainID, addressesSlice)
}

func buildVeloTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`velodrome.json`)
	tokenList.Name = `Velodrome`
	tokenList.LogoURI = `https://velodrome.finance/velodrome.svg`
	tokenList.Keywords = []string{`velodrome`, `optimism`}
	tokens := []models.TokenListToken{}
	tokens = append(tokens, fetchVeloLikeTokenList(10, common.HexToAddress(`0x7F45F1eA57E9231f846B2b4f5F8138F94295A726`))...)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `velodrome.json`, helpers.SavingMethodStandard)
}
