package main

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

func handleSmolAssetsTokenList(chainID uint64, tokenAddresses []common.Address) []TokenListToken {
	tokenList := []TokenListToken{}

	tokensInfo := retrieveBasicInformations(chainID, tokenAddresses)
	for _, address := range tokenAddresses {
		if token, ok := tokensInfo[address.Hex()]; ok {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				`https://assets.smold.app/api/token/`+strconv.FormatUint(chainID, 10)+`/`+address.Hex()+`/logo-128.png`,
				chainID,
				int(token.Decimals),
			); err == nil {
				tokenList = append(tokenList, newToken)
			}
		}
	}
	chainCoin := BASE_EXPLORERS_URI[chainID].Coin
	if chainCoin.Address == "" {
		chainCoin = ETHER
	}
	chainCoin.ChainID = chainID
	chainCoin.LogoURI = `https://assets.smold.app/api/token/` + strconv.FormatUint(chainID, 10) + `/` + chainCoin.Address + `/logo-128.png`
	tokenList = append(tokenList, chainCoin)
	return tokenList
}

func fetchSmolAssetsTokenList(chainID uint64) []TokenListToken {
	smolAssets := helpers.GetSmolAssetsPerChain(chainID)
	allTokensToAdd := []common.Address{}
	for _, token := range smolAssets {
		allTokensToAdd = append(allTokensToAdd, common.HexToAddress(token))
	}
	return handleSmolAssetsTokenList(chainID, allTokensToAdd)
}

func buildSmolAssetsTokenList() {
	tokenList := loadTokenListFromJsonFile(`smolAssets.json`)
	tokenList.Name = `SmolAssets`
	tokenList.Description = `A list of tokens supported by Smoldapp Token Assets repository`
	tokenList.LogoURI = `https://raw.githubusercontent.com/Migratooor/tokenLists/main/.github/tokenlistooor.svg`
	tokenList.Keywords = []string{`smol`, `tokenAssets`}
	tokens := []TokenListToken{}
	for chainID := range helpers.SUPPORTED_CHAIN_IDS {
		tokens = append(tokens, fetchSmolAssetsTokenList(chainID)...)
	}
	saveTokenListInJsonFile(tokenList, tokens, `smolAssets.json`, Standard)
}
