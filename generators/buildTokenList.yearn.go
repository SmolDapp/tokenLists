package main

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

type TExternalERC20Token struct {
	Address                   string   `json:"address"`
	UnderlyingTokensAddresses []string `json:"underlyingTokensAddresses"`
	Name                      string   `json:"name"`
	Symbol                    string   `json:"symbol"`
	Type                      string   `json:"type"`
	DisplayName               string   `json:"display_name"`
	DisplaySymbol             string   `json:"display_symbol"`
	Description               string   `json:"description"`
	Icon                      string   `json:"icon"`
	Decimals                  uint64   `json:"decimals"`
}
type TYearnTokenData struct {
	Address       string              `json:"address"`
	Name          string              `json:"name"`
	Symbol        string              `json:"symbol"`
	DisplayName   string              `json:"display_name"`
	DisplaySymbol string              `json:"display_symbol"`
	Token         TExternalERC20Token `json:"token"`
	Decimals      uint64              `json:"decimals"`
	ChainID       uint64              `json:"chainID"`
}

func fetchYearnTokenList() []TokenListToken {
	tokens := []TokenListToken{}
	list := helpers.FetchJSON[[]TYearnTokenData](`https://ydaemon.yearn.fi/vaults/all?limit=20000&migrable=ignore`)

	for _, vault := range list {
		chainIDStr := strconv.FormatUint(vault.ChainID, 10)

		/**************************************************************************
		** The API from Yearn returns a list of tokens for each chainID. For each
		** token, we can use the IsVault flag to know if it's a vault or not.
		** For each vault, we need to add the underlying tokens to the list of
		** tokens to add to the token list. As the data is not directly available
		** with the vault information, we can fetch the token details after.
		**************************************************************************/
		if newToken, err := SetToken(
			common.HexToAddress(vault.Address),
			helpers.SafeString(vault.DisplayName, vault.Name),
			helpers.SafeString(vault.DisplaySymbol, vault.Symbol),
			`https://assets.smold.app/api/token/`+chainIDStr+`/`+vault.Address+`/logo-128.png`,
			vault.ChainID,
			int(vault.Decimals),
		); err == nil {
			tokens = append(tokens, newToken)
		}

		if newToken, err := SetToken(
			common.HexToAddress(vault.Token.Address),
			helpers.SafeString(vault.Token.DisplayName, vault.Token.Name),
			helpers.SafeString(vault.Token.DisplaySymbol, vault.Token.Symbol),
			`https://assets.smold.app/api/token/`+chainIDStr+`/`+vault.Token.Address+`/logo-128.png`,
			vault.ChainID,
			int(vault.Token.Decimals),
		); err == nil {
			tokens = append(tokens, newToken)
		}
	}

	return tokens
}

func buildYearnTokenList() {
	tokenList := loadTokenListFromJsonFile(`yearn.json`)
	tokenList.Name = `Yearn Token List`
	tokenList.LogoURI = `https://assets.smold.app/api/token/1/0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e/logo.svg`
	tokenList.Keywords = []string{`yearn`, `yfi`, `yvault`, `ytoken`, `ycurve`, `yprotocol`, `vaults`}
	tokens := fetchYearnTokenList()

	saveTokenListInJsonFile(tokenList, tokens, `yearn.json`, Standard)
}
