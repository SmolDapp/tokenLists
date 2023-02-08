package main

import (
	"os"
	"strings"

	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

func main() {
	helpers.Init()
	ethereum.Init()
	if len(os.Args) < 2 {
		for Name, generator := range GENERATORS {
			logs.Info(`Running generator:`, strings.ToTitle(Name))
			generator.Exec()
			logs.Success(`Done!`)
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
