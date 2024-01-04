package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/contracts"
)

/**************************************************************************************************
** The multicall require a specific format for the call data. The following functions are helpers
** used to build them for some specific methods.
**************************************************************************************************/

var ERC20ABI, _ = contracts.ERC20MetaData.GetAbi()

func getName(name string, contractAddress common.Address) Call {
	parsedData, err := ERC20ABI.Pack("name")
	if err != nil {
		return Call{
			Target:   contractAddress,
			Abi:      ERC20ABI,
			Method:   `name`,
			CallData: nil,
			Name:     name,
		}
	}
	return Call{
		Target:   contractAddress,
		Abi:      ERC20ABI,
		Method:   `name`,
		CallData: parsedData,
		Name:     name,
	}
}
func getSymbol(name string, contractAddress common.Address) Call {
	parsedData, err := ERC20ABI.Pack("symbol")
	if err != nil {
		return Call{
			Target:   contractAddress,
			Abi:      ERC20ABI,
			Method:   `symbol`,
			CallData: nil,
			Name:     name,
		}
	}
	return Call{
		Target:   contractAddress,
		Abi:      ERC20ABI,
		Method:   `symbol`,
		CallData: parsedData,
		Name:     name,
	}
}
func getDecimals(name string, contractAddress common.Address) Call {
	parsedData, err := ERC20ABI.Pack("decimals")
	if err != nil {
		return Call{
			Target:   contractAddress,
			Abi:      ERC20ABI,
			Method:   `decimals`,
			CallData: nil,
			Name:     name,
		}
	}
	return Call{
		Target:   contractAddress,
		Abi:      ERC20ABI,
		Method:   `decimals`,
		CallData: parsedData,
		Name:     name,
	}
}

/**************************************************************************************************
** fetchBasicInformations will, for a list of addresses, fetch all the relevant basic information
** for the related token. This includes the name, the symbol and the decimals.
**
** Arguments:
** - chainID: the chain ID of the network we are working on
** - tokens: a list of addresses of the tokens we want to fetch the information for
**
** Returns:
** - a list of TERC20Token containing the basic information for the tokens
**************************************************************************************************/
type TERC20 struct {
	Address  common.Address
	Name     string
	Symbol   string
	ChainID  uint64
	Decimals uint64
}

func FetchBasicInformations(chainID uint64, tokens []common.Address) map[string]*TERC20 {
	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	caller := MulticallClientForChainID[chainID]
	calls := []Call{}
	for _, token := range tokens {
		calls = append(calls, getName(token.String(), token))
		calls = append(calls, getSymbol(token.String(), token))
		calls = append(calls, getDecimals(token.String(), token))
	}

	/**********************************************************************************************
	** Regular fix for some RPC, which limit the number of calls in a multicall to a very low
	** number.
	**********************************************************************************************/
	maxBatch := uint64(420)
	if chainID == 250 || chainID == 56 || chainID == 137 {
		maxBatch = 420
	}

	/**********************************************************************************************
	** Then we can proceed the responses.
	**********************************************************************************************/
	tokenList := make(map[string]*TERC20)
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	for _, token := range tokens {
		rawName := response[token.String()+`name`]
		rawSymbol := response[token.String()+`symbol`]
		rawDecimals := response[token.String()+`decimals`]

		newToken := &TERC20{
			Address:  token,
			Name:     decodeString(rawName),
			Symbol:   decodeString(rawSymbol),
			Decimals: decodeUint64(rawDecimals),
		}
		tokenList[token.Hex()] = newToken
	}

	return tokenList
}

func FetchNames(chainID uint64, tokens []common.Address) map[string]string {
	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	caller := MulticallClientForChainID[chainID]
	calls := []Call{}
	for _, token := range tokens {
		calls = append(calls, getName(token.String(), token))
	}

	/**********************************************************************************************
	** Regular fix for some RPC, which limit the number of calls in a multicall to a very low
	** number.
	**********************************************************************************************/
	maxBatch := uint64(420)
	if chainID == 250 || chainID == 56 || chainID == 137 {
		maxBatch = 420
	}

	/**********************************************************************************************
	** Then we can proceed the responses.
	**********************************************************************************************/
	nameList := make(map[string]string)
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	for _, token := range tokens {
		rawName := response[token.String()+`name`]
		nameList[token.Hex()] = decodeString(rawName)
	}

	return nameList
}

func FetchDecimals(chainID uint64, tokens []common.Address) map[string]uint64 {
	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	caller := MulticallClientForChainID[chainID]
	calls := []Call{}
	for _, token := range tokens {
		calls = append(calls, getDecimals(token.String(), token))
	}

	/**********************************************************************************************
	** Regular fix for some RPC, which limit the number of calls in a multicall to a very low
	** number.
	**********************************************************************************************/
	maxBatch := uint64(420)
	if chainID == 250 || chainID == 56 || chainID == 137 {
		maxBatch = 42
	}

	/**********************************************************************************************
	** Then we can proceed the responses.
	**********************************************************************************************/
	decimalsList := make(map[string]uint64)
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	for _, token := range tokens {
		rawDecimals := response[token.String()+`decimals`]
		decimalsList[token.Hex()] = decodeUint64(rawDecimals)
	}

	return decimalsList
}
