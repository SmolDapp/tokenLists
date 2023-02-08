package main

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

type TYearnTokenData struct {
	Address          string   `json:"address"`
	Name             string   `json:"name"`
	Symbol           string   `json:"symbol"`
	DisplayName      string   `json:"display_name"`
	DisplaySymbol    string   `json:"display_symbol"`
	Description      string   `json:"description"`
	Website          string   `json:"website"`
	Categories       []string `json:"categories"`
	UnderlyingTokens []string `json:"underlyingTokens"`
	Decimals         uint64   `json:"decimals"`
	IsVault          bool     `json:"isVault"`
}

func fetchYearnTokenList() []TokenListToken {
	tokens := []TokenListToken{}
	list := helpers.FetchJSON[map[string]map[string]TYearnTokenData](`https://ydaemon.yearn.finance/tokens/all`)

	for chainIDStr, tokensList := range list {
		chainID, _ := strconv.ParseUint(chainIDStr, 10, 64)
		vaultsUnderlyingAddress := []common.Address{}

		/**************************************************************************
		** The API from Yearn returns a list of tokens for each chainID. For each
		** token, we can use the IsVault flag to know if it's a vault or not.
		** For each vault, we need to add the underlying tokens to the list of
		** tokens to add to the token list. As the data is not directly available
		** with the vault information, we can fetch the token details after.
		**************************************************************************/
		for _, token := range tokensList {
			if !token.IsVault {
				continue
			}
			if newToken, err := SetToken(
				common.HexToAddress(token.Address),
				helpers.SafeString(token.DisplayName, token.Name),
				helpers.SafeString(token.DisplaySymbol, token.Symbol),
				`https://raw.githubusercontent.com/yearn/yearn-assets/master/icons/multichain-tokens/`+chainIDStr+`/`+token.Address+`/logo-128.png`,
				chainID,
				int(token.Decimals),
			); err == nil {
				tokens = append(tokens, newToken)
			}
			for _, underlyingToken := range token.UnderlyingTokens {
				vaultsUnderlyingAddress = append(vaultsUnderlyingAddress, common.HexToAddress(underlyingToken))
			}
		}

		/**************************************************************************
		** For each underlying token for this chain, we fetch the basic information
		** and add it to the list of tokens to add to the token list.
		** This fetch could be bypassed by checking the list returned by the API
		** and checking if the token is in the list.
		**************************************************************************/
		underlyingTokensInfo := ethereum.FetchBasicInformations(chainID, vaultsUnderlyingAddress)
		for _, token := range underlyingTokensInfo {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				`https://raw.githubusercontent.com/yearn/yearn-assets/master/icons/multichain-tokens/`+chainIDStr+`/`+token.Address.Hex()+`/logo-128.png`,
				chainID,
				int(token.Decimals),
			); err == nil {
				tokens = append(tokens, newToken)
			}
		}

	}
	return tokens
}

func buildYearnTokenList() {
	tokenList := loadTokenListFromJsonFile(`yearn.json`)
	tokenList.Name = `Yearn Token List`
	tokenList.LogoURI = `https://raw.githubusercontent.com/yearn/yearn-assets/master/icons/multichain-tokens/1/0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e/logo.svg`
	tokenList.Keywords = []string{`yearn`, `yfi`, `yvault`, `ytoken`, `ycurve`, `yprotocol`, `vaults`}
	tokens := fetchYearnTokenList()

	saveTokenListInJsonFile(tokenList, tokens, `yearn.json`, Standard)
}
