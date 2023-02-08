package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

type JSONSaveTokensMethods string

const (
	Standard JSONSaveTokensMethods = "Standard"
	Append   JSONSaveTokensMethods = "Append"
)

// loadTokenListFromJsonFile loads a token list from a json file
func loadTokenListFromJsonFile(filePath string) TokenListData {
	var tokenList TokenListData
	content, err := ioutil.ReadFile(helpers.BASE_PATH + `/lists/` + filePath)
	if err != nil {
		logs.Error(err)
		return InitTokenList()
	}
	if err = json.Unmarshal(content, &tokenList); err != nil {
		logs.Error(err)
		return InitTokenList()
	}

	tokenList.PreviousTokensMap = make(map[string]TokenListToken)
	for _, token := range tokenList.Tokens {
		if helpers.IsChainIDIgnored(token.ChainID) {
			continue
		}
		key := getKey(token.ChainID, common.HexToAddress(token.Address))
		tokenList.PreviousTokensMap[key] = token
	}
	tokenList.NextTokensMap = make(map[string]TokenListToken)
	return tokenList
}

// saveTokenListInJsonFile saves a token list in a json file
func saveTokenListInJsonFile(
	tokenList TokenListData,
	tokens []TokenListToken,
	filePath string,
	method JSONSaveTokensMethods,
) error {
	/**************************************************************************
	** First part is transforming the token list into a map. This will allow
	** us to detect the changes in the token list and to directly access a
	** token by its address.
	** If the method is set to "Append", we will first need to load the
	** previous token list.
	**************************************************************************/
	if method == Append {
		for _, token := range tokenList.PreviousTokensMap {
			if helpers.IsChainIDIgnored(token.ChainID) {
				continue
			}
			if (token.Name == `` || token.Symbol == `` || token.Decimals == 0) || helpers.IsIgnoredToken(token.ChainID, common.HexToAddress(token.Address)) {
				continue
			}
			tokenList.NextTokensMap[getKey(token.ChainID, common.HexToAddress(token.Address))] = token
		}
	}

	for _, token := range tokens {
		if helpers.IsChainIDIgnored(token.ChainID) {
			continue
		}
		tokenList.NextTokensMap[getKey(token.ChainID, common.HexToAddress(token.Address))] = token
	}

	if len(tokenList.NextTokensMap) == 0 {
		return errors.New(`token list is empty`)
	}
	tokenList.Timestamp = time.Now().UTC().Format(`02/01/2006 15:04:05`)
	tokenList.Tokens = []TokenListToken{}

	/**************************************************************************
	** Detect the changes in the token list.
	** If a token is removed, the major version is bumped.
	** If a token is added, the minor version is bumped.
	** If a token is modified, the patch version is bumped.
	** Skip if we are not using the standard method.
	**************************************************************************/
	shouldBumpMajor := false
	shouldBumpMinor := false
	shouldBumpPatch := false
	for _, token := range tokenList.NextTokensMap {
		key := getKey(token.ChainID, common.HexToAddress(token.Address))
		if _, ok := tokenList.PreviousTokensMap[key]; !ok {
			shouldBumpMinor = true
		} else if !reflect.DeepEqual(token, tokenList.PreviousTokensMap[key]) {
			shouldBumpPatch = true
		}
		tokenList.Tokens = append(tokenList.Tokens, token)
		delete(tokenList.PreviousTokensMap, key)
	}
	if len(tokenList.PreviousTokensMap) > 0 {
		shouldBumpMajor = true
	}

	if shouldBumpMajor {
		tokenList.Version.Major++
		tokenList.Version.Minor = 0
		tokenList.Version.Patch = 0
	} else if shouldBumpMinor {
		tokenList.Version.Minor++
		tokenList.Version.Patch = 0
	} else if shouldBumpPatch {
		tokenList.Version.Patch++
	}

	/**************************************************************************
	** To make it easy to work with the list, we will sort the token by their
	** chainID in ascending order.
	**************************************************************************/
	tokeListPerChainID := make(map[uint64][]TokenListToken)
	keys := make([]string, 0, len(tokenList.NextTokensMap))
	for k := range tokenList.NextTokensMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		chainID, _ := strconv.ParseUint(strings.Split(k, `_`)[0], 10, 64)
		if _, ok := tokeListPerChainID[chainID]; !ok {
			tokeListPerChainID[chainID] = []TokenListToken{}
		}

		token := tokenList.NextTokensMap[k]
		if (token.Name == `` || token.Symbol == `` || token.Decimals == 0) || helpers.IsIgnoredToken(token.ChainID, common.HexToAddress(token.Address)) {
			continue
		}
		tokenList.Tokens[i] = tokenList.NextTokensMap[k]
		tokeListPerChainID[chainID] = append(tokeListPerChainID[chainID], tokenList.NextTokensMap[k])
	}

	/**************************************************************************
	** Then we will just save the unified token list in a json file as well as
	** each individual token list per chainID.
	**************************************************************************/
	jsonData, err := json.MarshalIndent(tokenList, "", "  ")
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(helpers.BASE_PATH+`/lists/`+filePath, jsonData, 0644); err != nil {
		return err
	}

	for chainID, tokens := range tokeListPerChainID {
		if helpers.IsChainIDIgnored(chainID) {
			continue
		}
		chainIDStr := strconv.FormatUint(chainID, 10)
		tokenList.Tokens = tokens
		jsonData, err := json.MarshalIndent(tokenList, "", "  ")
		if err != nil {
			return err
		}
		helpers.CreateFile(helpers.BASE_PATH + `/lists/` + chainIDStr)
		if err = ioutil.WriteFile(helpers.BASE_PATH+`/lists/`+chainIDStr+`/`+filePath, jsonData, 0644); err != nil {
			logs.Error(err)
			return err
		}
	}

	return nil
}

// getKey returns the key of a token in a specific format to make it sortable
func getKey(chainID uint64, address common.Address) string {
	chainIDStr := strconv.FormatUint(chainID, 10)
	chainIDStr = strings.Repeat("0", 18-len(chainIDStr)) + chainIDStr
	return chainIDStr + `_` + address.Hex()
}

func initSyncMap[T any](chainIDs map[uint64]T) sync.Map {
	tokensForChainIDSyncMap := sync.Map{}
	for chainID := range chainIDs {
		tokensForChainIDSyncMap.Store(chainID, []TokenListToken{})
	}
	return tokensForChainIDSyncMap
}

func extractSyncMap(mapper sync.Map) []TokenListToken {
	tokenList := []TokenListToken{}
	mapper.Range(func(chainID, syncMapRaw interface{}) bool {
		syncMap, _ := syncMapRaw.([]TokenListToken)
		tokenList = append(tokenList, syncMap...)
		return true
	})
	return tokenList
}
