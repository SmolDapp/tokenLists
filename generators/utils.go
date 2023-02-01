package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/helpers"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

type JSONTokensMethods string

const (
	Standard JSONTokensMethods = "Standard"
	Append   JSONTokensMethods = "Append"
)

// create creates a file with the given path and returns the file object
func create(p string) error {
	if err := os.MkdirAll(p, 0770); err != nil {
		return err
	}
	return nil
}

// loadTokenListFromJsonFile loads a token list from a json file
func loadTokenListFromJsonFile(filePath string) TokenListData {
	var tokenList TokenListData
	content, err := ioutil.ReadFile(helpers.BASE_PATH + `/lists/` + filePath)
	if err != nil {
		logs.Error(err)
		return initEmptyTokenList()
	}
	err = json.Unmarshal(content, &tokenList)
	if err != nil {
		logs.Error(err)
		return initEmptyTokenList()
	}

	tokenList.PreviousTokensMap = make(map[string]TokenListToken)
	for _, token := range tokenList.Tokens {
		if helpers.IsIgnoredChain(uint64(token.ChainID)) {
			continue
		}
		key := GetKey(uint64(token.ChainID), common.HexToAddress(token.Address))
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
	method JSONTokensMethods,
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
			if helpers.IsIgnoredChain(uint64(token.ChainID)) {
				continue
			}
			if (token.Name == `` || token.Symbol == `` || token.Decimals == 0) || helpers.IsIgnoredToken(uint64(token.ChainID), common.HexToAddress(token.Address)) {
				continue
			}
			tokenList.NextTokensMap[GetKey(uint64(token.ChainID), common.HexToAddress(token.Address))] = token
		}
	}

	for _, token := range tokens {
		if helpers.IsIgnoredChain(uint64(token.ChainID)) {
			continue
		}
		tokenList.NextTokensMap[GetKey(uint64(token.ChainID), common.HexToAddress(token.Address))] = token
	}

	if len(tokenList.NextTokensMap) == 0 {
		return errors.New(`token list is empty`)
	}
	tokenList.Timestamp = time.Now().Format(time.RFC3339)
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
		key := GetKey(uint64(token.ChainID), common.HexToAddress(token.Address))
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
		if (token.Name == `` || token.Symbol == `` || token.Decimals == 0) || helpers.IsIgnoredToken(uint64(token.ChainID), common.HexToAddress(token.Address)) {
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
		if helpers.IsIgnoredChain(chainID) {
			continue
		}
		chainIDStr := strconv.FormatUint(chainID, 10)
		tokenList.Tokens = tokens
		jsonData, err := json.MarshalIndent(tokenList, "", "  ")
		if err != nil {
			return err
		}
		create(helpers.BASE_PATH + `/lists/` + chainIDStr)
		if err = ioutil.WriteFile(helpers.BASE_PATH+`/lists/`+chainIDStr+`/`+filePath, jsonData, 0644); err != nil {
			logs.Error(err)
			return err
		}
	}

	return nil
}

// initEmptyTokenList initializes an empty token list
func initEmptyTokenList() TokenListData {
	newTokenList := TokenListData{
		Name:      ``,
		Timestamp: ``,
		Tags:      []string{},
		Keywords:  []string{},
		Tokens:    []TokenListToken{},
	}
	newTokenList.Version.Major = 0
	newTokenList.Version.Minor = 0
	newTokenList.Version.Patch = 0
	newTokenList.PreviousTokensMap = make(map[string]TokenListToken)
	newTokenList.NextTokensMap = make(map[string]TokenListToken)

	return newTokenList
}

// GetKey returns the key of a token in a specific format to make it sortable
func GetKey(chainID uint64, address common.Address) string {
	chainIDStr := strconv.FormatUint(uint64(chainID), 10)
	chainIDStr = strings.Repeat("0", 18-len(chainIDStr)) + chainIDStr
	return chainIDStr + `_` + address.Hex()
}

// decodeString decodes a string from a slice of interfaces
func decodeString(something []interface{}) string {
	if len(something) == 0 {
		return ""
	}
	return something[0].(string)
}

// decodeUint64 decodes a uint64 from a slice of interfaces
func decodeUint64(something []interface{}) uint64 {
	if len(something) == 0 {
		return 0
	}
	return uint64(something[0].(uint8))
}
