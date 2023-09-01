package main

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/contracts"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

func handleVeloTokenList(chainID uint64, tokenAddresses []common.Address) []TokenListToken {
	tokenList := []TokenListToken{}

	tokensInfo := retrieveBasicInformations(chainID, tokenAddresses)
	for _, address := range tokenAddresses {
		if token, ok := tokensInfo[address.Hex()]; ok {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				`https://assets.smold.app/api/token/`+strconv.FormatUint(chainID, 10)+`/`+token.Address.Hex()+`/logo-128.png`,
				chainID,
				int(token.Decimals),
			); err == nil {
				tokenList = append(tokenList, newToken)
			}
		}
	}
	chainCoin := ETHER
	chainCoin.ChainID = chainID
	tokenList = append(tokenList, chainCoin)
	return tokenList
}

func fetchVeloLikeTokenList(chainID uint64, sugarAddress common.Address) []TokenListToken {
	client := ethereum.GetRPC(chainID)
	veloSugar, err := contracts.NewVeloSugarV2Caller(sugarAddress, client)
	if err != nil {
		logs.Error(err)
		return []TokenListToken{}
	}
	allTokens, err := veloSugar.All(nil, big.NewInt(10_000), big.NewInt(0), common.Address{})
	if err != nil {
		logs.Error(err)
		return []TokenListToken{}
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
	tokenList := loadTokenListFromJsonFile(`velodrome.json`)
	tokenList.Name = `Velodrome`
	tokenList.LogoURI = `https://velodrome.finance/velodrome.svg`
	tokenList.Keywords = []string{`velodrome`, `optimism`}
	tokens := []TokenListToken{}
	tokens = append(tokens, fetchVeloLikeTokenList(10, common.HexToAddress(`0x7F45F1eA57E9231f846B2b4f5F8138F94295A726`))...)
	saveTokenListInJsonFile(tokenList, tokens, `velodrome.json`, Standard)
}
