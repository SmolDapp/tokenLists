package main

import (
	"github.com/ethereum/go-ethereum/common"
)

func buildAeroTokenList() {
	tokenList := loadTokenListFromJsonFile(`aerodrome.json`)
	tokenList.Name = `Aerodrome`
	tokenList.LogoURI = `https://aerodrome.finance/aerodrome.svg`
	tokenList.Keywords = []string{`aerodrome`, `base`, `velodrome`}
	tokens := []TokenListToken{}
	tokens = append(tokens, fetchVeloLikeTokenList(8453, common.HexToAddress(`0x2073d8035bb2b0f2e85aaf5a8732c6f397f9ff9b`))...)
	saveTokenListInJsonFile(tokenList, tokens, `aerodrome.json`, Standard)
}
