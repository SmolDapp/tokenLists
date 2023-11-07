package main

import (
	"os"
	"strings"

	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

var ALL_EXISTING_TOKENS = map[uint64]map[string]TokenListToken{}

func main() {
	helpers.Init()
	ethereum.Init()
	loadAllTokenLogoURI()

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
			if _, ok := GENERATORS[arg]; ok {
				logs.Info(`Running generator:`, strings.ToTitle(arg))
				GENERATORS[arg].Exec()
				logs.Success(`Done!`)
			}
		}
	}

	buildTokenListooorList()
	buildSummary()
}
