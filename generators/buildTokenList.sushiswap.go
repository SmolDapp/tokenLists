package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

type TSushiContracts struct {
	ContractAddress common.Address
	BlockNumber     *big.Int
}

var SushiswapContractsPerChainID = map[uint64][]TSushiContracts{
	1: {
		{
			ContractAddress: common.HexToAddress(`0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac`),
			BlockNumber:     big.NewInt(10_794_229),
		},
	},
	56: {
		{
			ContractAddress: common.HexToAddress(`0xc35DADB65012eC5796536bD9864eD8773aBc74C4`),
			BlockNumber:     big.NewInt(5_205_069),
		},
	},
	100: {
		{
			ContractAddress: common.HexToAddress(`0xc35DADB65012eC5796536bD9864eD8773aBc74C4`),
			BlockNumber:     big.NewInt(14_735_904),
		},
	},
	137: {
		{
			ContractAddress: common.HexToAddress(`0xc35DADB65012eC5796536bD9864eD8773aBc74C4`),
			BlockNumber:     big.NewInt(11_333_218),
		},
	},
	250: {
		{
			ContractAddress: common.HexToAddress(`0xc35DADB65012eC5796536bD9864eD8773aBc74C4`),
			BlockNumber:     big.NewInt(2_457_879),
		},
	},
	42161: {
		{
			ContractAddress: common.HexToAddress(`0xc35DADB65012eC5796536bD9864eD8773aBc74C4`),
			BlockNumber:     big.NewInt(70),
		},
	},
	43114: {
		{
			ContractAddress: common.HexToAddress(`0xc35DADB65012eC5796536bD9864eD8773aBc74C4`),
			BlockNumber:     big.NewInt(506_190),
		},
	},
}

func buildSushiswapTokenList() {
	tokenList := loadTokenListFromJsonFile(`sushiswap.json`)
	originalTokenList := helpers.FetchJSON[TokenListData](`https://token-list.sushi.com/`)
	tokenList.Name = originalTokenList.Name
	tokenList.LogoURI = originalTokenList.LogoURI
	tokenList.Keywords = originalTokenList.Keywords
	tokenList = tokenList.Assign(originalTokenList.Tokens)

	saveTokenListInJsonFile(tokenList, tokenList.Tokens, `sushiswap.json`, Standard)
}
