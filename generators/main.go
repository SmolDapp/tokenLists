package main

import (
	"os"

	"github.com/migratooor/tokenLists/generators/common/logs"
)

func main() {
	logs.Info(`Welcome to the tokenLists generator!`)
	logs.Pretty(os.Getenv(`ENV_TEXT`), os.Getenv(`ENV_TEXT`) == `working`)
	// buildLedgersTokenList()
	// buildPortalsTokenList()
	// buildWidoTokenList()
	// buildCoingeckoTokenList()
	// build1InchTokenList()
	// buildParaswapTokenList()
	// buildDefillamaTokenList()
	buildYearnTokenList()
	// buildCurveTokenList()
	// buildCowswapTokenList()

	// buildUniswapTokenList()
	// buildUniswapPairsTokenList()

	// buildSushiswapTokenList()
	// buildSushiswapPairsTokenList()
	logs.Info(`Done!`)
}
