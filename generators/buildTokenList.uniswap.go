package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

type TUniContracts struct {
	ContractAddress common.Address
	BlockNumber     *big.Int
	Type            int
}

var UniswapContractsPerChainID = map[uint64][]TUniContracts{
	1: {
		{
			ContractAddress: common.HexToAddress(`0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f`),
			BlockNumber:     big.NewInt(10_000_835),
			Type:            2,
		},
		{
			ContractAddress: common.HexToAddress(`0x1F98431c8aD98523631AE4a59f267346ea31F984`),
			BlockNumber:     big.NewInt(12_369_621),
			Type:            3,
		},
	},
	10: {
		{
			ContractAddress: common.HexToAddress(`0x1F98431c8aD98523631AE4a59f267346ea31F984`),
			BlockNumber:     big.NewInt(100_000),
			Type:            3,
		},
	},
	137: {
		{
			ContractAddress: common.HexToAddress(`0x1F98431c8aD98523631AE4a59f267346ea31F984`),
			BlockNumber:     big.NewInt(22_757_547),
			Type:            3,
		},
	},
	42161: {
		{
			ContractAddress: common.HexToAddress(`0x1F98431c8aD98523631AE4a59f267346ea31F984`),
			BlockNumber:     big.NewInt(165),
			Type:            3,
		},
	},
}

func buildUniswapTokenList() {
	tokenList := loadTokenListFromJsonFile(`uniswap.json`)
	originalTokenList := helpers.FetchJSON[TokenListData[TokenListToken]](`https://tokens.uniswap.org`)
	tokenList.Name = helpers.SafeString(originalTokenList.Name, `Uniswap Token List`)
	tokenList.LogoURI = helpers.SafeString(originalTokenList.LogoURI, `ipfs://QmNa8mQkrNKp1WEEeGjFezDmDeodkWRevGFN8JCV7b4Xir"`)
	tokenList.Keywords = originalTokenList.Keywords
	tokenList = tokenList.Assign(originalTokenList.Tokens)

	tokens := fetchTokenList(tokenList.Tokens)
	saveTokenListInJsonFile(tokenList, tokens, `uniswap.json`, Standard)
}
