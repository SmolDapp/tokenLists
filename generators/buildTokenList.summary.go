package main

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"time"

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
	Metadata    struct {
		SupportedChains  []int  `json:"supportedChains"`
		GenerationMethod string `json:"generationMethod"`
	} `json:"metadata"`
	Version struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	} `json:"version"`
}

// TTokenListSummary is the summary of a token list for the tokenListooor project
type TTokenListSummary struct {
	Name      string              `json:"name"`
	Timestamp int64               `json:"timestamp"`
	LogoURI   string              `json:"logoURI"`
	Lists     []TMinTokenListData `json:"lists"`
}

var BASE_URI = `https://raw.githubusercontent.com/Migratooor/tokenLists/main/`

func listSupportedChains(list []TokenListToken) []int {
	detectedChainsMap := map[int]bool{}
	detectedChains := []int{}
	for _, token := range list {
		detectedChainsMap[int(token.ChainID)] = true
	}
	for chainID := range detectedChainsMap {
		detectedChains = append(detectedChains, chainID)
	}
	sort.Ints(detectedChains)
	return detectedChains
}

func buildSummary() {
	tokenListSummary := TTokenListSummary{}
	tokenListSummary.Name = `Tokenlistooor summary`
	tokenListSummary.LogoURI = BASE_URI + `.github/tokenlistooor.svg`
	tokenListSummary.Timestamp = time.Now().UTC().Unix()
	for name, data := range GENERATORS {
		tokenList := loadTokenListFromJsonFile(name + `.json`)
		listElement := TMinTokenListData{
			Name:        tokenList.Name,
			Timestamp:   tokenList.Timestamp,
			LogoURI:     tokenList.LogoURI,
			URI:         BASE_URI + `lists/` + name + `.json`,
			Keywords:    tokenList.Keywords,
			Version:     tokenList.Version,
			TokenCount:  len(tokenList.Tokens),
			Description: data.Description,
		}
		listElement.Metadata.SupportedChains = listSupportedChains(tokenList.Tokens)
		listElement.Metadata.GenerationMethod = string(data.GenerationMethod)
		tokenListSummary.Lists = append(tokenListSummary.Lists, listElement)
	}

	//Also add the tokenListooor list
	{
		tokenListooorList := loadTokenListFromJsonFile(`tokenListooor.json`)
		listElement := TMinTokenListData{
			Name:        tokenListooorList.Name,
			Timestamp:   tokenListooorList.Timestamp,
			LogoURI:     tokenListooorList.LogoURI,
			URI:         BASE_URI + `lists/tokenlistooor.json`,
			Keywords:    tokenListooorList.Keywords,
			Version:     tokenListooorList.Version,
			TokenCount:  len(tokenListooorList.Tokens),
			Description: `A curated list of tokens from all the token lists on tokenlistooor.`,
		}
		listElement.Metadata.SupportedChains = listSupportedChains(tokenListooorList.Tokens)
		listElement.Metadata.GenerationMethod = string(GenerationAPI)
		tokenListSummary.Lists = append(tokenListSummary.Lists, listElement)
	}

	sort.Slice(tokenListSummary.Lists, func(i, j int) bool {
		return tokenListSummary.Lists[i].Name < tokenListSummary.Lists[j].Name
	})

	jsonData, _ := json.MarshalIndent(tokenListSummary, "", "  ")
	ioutil.WriteFile(helpers.BASE_PATH+`/lists/summary.json`, jsonData, 0644)
}
