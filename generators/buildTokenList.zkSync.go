package main

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/static"
)

func handleZkSyncTokenList(chainID uint64, tokens []static.TStaticElement) []TokenListToken {
	tokenList := []TokenListToken{}
	tokenAddresses := []common.Address{}
	for _, token := range tokens {
		tokenAddresses = append(tokenAddresses, token.Address)
	}

	tokensInfo := retrieveBasicInformations(chainID, tokenAddresses)
	for index, address := range tokenAddresses {
		if token, ok := tokensInfo[address.Hex()]; ok {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				tokens[index].Icon,
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

func fetchZkSyncTokenList() []TokenListToken {
	tokens := static.ZKSYNC_STATIC_TOKENLIST
	return handleZkSyncTokenList(324, tokens[324])
}

func buildZkSyncTokenList() {
	tokenList := loadTokenListFromJsonFile(`zksync.json`)
	tokenList.Name = `zkSync`
	tokenList.LogoURI = `https://assets.smold.app/api/chain/324/logo-128.png`
	tokenList.Keywords = []string{`zksync`}
	tokens := []TokenListToken{}
	tokens = append(tokens, fetchZkSyncTokenList()...)
	saveTokenListInJsonFile(tokenList, tokens, `zksync.json`, Standard)
}
