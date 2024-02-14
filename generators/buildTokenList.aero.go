package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

func buildAeroTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`aerodrome.json`)
	tokenList.Name = `Aerodrome`
	tokenList.LogoURI = `https://aerodrome.finance/aerodrome.svg`
	tokenList.Keywords = []string{`aerodrome`, `base`, `velodrome`}
	tokens := []models.TokenListToken{}
	tokens = append(tokens, fetchVeloLikeTokenList(8453, common.HexToAddress(`0x2073d8035bb2b0f2e85aaf5a8732c6f397f9ff9b`))...)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `aerodrome.json`, helpers.SavingMethodStandard)
}
