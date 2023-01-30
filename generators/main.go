package main

import (
	"github.com/migratooor/tokenLists/generators/common/logs"
)

func main() {
	logs.Info(`Welcome to the tokenLists generator!`)
	buildLedgersTokenList()
	logs.Info(`Done!`)
}
