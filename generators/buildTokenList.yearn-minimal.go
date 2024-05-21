package main

import (
	"strconv"

	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

func fetchYearnMinTokenList() []models.TokenListToken {
	list := helpers.FetchJSON[map[uint64]map[string]TYearnTokenData](`https://ydaemon.yearn.fi/tokens/all`)
	listPerChainID := []models.TokenListToken{}

	for chainID, listPerChain := range list {
		for _, vault := range listPerChain {
			chainIDStr := strconv.FormatUint(chainID, 10)

			isAjnaToken := vault.Symbol == `bwAJNA` || vault.Symbol == `AJNA`
			isdYFI := vault.Symbol == `dYFI`
			isOptimism := vault.Symbol == `OP`
			if vault.Category != `yVault` && !isAjnaToken && !isdYFI && !isOptimism {
				continue
			}

			/**************************************************************************
			** The API from Yearn returns a list of tokens for each chainID. For each
			** token, we can use the IsVault flag to know if it's a vault or not.
			** For each vault, we need to add the underlying tokens to the list of
			** tokens to add to the token list. As the data is not directly available
			** with the vault information, we can fetch the token details after.
			**************************************************************************/
			listPerChainID = append(listPerChainID, models.TokenListToken{
				Address: vault.Address,
				Name:    vault.Name,
				Symbol:  vault.Symbol,
				LogoURI: `https://assets.smold.app/api/token/` + chainIDStr + `/` + vault.Address + `/logo-128.png`,
				ChainID: chainID,
			})

			for _, underlyingTokenAddress := range vault.UnderlyingTokensAddresses {
				listPerChainID = append(listPerChainID, models.TokenListToken{
					Address: underlyingTokenAddress,
					ChainID: chainID,
				})
			}
		}
	}

	return helpers.GetTokensFromList(listPerChainID)
}

func buildYearnMinimalTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`yearn-min.json`)
	tokenList.Name = `Yearn Minimal Token List`
	tokenList.LogoURI = `https://assets.smold.app/api/token/1/0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e/logo.svg`
	tokenList.Keywords = []string{`yearn`, `yfi`, `yvault`, `ytoken`, `ycurve`, `yprotocol`, `vaults`}
	tokens := fetchYearnMinTokenList()

	helpers.SaveTokenListInJsonFile(tokenList, tokens, `yearn-min.json`, helpers.SavingMethodStandard)
}
