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
)

func handleRegistryToken(chainID uint64, tokenAddresses []string, icons map[string]string) []models.TokenListToken {
	tokenList := helpers.GetTokensFromAddressesWithIcons_NOREPLACEMENT(chainID, tokenAddresses, icons)
	tokenList = append(tokenList, chains.CHAINS[chainID].Coin)
	return tokenList
}

func fetchTokenRegistryTokens(chainID uint64, registryAddress common.Address) []models.TokenListToken {
	client := ethereum.GetRPC(chainID).(*ethclient.Client)
	registry, err := contracts.NewTokenRegistryCaller(registryAddress, client)

	if err != nil {
		logs.Error(err)
		return []models.TokenListToken{}
	}
	addressesSlice := []string{}
	icons := make(map[string]string)
	limit := big.NewInt(100)
	for i := 0; i < 100; i++ {
		offset := big.NewInt(int64(i * 100))
		allTokens, err := registry.ListTokens(nil, offset, limit, 1, true)
		if err != nil || len(allTokens.Tokens) == 0 {
			logs.Error(err)
			break
		}
		for _, token := range allTokens.Tokens {
			addressesSlice = append(addressesSlice, token.ContractAddress.Hex())
			icons[token.ContractAddress.Hex()] = token.LogoURI
		}
	}
	return handleRegistryToken(chainID, addressesSlice, icons)
}

func buildTokenRegistryTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`tokenRegistry.json`)
	tokenList.Name = `Token Registry`
	tokenList.LogoURI = `https://tokenregistry.builtby.mom/favicons/web-app-manifest-512x512.png`
	tokenList.Keywords = []string{`tokenRegistry`}
	tokens := []models.TokenListToken{}
	tokens = append(tokens, fetchTokenRegistryTokens(10, common.HexToAddress(`0x1f393eCAadbDCC3016AF9671450E4486B53FfE19`))...)
	tokens = append(tokens, fetchTokenRegistryTokens(8453, common.HexToAddress(`0x1f393eCAadbDCC3016AF9671450E4486B53FfE19`))...)

	helpers.SaveTokenListInJsonFile(tokenList, tokens, `tokenRegistry.json`, helpers.SavingMethodStandard_NOREPLACEMENT)
}
