package main

import (
	"context"

	graphql "github.com/hasura/go-graphql-client"
)

var chainIDToBIP44 = map[string]uint64{
	`60`:   1,     // Ethereum
	`614`:  10,    // Optimism
	`714`:  56,    // Binance Smart Chain
	`700`:  100,   // xDai/Gnosis
	`966`:  137,   // Polygon
	`1007`: 250,   // Fantom
	`804`:  324,   // ZKSync
	`9001`: 42161, // Arbitrum
	`9000`: 43114, // Avalanche
}

func fetchTNSTokeList() []TokenListToken {
	listPerChainID := []TokenListToken{}
	client := graphql.NewClient("https://api.thegraph.com/subgraphs/name/mike-data-nexus/tkn-_sg", nil)

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
		return listPerChainID
	}

	for _, domain := range query.Domains {
		for _, address := range domain.Resolver.Addresses {
			expectedChainID := chainIDToBIP44[address.CoinType]
			if expectedChainID == 0 {
				continue
			}
			listPerChainID = append(listPerChainID, TokenListToken{
				Address: address.Address,
				ChainID: expectedChainID,
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
