package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

type TPortalTokenData struct {
	Key       string            `json:"key"`
	Name      string            `json:"name"`
	Decimals  int               `json:"decimals"`
	Symbol    string            `json:"symbol"`
	Price     float64           `json:"price"`
	Address   string            `json:"address"`
	Addresses map[string]string `json:"addresses"`
	Platform  string            `json:"platform"`
	Network   string            `json:"network"`
	Images    []string          `json:"images"`
	UpdatedAt string            `json:"updatedAt"`
	CreatedAt string            `json:"createdAt"`
	Tokens    []string          `json:"tokens"`
	Liquidity float64           `json:"liquidity"`
}
type TPortalList struct {
	PageItems  int
	TotalItems int
	More       bool
	Page       int
	Tokens     []TPortalTokenData
}

func portalsMapNetworkNameToChainID(network string) uint64 {
	switch network {
	case `ethereum`:
		return 1
	case `optimism`:
		return 10
	case `fantom`:
		return 250
	case `arbitrum`:
		return 42161
	case `polygon`:
		return 137
	case `avalanche`:
		return 43114
	case `bsc`:
		return 56
	}
	return 0
}

func fetchPortalsTokenList() []TokenListToken {
	limit := 250
	page := 0
	tokens := []TokenListToken{}

	for {
		req, _ := http.NewRequest("GET", `https://api.portals.fi/v2/tokens?limit=`+strconv.FormatInt(int64(limit), 10)+`&page=`+strconv.FormatInt(int64(page), 10), nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			logs.Error(err)
			break
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logs.Error(err)
			break
		}
		if (resp.StatusCode < 200) || (resp.StatusCode > 299) {
			logs.Error(`impossible to fetch Portals token list`)
			break
		}

		var list TPortalList
		if err := json.Unmarshal(body, &list); err != nil {
			logs.Error(err)
			break
		}

		for _, v := range list.Tokens {
			logoURI := ``
			if len(v.Images) > 0 {
				logoURI = v.Images[0]
			}
			chainID := portalsMapNetworkNameToChainID(v.Network)
			if chainID == 0 || helpers.IsIgnoredChain(chainID) {
				continue
			}
			if (v.Name == `` || v.Symbol == `` || v.Decimals == 0) || helpers.IsIgnoredToken(chainID, common.HexToAddress(v.Address)) {
				continue
			}

			tokens = append(tokens, TokenListToken{
				ChainID:  int(chainID),
				Decimals: v.Decimals,
				Address:  common.HexToAddress(v.Address).Hex(),
				Name:     v.Name,
				Symbol:   v.Symbol,
				LogoURI:  logoURI,
			})
		}
		if !list.More {
			break
		}
		page++
	}
	return tokens
}

func buildPortalsTokenList() {
	tokenList := loadTokenListFromJsonFile(`portals.json`)
	tokenList.Name = "Portals Token List"
	tokenList.LogoURI = "https://portals-assets-bucket.s3.amazonaws.com/logo.png"

	tokens := fetchPortalsTokenList()
	saveTokenListInJsonFile(tokenList, tokens, `portals.json`, Standard)
}
