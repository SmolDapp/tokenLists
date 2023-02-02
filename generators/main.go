package main

import (
	"os"
	"strings"

	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

var instructionToFunction = map[string]func(){
	`coingecko`:       buildCoingeckoTokenList,
	`1inch`:           build1InchTokenList,
	`paraswap`:        buildParaswapTokenList,
	`defillama`:       buildDefillamaTokenList,
	`yearn`:           buildYearnTokenList,
	`curve`:           buildCurveTokenList,
	`cowswap`:         buildCowswapTokenList,
	`uniswap`:         buildUniswapTokenList,
	`sushiswap`:       buildSushiswapTokenList,
	`ledger`:          buildLedgersTokenList,
	`portals`:         buildPortalsTokenList,
	`wido`:            buildWidoTokenList,
	`uniswap-pairs`:   buildUniswapPairsTokenList,
	`sushiswap-pairs`: buildSushiswapPairsTokenList,
}

func main() {
	helpers.Init()
	ethereum.Init()

	if len(os.Args) < 2 {
		for name, generator := range instructionToFunction {
			logs.Info(`Running generator:`, strings.ToTitle(name))
			generator()
			logs.Success(`Done!`)
		}
	} else {
		for _, arg := range os.Args[1:] {
			logs.Info(`Running generator:`, strings.ToTitle(arg))
			instructionToFunction[arg]()
			logs.Success(`Done!`)
		}
	}
	buildSummary()
}
