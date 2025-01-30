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
		key := GetKey(token.ChainID, token.Address)
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
	listName := BASE_PATH + `/lists/` + filePath
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
			if (token.Name == `` || token.Symbol == `` || token.Decimals == 0) || chains.IsTokenIgnored(token.ChainID, token.Address) {
				continue
			}
			newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				token.LogoURI,
				token.ChainID,
				token.Decimals,
			)
			if err != nil {
				continue
			}
			newToken.Occurrence = token.Occurrence
			tokenList.NextTokensMap[GetKey(token.ChainID, token.Address)] = newToken
		}
	}

	for _, token := range tokens {
		if !chains.IsChainIDSupported(token.ChainID) {
			continue
		}
		newToken, err := SetToken(
			token.Address,
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
		newToken.Occurrence = token.Occurrence
		tokenList.NextTokensMap[GetKey(token.ChainID, token.Address)] = newToken
	}

	/**************************************************************************
	** If the list is empty, we skip
	**************************************************************************/
	if len(tokenList.NextTokensMap) == 0 {
		return errors.New(`token list is empty for ` + listName)
	}

	/**************************************************************************
	** If the chain contains only the default eeee coin or only the extra tokens
	** we don't need to save the list.
	**************************************************************************/
	baseCoinCount := 0
	for _, token := range tokenList.NextTokensMap {
		if !chains.IsChainIDSupported(token.ChainID) {
			continue
		}
		if strings.EqualFold(`0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`, token.Address) {
			baseCoinCount++
		}
	}
	for _, chainID := range chains.SUPPORTED_CHAIN_IDS {
		chain := chains.CHAINS[chainID]
		baseCoinCount += len(chain.ExtraTokens)
	}

	if len(tokenList.NextTokensMap) <= baseCoinCount {
		return errors.New(`token list is empty for ` + listName)
	}

	tokenList.Timestamp = time.Now().Format(time.RFC3339)
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
		key := GetKey(token.ChainID, token.Address)
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
		return errors.New(`no changes detected for ` + listName)
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
		if (token.Name == `` || token.Symbol == `` || token.Decimals == 0) || chains.IsTokenIgnored(token.ChainID, token.Address) {
			continue
		}
		tokenList.Tokens[i] = tokenList.NextTokensMap[k]
		tokenListPerChainID[chainID] = append(tokenListPerChainID[chainID], tokenList.NextTokensMap[k])
	}

	if filePath == `popular.json` {
		occurrence := func(p1, p2 *models.TokenListToken) bool {
			return p1.Occurrence > p2.Occurrence
		}
		By(occurrence).Sort(tokenList.Tokens)
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

		if len(tokens) <= len(chains.CHAINS[chainID].ExtraTokens)+1 {
			logs.Info(`No need to save the list ` + filePath + ` for chainID ` + chainIDStr)
			continue //If we have as much tokens as the extra tokens, we don't need to save the list, this is the default list
		}

		if filePath == `popular.json` {
			occurrence := func(p1, p2 *models.TokenListToken) bool {
				return p1.Occurrence > p2.Occurrence
			}
			By(occurrence).Sort(tokens)
		}
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

// SaveChainListInJsonFile saves a chain list in a json file
func SaveChainListInJsonFile(
	tokenList models.TokenListData[models.TokenListToken],
) error {
	for _, token := range tokenList.Tokens {
		if !chains.IsChainIDSupported(token.ChainID) {
			continue
		}
		newToken, err := SetToken(
			token.Address,
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
		newToken.Occurrence = token.Occurrence
		tokenList.NextTokensMap[GetKey(token.ChainID, token.Address)] = newToken
	}

	/**************************************************************************
	** If the list is empty, we skip
	**************************************************************************/
	if len(tokenList.NextTokensMap) == 0 {
		logs.Error(`token list ` + tokenList.Name + ` is empty`)
		return errors.New(`token list is empty`)
	}

	/**************************************************************************
	** If the chain contains only the default eeee coin or only the extra tokens
	** we don't need to save the list.
	**************************************************************************/
	baseCoinCount := 0
	for _, token := range tokenList.NextTokensMap {
		if !chains.IsChainIDSupported(token.ChainID) {
			continue
		}
		if strings.EqualFold(`0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`, token.Address) {
			baseCoinCount++
		}
		if strings.EqualFold(`So11111111111111111111111111111111111111111`, token.Address) {
			baseCoinCount++
		}
	}
	for _, chainID := range chains.SUPPORTED_CHAIN_IDS {
		chain := chains.CHAINS[chainID]
		baseCoinCount += len(chain.ExtraTokens)
	}

	if len(tokenList.NextTokensMap) <= baseCoinCount {
		logs.Error(`token list ` + tokenList.Name + ` is empty`)
		return errors.New(`token list is empty`)
	}

	tokenList.Timestamp = time.Now().Format(time.RFC3339)
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
		key := GetKey(token.ChainID, token.Address)
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
		logs.Warning(`no changes detected for ` + tokenList.Name)
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
		if (token.Name == `` || token.Symbol == `` || token.Decimals == 0) || chains.IsTokenIgnored(token.ChainID, token.Address) {
			continue
		}
		tokenList.Tokens[i] = tokenList.NextTokensMap[k]
		tokenListPerChainID[chainID] = append(tokenListPerChainID[chainID], tokenList.NextTokensMap[k])
	}

	/**************************************************************************
	** Then, for each chain, we will save the list in a json file
	**************************************************************************/
	for chainID, tokens := range tokenListPerChainID {
		if !chains.IsChainIDSupported(chainID) {
			logs.Info(`ChainID ` + strconv.FormatUint(chainID, 10) + ` is not supported`)
			continue
		}
		logs.Info(`Saving chain list for chainID ` + strconv.FormatUint(chainID, 10))
		if chains.CHAINS[chainID].IsTestNet {
			logs.Info(`ChainID ` + strconv.FormatUint(chainID, 10) + ` is a testnet`)
			continue
		}
		chainIDStr := strconv.FormatUint(chainID, 10)

		if len(tokens) <= len(chains.CHAINS[chainID].ExtraTokens)+1 {
			logs.Info(`No need to save the list for chainID ` + chainIDStr)
			continue //If we have as much tokens as the extra tokens, we don't need to save the list, this is the default list
		}

		/**************************************************************************
		** Sort by occurence
		**************************************************************************/
		occurrence := func(p1, p2 *models.TokenListToken) bool {
			return p1.Occurrence > p2.Occurrence
		}
		By(occurrence).Sort(tokens)

		/**************************************************************************
		** Save the list
		**************************************************************************/
		tokenList.Name = chains.CHAINS[chainID].Name
		tokenList.LogoURI = `https://assets.smold.app/chains/` + chainIDStr + `/logo-128.png`
		tokenList.Tokens = tokens
		jsonData, err := json.MarshalIndent(tokenList, "", "  ")
		if err != nil {
			logs.Error(err)
			return err
		}
		if err = os.WriteFile(BASE_PATH+`/lists/`+chainIDStr+`.json`, jsonData, 0644); err != nil {
			logs.Error(err)
			return err
		}
	}

	return nil
}
