package main

import (
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

func main() {
	helpers.Init()
	ethereum.Init()
	logs.Info(`Welcome to the tokenLists generator!`)
	// buildLedgersTokenList()
	// buildPortalsTokenList()
	// buildWidoTokenList()
	buildCoingeckoTokenList()
	// build1InchTokenList()
	// buildParaswapTokenList()
	// buildDefillamaTokenList()
	// buildYearnTokenList()
	// buildCurveTokenList()
	// buildCowswapTokenList()

	// buildUniswapTokenList()
	// buildUniswapPairsTokenList()

	// buildSushiswapTokenList()
	// buildSushiswapPairsTokenList()
	logs.Info(`Done!`)
}
