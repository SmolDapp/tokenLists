package main

import (
	"context"
	"strconv"
	"strings"

	graphql "github.com/hasura/go-graphql-client"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

var chainIDToBIP44 = map[string]uint64{
	`0x3c`:       1,     // Ethereum
	`0x8000000a`: 10,    // Optimism
	`0x80000038`: 56,    // Binance Smart Chain
	`0x80000064`: 100,   // xDai/Gnosis
	`0x80000089`: 137,   // Polygon
	`0x800000fa`: 250,   // Fantom
	`0x80000144`: 324,   // ZKSync
	`0x8000a4b1`: 42161, // Arbitrum
	`0x8000a86a`: 43114, // Avalanche
}

func fetchTNSTokeList() []TokenListToken {
	listPerChainID := []TokenListToken{}
	client := graphql.NewClient(
		`https://api.thegraph.com/subgraphs/name/mike-data-nexus/tkn-_sg`,
		nil,
	)
	var query struct {
		Domains []struct {
			Resolver struct {
				Avatar    string
				Addresses []struct {
					CoinType string
					Address  string
				}
			}
		} `graphql:"domains(where: {name_ends_with: \".tkn.eth\"}, first: 1000)"`
	}
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		logs.Error(err)
		return listPerChainID
	}

	for _, domain := range query.Domains {
		for _, address := range domain.Resolver.Addresses {
			coinTypeToInt, err := strconv.ParseInt(address.CoinType, 0, 64)
			if err != nil {
				logs.Error(err)
				continue
			}
			coinTypeHex := strconv.FormatInt(coinTypeToInt, 16)
			expectedChainID := chainIDToBIP44[strings.ToLower(`0x`+coinTypeHex)]
			if expectedChainID == 0 {
				continue
			}
			listPerChainID = append(listPerChainID, TokenListToken{
				Address: address.Address,
				ChainID: expectedChainID,
				LogoURI: domain.Resolver.Avatar,
			})
		}
	}

	return fetchTokenList(listPerChainID)
}

func buildTNSTokenList() {
	tokenList := loadTokenListFromJsonFile(`tns.json`)
	tokenList.Name = `Token Name Service`
	tokenList.LogoURI = `https://logo.assets.tkn.eth.limo/`
	tokenList.Keywords = []string{`tns`, `token`, `tokendao`, `tkn`, `tkr`}
	tokenList.Description = `Token Name Service is a decentralized naming service for tokens on Ethereum.`
	tokens := fetchTNSTokeList()

	saveTokenListInJsonFile(tokenList, tokens, `tns.json`, Standard)
}
