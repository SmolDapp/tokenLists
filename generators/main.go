package main

import (
	"os"
	"strings"

	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

type TGenerators struct {
	Exec        func()
	Name        string
	Description string
}

var instructionToFunction = map[string]TGenerators{
	`coingecko`: {
		Exec:        buildCoingeckoTokenList,
		Name:        `CoinGecko`,
		Description: `A list of tokens available showing in CoinGecko data agregator.`,
	},
	`1inch`: {
		Exec:        build1InchTokenList,
		Name:        `1Inch`,
		Description: `A list of tokens available in 1Inch DeFi / DEX aggregator`,
	},
	`paraswap`: {
		Exec:        buildParaswapTokenList,
		Name:        `Paraswap`,
		Description: `A list of tokens available for trading on Paraswap DEX`,
	},
	`defillama`: {
		Exec:        buildDefillamaTokenList,
		Name:        `DefiLlama`,
		Description: `A list of tokens available in DefiLlama token service`,
	},
	`yearn`: {
		Exec:        buildYearnTokenList,
		Name:        `Yearn`,
		Description: `A list of tokens available for depositing in Yearn, as well as the tokens that represent yVaults.`,
	},
	`curve`: {
		Exec:        buildCurveTokenList,
		Name:        `Curve`,
		Description: `A list of tokens available for trading on Curve, the largest stableswap.`,
	},
	`cowswap`: {
		Exec:        buildCowswapTokenList,
		Name:        `Cow Swap`,
		Description: `A list of tokens available for trading on CoW Swap, a DEX focused on MEV protection.`,
	},
	`uniswap`: {
		Exec:        buildUniswapTokenList,
		Name:        `UniSwap`,
		Description: `A list of tokens available on UniSwap DEX.`,
	},
	`sushiswap`: {
		Exec:        buildSushiswapTokenList,
		Name:        `SushiSwap`,
		Description: `A list of tokens available on SushiSwap DEX.`,
	},
	`ledger`: {
		Exec:        buildLedgersTokenList,
		Name:        `Ledger`,
		Description: `A list of tokens supported in Ledger Live App`,
	},
	`portals`: {
		Exec:        buildPortalsTokenList,
		Name:        `Portals`,
		Description: `A list of tokens available for trading on Portals DEX.`,
	},
	`wido`: {
		Exec:        buildWidoTokenList,
		Name:        `Wido`,
		Description: `A list of tokens supported by the Wido Router`,
	},
	`uniswap-pairs`: {
		Exec:        buildUniswapPairsTokenList,
		Name:        `UniSwap (pairs)`,
		Description: `A list of token pairs (liquidity pools) available for trading on UniSwap.`,
	},
	`sushiswap-pairs`: {
		Exec:        buildSushiswapPairsTokenList,
		Name:        `SushiSwap (pairs)`,
		Description: `A list of token pairs (liquidity pools) available for trading on SushiSwap.`,
	},
	`optimism`: {
		Exec:        buildOptimismTokenList,
		Name:        `Optimism`,
		Description: `A list of okens used as the source of truth for the Optimism Gateway.`,
	},
	`tokenlistooor`: {
		Exec:        buildTokenListooorList,
		Name:        `Tokenlistooor`,
		Description: `An aggregated list of tokens from Paraswap, Yearn, and Curve`,
	},
}

func main() {
	helpers.Init()
	ethereum.Init()
	if len(os.Args) < 2 {
		for Name, generator := range instructionToFunction {
			logs.Info(`Running generator:`, strings.ToTitle(Name))
			generator.Exec()
			logs.Success(`Done!`)
		}
	} else {
		for _, arg := range os.Args[1:] {
			logs.Info(`Running generator:`, strings.ToTitle(arg))
			instructionToFunction[arg].Exec()
			logs.Success(`Done!`)
		}
	}
	buildSummary()
}
