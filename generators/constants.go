package main

import "github.com/ethereum/go-ethereum/common"

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
var XDAI = TokenListToken{
	Address:  `0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`,
	Name:     `xDai`,
	Symbol:   `xDAI`,
	LogoURI:  ``,
	ChainID:  100,
	Decimals: 18,
}
