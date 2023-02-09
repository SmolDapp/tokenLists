package main

import (
	"os"
	"strings"

	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

func loadAllLists() {
	allTokens := make(map[uint64]map[string]TokenListToken)
	allTokensPlain := []TokenListToken{}

	for name := range GENERATORS {
		logs.Info(`Loading token list: ` + name)
		tokenList := loadTokenListFromJsonFile(name + `.json`)
		for _, token := range tokenList.Tokens {
			if _, ok := allTokens[token.ChainID]; !ok {
				allTokens[token.ChainID] = make(map[string]TokenListToken)
			}
			if existingToken, ok := allTokens[token.ChainID][helpers.ToAddress(token.Address)]; ok {
				// If the token already exists, we keep the existing data and only
				// update the missing ones.
				allTokens[token.ChainID][helpers.ToAddress(token.Address)] = TokenListToken{
					Address:  existingToken.Address,
					Name:     helpers.SafeString(existingToken.Name, token.Name),
					Symbol:   helpers.SafeString(existingToken.Symbol, token.Symbol),
					LogoURI:  helpers.SafeString(existingToken.LogoURI, token.LogoURI),
					Decimals: helpers.SafeInt(existingToken.Decimals, token.Decimals),
					ChainID:  token.ChainID,
				}
			} else {
				allTokens[token.ChainID][helpers.ToAddress(token.Address)] = TokenListToken{
					Address:  helpers.ToAddress(token.Address),
					Name:     helpers.SafeString(token.Name, ``),
					Symbol:   helpers.SafeString(token.Symbol, ``),
					LogoURI:  helpers.SafeString(token.LogoURI, ``),
					Decimals: helpers.SafeInt(token.Decimals, 18),
					ChainID:  token.ChainID,
				}
			}
		}
	}

	for chainID, tokens := range allTokens {
		logs.Info(chainID, len(tokens))
		for _, token := range tokens {
			if token.LogoURI == `` {
				allTokensPlain = append(allTokensPlain, token)

			}
		}
	}
	tokenList := loadTokenListFromJsonFile(`tokenlistor.json`)
	saveTokenListInJsonFile(tokenList, allTokensPlain, `tokenlistor.json`, Standard)

}

func main() {
	// loadAllLists()
	helpers.Init()
	ethereum.Init()
	// return
	if len(os.Args) < 2 {
		for name, generator := range GENERATORS {
			logs.Info(`Running generator:`, strings.ToTitle(name))
			generator.Exec()
			logs.Success(`Done!`)
		}
	} else if (len(os.Args) == 2) && (os.Args[1] == `tokens`) {
		for name, generator := range GENERATORS {
			if generator.GeneratorType == GeneratorToken {
				logs.Info(`Running generator:`, strings.ToTitle(name))
				generator.Exec()
				logs.Success(`Done!`)
			}
		}
	} else if (len(os.Args) == 2) && (os.Args[1] == `pools`) {
		for name, generator := range GENERATORS {
			if generator.GeneratorType == GeneratorPool {
				logs.Info(`Running generator:`, strings.ToTitle(name))
				generator.Exec()
				logs.Success(`Done!`)
			}
		}
	} else {
		for _, arg := range os.Args[1:] {
			logs.Info(`Running generator:`, strings.ToTitle(arg))
			GENERATORS[arg].Exec()
			logs.Success(`Done!`)
		}
	}

	buildSummary()
}
