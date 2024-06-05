package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/models"
)

func fetchFilscanTokens(chainID uint64) []models.TokenListToken {
	type TFilescanTokenData struct {
		Address   string `json:"contract_id"`
		Name      string `json:"token_name"`
		LogoURI   string `json:"icon_url"`
		MarketCap string `json:"market_cap"`
	}
	type TFilescanList struct {
		Result struct {
			Tokens []TFilescanTokenData `json:"items"`
		} `json:"result"`
	}

	explorerBaseUri := `https://api-v2.filscan.io/api/v1/ERC20List`
	tokenAddresses := []common.Address{}
	tokenIcons := map[string]string{}

	uri := explorerBaseUri
	list := helpers.FetchJSONPost[TFilescanList](uri)

	for _, token := range list.Result.Tokens {
		if token.MarketCap == `0` || token.MarketCap == `` {
			continue
		}
		tokenAddresses = append(tokenAddresses, common.HexToAddress(token.Address))
		tokenIcons[common.HexToAddress(token.Address).Hex()] = token.LogoURI
	}
	return helpers.GetTokensFromAddressesWithIcons(chainID, tokenAddresses, tokenIcons)
}

func buildFilecoinTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`filecoin.json`)
	tokenList.Name = `filecoin`
	tokenList.LogoURI = `https://assets.smold.app/api/chains/314/logo-128.png`
	tokenList.Keywords = []string{`filecoin`}
	tokens := []models.TokenListToken{}
	tokens = append(tokens, fetchFilscanTokens(314)...)

	// We are ignoring the blockscout token list for now as we have no way to filter out the tokens with 0 market cap
	// tokens = append(tokens, fetchBlockScoutV6TokenList(314)...)

	helpers.SaveTokenListInJsonFile(tokenList, tokens, `filecoin.json`, helpers.SavingMethodStandard)
}
