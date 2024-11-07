package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/contracts"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/common/utils"
)

func handleVeloTokenList(chainID uint64, tokenAddresses []string) []models.TokenListToken {
	tokenList := helpers.GetTokensFromAddresses(chainID, tokenAddresses)
	tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
	return tokenList
}

func fetchVeloLikeTokenList(chainID uint64, sugarAddress common.Address) []models.TokenListToken {
	client := ethereum.GetRPC(chainID).(*ethclient.Client)
	veloSugar, err := contracts.NewVeloSugarV2Caller(sugarAddress, client)
	if err != nil {
		logs.Error(err)
		return []models.TokenListToken{}
	}
	addressesMap := make(map[string]bool)
	addressesSlice := []string{}

	for i := 0; i < 100; i++ {
		offset := big.NewInt(int64(i * 25))
		allTokens, err := veloSugar.All(nil, big.NewInt(25), offset, common.Address{})
		if err != nil {
			logs.Error(err)
			break
		}
		for _, token := range allTokens {
			addressesMap[utils.ToAddress(token.Token0.Hex())] = true
			addressesMap[utils.ToAddress(token.Token1.Hex())] = true
			addressesMap[utils.ToAddress(token.EmissionsToken.Hex())] = true
			addressesMap[utils.ToAddress(token.Lp.Hex())] = true
		}
		// Used to remove the duplicates
		for address := range addressesMap {
			addressesSlice = append(addressesSlice, address)
		}
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
