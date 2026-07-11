package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
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
	tokenAddresses := []string{}
	tokenIcons := map[string]string{}

	list := helpers.FetchJSONPost[TFilescanList](explorerBaseUri)

	for _, token := range list.Result.Tokens {
		if token.MarketCap == `0` || token.MarketCap == `` {
			continue
		}
		tokenAddresses = append(tokenAddresses, token.Address)
		tokenIcons[common.HexToAddress(token.Address).Hex()] = token.LogoURI
	}
	return helpers.GetTokensFromAddressesWithIcons(chainID, tokenAddresses, tokenIcons)
}

func buildFilecoinTokenList() {
	tokenList := helpers.LoadTokenListFromJsonFile(`filecoin.json`)
	tokenList.Name = `filecoin`
	tokenList.LogoURI = chains.CHAINS[314].LogoURI
	tokenList.Keywords = []string{`filecoin`}
	tokens := []models.TokenListToken{}
	tokens = append(tokens, fetchFilscanTokens(314)...)

	// We are ignoring the blockscout token list for now as we have no way to filter out the tokens with 0 market cap

	helpers.SaveTokenListInJsonFile(tokenList, tokens, `filecoin.json`, helpers.SavingMethodStandard)
}
