package main

import (
	"strconv"

	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

type TTokenType string

// TYearnTokenData contains the basic information of an ERC20 token
type TYearnTokenData struct {
	Address                   string     `json:"address"`
	UnderlyingTokensAddresses []string   `json:"underlyingTokens"`
	Type                      TTokenType `json:"type"`
	Name                      string     `json:"name"`
	Symbol                    string     `json:"symbol"`
	DisplayName               string     `json:"displayName"`
	DisplaySymbol             string     `json:"displaySymbol"`
	Description               string     `json:"description"`
	Category                  string     `json:"category"`
	Icon                      string     `json:"icon"`
	Decimals                  uint64     `json:"decimals"`
}

func fetchYearnTokenList() []models.TokenListToken {
	list := helpers.FetchJSON[map[uint64]map[string]TYearnTokenData](`https://ydevmon.ycorpo.com/tokens/all`)
	listPerChainID := []models.TokenListToken{}

	for chainID, listPerChain := range list {
		for _, vault := range listPerChain {
			chainIDStr := strconv.FormatUint(chainID, 10)

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

func buildYearnTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`yearn.json`)
	tokenList.Name = `Yearn Token List`
	tokenList.LogoURI = `https://assets.smold.app/api/token/1/0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e/logo.svg`
	tokenList.Keywords = []string{`yearn`, `yfi`, `yvault`, `ytoken`, `ycurve`, `yprotocol`, `vaults`}
	tokens := fetchYearnTokenList()

	helpers.SaveTokenListInJsonFile(tokenList, tokens, `yearn.json`, helpers.SavingMethodStandard)
}
