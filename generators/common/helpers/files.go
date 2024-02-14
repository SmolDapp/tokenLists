package helpers

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/logs"
	"github.com/migratooor/tokenLists/generators/common/models"
)

type JSONSaveTokensMethods string

const (
	SavingMethodStandard JSONSaveTokensMethods = "SavingMethodStandard"
	SavingMethodAppend   JSONSaveTokensMethods = "SavingMethodAppend"
)

// BASE_PATH is the base path to access the data informations
var BASE_PATH, _ = filepath.Abs(getCurrentPath() + `../../../../`)

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)

	return path.Dir(filename)
}

// CreateFile creates a file with the given path and returns the file object
func CreateFile(p string) error {
	if err := os.MkdirAll(p, 0770); err != nil {
		return err
	}
	return nil
}

// LoadTokenListFromJsonFile loads a token list from a json file.
func LoadTokenListFromJsonFile(filePath string) models.TokenListData[models.TokenListToken] {
	var tokenList models.TokenListData[models.TokenListToken]
	content, err := os.ReadFile(BASE_PATH + `/lists/` + filePath)
	if err != nil {
		logs.Error(err)
		if errors.Is(err, os.ErrNotExist) {
			os.WriteFile(BASE_PATH+`/lists/`+filePath, []byte(`{}`), 0644)
		}
		return models.InitTokenList()
	}
	if err = json.Unmarshal(content, &tokenList); err != nil {
		logs.Error(err)
		return models.InitTokenList()
	}

	tokenList.PreviousTokensMap = make(map[string]models.TokenListToken)
	for _, token := range tokenList.Tokens {
		if !chains.IsChainIDSupported(token.ChainID) {
			continue
		}
		key := GetKey(token.ChainID, common.HexToAddress(token.Address))
		tokenList.PreviousTokensMap[key] = token
	}
	tokenList.NextTokensMap = make(map[string]models.TokenListToken)
	return tokenList
}

// SaveTokenListInJsonFile saves a token list in a json file
func SaveTokenListInJsonFile(
	tokenList models.TokenListData[models.TokenListToken],
	tokensMaybeDuplicates []models.TokenListToken,
	filePath string,
	method JSONSaveTokensMethods,
) error {
	tokens := []models.TokenListToken{}
	addresses := make(map[string]bool)
	for _, token := range tokensMaybeDuplicates {
		key := token.Address + strconv.FormatUint(token.ChainID, 10)
		if _, ok := addresses[key]; !ok {
			addresses[key] = true
			tokens = append(tokens, token)
		}
	}

	/**************************************************************************
	** First part is transforming the token list into a map. This will allow
	** us to detect the changes in the token list and to directly access a
	** token by its address.
	** If the method is set to "SavingMethodAppend", we will first need to load the
	** previous token list.
	**************************************************************************/
	if method == SavingMethodAppend {
		for _, token := range tokenList.PreviousTokensMap {
			if !chains.IsChainIDSupported(token.ChainID) {
				continue
			}
			if (token.Name == `` || token.Symbol == `` || token.Decimals == 0) || chains.IsTokenIgnored(token.ChainID, common.HexToAddress(token.Address)) {
				continue
			}
			newToken, err := SetToken(
				common.HexToAddress(token.Address),
				token.Name,
				token.Symbol,
				token.LogoURI,
				token.ChainID,
				token.Decimals,
			)
			if err != nil {
				continue
			}
			tokenList.NextTokensMap[GetKey(token.ChainID, common.HexToAddress(token.Address))] = newToken
		}
	}

	for _, token := range tokens {
		if !chains.IsChainIDSupported(token.ChainID) {
			continue
		}
		newToken, err := SetToken(
			common.HexToAddress(token.Address),
			token.Name,
			token.Symbol,
			token.LogoURI,
			token.ChainID,
			token.Decimals,
		)
		if err != nil {
			logs.Error(err)
			continue
		}
		tokenList.NextTokensMap[GetKey(token.ChainID, common.HexToAddress(token.Address))] = newToken
	}

	if len(tokenList.NextTokensMap) == 0 {
		return errors.New(`token list is empty`)
	}
	tokenList.Timestamp = time.Now().UTC().Format(`02/01/2006 15:04:05`)
	tokenList.Tokens = []models.TokenListToken{}

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
		key := GetKey(token.ChainID, common.HexToAddress(token.Address))
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

	/**************************************************************************
	** If there are no changes, we will just return.
	**************************************************************************/
	if !shouldBumpMajor && !shouldBumpMinor && !shouldBumpPatch {
		return nil
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
	tokenListPerChainID := make(map[uint64][]models.TokenListToken)
	keys := make([]string, 0, len(tokenList.NextTokensMap))
	for k := range tokenList.NextTokensMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		chainID, _ := strconv.ParseUint(strings.Split(k, `_`)[0], 10, 64)
		if _, ok := tokenListPerChainID[chainID]; !ok {
			tokenListPerChainID[chainID] = []models.TokenListToken{}
		}

		token := tokenList.NextTokensMap[k]
		if (token.Name == `` || token.Symbol == `` || token.Decimals == 0) || chains.IsTokenIgnored(token.ChainID, common.HexToAddress(token.Address)) {
			continue
		}
		tokenList.Tokens[i] = tokenList.NextTokensMap[k]
		tokenListPerChainID[chainID] = append(tokenListPerChainID[chainID], tokenList.NextTokensMap[k])
	}

	/**************************************************************************
	** Then we will just save the unified token list in a json file as well as
	** each individual token list per chainID.
	**************************************************************************/
	jsonData, err := json.MarshalIndent(tokenList, "", "  ")
	if err != nil {
		return err
	}
	if err = os.WriteFile(BASE_PATH+`/lists/`+filePath, jsonData, 0644); err != nil {
		return err
	}

	for chainID, tokens := range tokenListPerChainID {
		if !chains.IsChainIDSupported(chainID) {
			continue
		}
		chainIDStr := strconv.FormatUint(chainID, 10)
		tokenList.Tokens = tokens

		jsonData, err := json.MarshalIndent(tokenList, "", "  ")
		if err != nil {
			logs.Error(err)
			return err
		}
		if err := CreateFile(BASE_PATH + `/lists/` + chainIDStr); err != nil {
			logs.Error(err)
			return err
		}

		if err = os.WriteFile(BASE_PATH+`/lists/`+chainIDStr+`/`+filePath, jsonData, 0644); err != nil {
			logs.Error(err)
			return err
		}
	}

	return nil
}
