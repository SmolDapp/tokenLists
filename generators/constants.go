package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

var ALL_EXISTING_TOKENS = map[uint64]map[string]TokenListToken{}

func init() {
	chainCoins := []TokenListToken{}
	chainCoins = append(chainCoins, ETHER)
	chainCoins = append(chainCoins, FTM)
	chainCoins = append(chainCoins, BSC)
	chainCoins = append(chainCoins, MATIC)
	chainCoins = append(chainCoins, MATIC_ZKEVM)
	chainCoins = append(chainCoins, XDAI)
	chainCoins = append(chainCoins, addEtherLikeToken(10))
	chainCoins = append(chainCoins, addEtherLikeToken(8453))
	chainCoins = append(chainCoins, addEtherLikeToken(324))

	for _, coin := range chainCoins {
		if _, ok := ALL_EXISTING_TOKENS[coin.ChainID]; !ok {
			ALL_EXISTING_TOKENS[coin.ChainID] = make(map[string]TokenListToken)
		}
		if coin.Name == `` && coin.Symbol == `` {
			logs.Warning(`Missing informations for token:`, coin.Address, `on chain:`, coin.ChainID)
		}
		ALL_EXISTING_TOKENS[coin.ChainID][coin.Address] = TokenListToken{
			Address:    common.HexToAddress(coin.Address).Hex(),
			Name:       coin.Name,
			Symbol:     coin.Symbol,
			LogoURI:    ``,
			Decimals:   int(coin.Decimals),
			ChainID:    coin.ChainID,
			Occurrence: 1,
		}
	}
}

func addEtherToken(chainId uint64, tokenList []TokenListToken) []TokenListToken {
	if newToken, err := SetToken(
		common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
		`Ethereum`,
		`ETH`,
		``,
		chainId,
		18,
	); err == nil {
		tokenList = append(tokenList, newToken)
	}
	return tokenList
}

func addEtherLikeToken(chainID uint64) TokenListToken {
	return TokenListToken{
		Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
		Name:     `Ethereum`,
		Symbol:   `ETH`,
		LogoURI:  ``,
		ChainID:  chainID,
		Decimals: 18,
	}
}

var ETHER = TokenListToken{
	Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
	Name:     `Ethereum`,
	Symbol:   `ETH`,
	LogoURI:  ``,
	ChainID:  1,
	Decimals: 18,
}
var FTM = TokenListToken{
	Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
	Name:     `Fantom`,
	Symbol:   `FTM`,
	LogoURI:  ``,
	ChainID:  250,
	Decimals: 18,
}
var BSC = TokenListToken{
	Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
	Name:     `Binance Smart Chain`,
	Symbol:   `BNB`,
	LogoURI:  ``,
	ChainID:  56,
	Decimals: 18,
}
var MATIC = TokenListToken{
	Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
	Name:     `Matic`,
	Symbol:   `MATIC`,
	LogoURI:  ``,
	ChainID:  137,
	Decimals: 18,
}
var MATIC_ZKEVM = TokenListToken{
	Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
	Name:     `Matic`,
	Symbol:   `MATIC`,
	LogoURI:  ``,
	ChainID:  1101,
	Decimals: 18,
}
var XDAI = TokenListToken{
	Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
	Name:     `xDai`,
	Symbol:   `xDAI`,
	LogoURI:  ``,
	ChainID:  100,
	Decimals: 18,
}
