package main

import (
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/static"
)

func fetchSophonStaticTokenList(chainID uint64) []models.TokenListToken {
	tokens := static.SOPHON_STATIC_TOKENLIST
	return handleStaticTokenList(chainID, tokens[chainID])
}

func buildSophonStaticTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`sophon-static.json`)
	tokenList.Name = `Sophon (Static)`
	tokenList.Description = `Sophon aims to make the experiences people already love more valuable, more connected, and more rewarding`
	tokenList.LogoURI = `https://assets.smold.app/chains/50104/logo-128.png`
	tokenList.Keywords = []string{`Sophon`}
	tokens := []models.TokenListToken{}
	tokens = append(tokens, fetchSophonStaticTokenList(50104)...)
	helpers.SaveTokenListInJsonFile(tokenList, tokens, `sophon-static.json`, helpers.SavingMethodStandard)
}
