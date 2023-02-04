package main

import (
	"encoding/json"
	"io/ioutil"
	"sort"

	"github.com/migratooor/tokenLists/generators/common/helpers"
)

// TMinTokenListData is the minimal data of a token list for the tokenListooor project
type TMinTokenListData struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Timestamp   string   `json:"timestamp"`
	LogoURI     string   `json:"logoURI"`
	URI         string   `json:"URI"`
	Keywords    []string `json:"keywords"`
	TokenCount  int      `json:"tokenCount"`
	Version     struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	} `json:"version"`
}

// TTokenListSummary is the summary of a token list for the tokenListooor project
type TTokenListSummary struct {
	Name      string              `json:"name"`
	Timestamp string              `json:"timestamp"`
	LogoURI   string              `json:"logoURI"`
	Lists     []TMinTokenListData `json:"tokens"`
}

var BASE_URI = `https://raw.githubusercontent.com/Migratooor/tokenLists/main/`

func buildSummary() {
	tokenListSummary := TTokenListSummary{}
	tokenListSummary.Name = `Tokenlistooor summary`
	tokenListSummary.LogoURI = BASE_URI + `.github/tokenlistooor.svg`
	for name := range instructionToFunction {
		tokenList := loadTokenListFromJsonFile(name + `.json`)
		tokenListSummary.Lists = append(tokenListSummary.Lists, TMinTokenListData{
			Name:       tokenList.Name,
			Timestamp:  tokenList.Timestamp,
			LogoURI:    tokenList.LogoURI,
			URI:        BASE_URI + `lists/` + name + `.json`,
			Keywords:   tokenList.Keywords,
			Version:    tokenList.Version,
			TokenCount: len(tokenList.Tokens),
		})
	}

	sort.Slice(tokenListSummary.Lists, func(i, j int) bool {
		return tokenListSummary.Lists[i].Name < tokenListSummary.Lists[j].Name
	})

	jsonData, _ := json.MarshalIndent(tokenListSummary, "", "  ")
	ioutil.WriteFile(helpers.BASE_PATH+`/lists/summary.json`, jsonData, 0644)
}
