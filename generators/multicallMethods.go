package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/contracts"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/helpers"
)

/**************************************************************************************************
** The multicall require a specific format for the call data. The following functions are helpers
** used to build them for some specific methods.
**************************************************************************************************/

var ERC20ABI, _ = contracts.ERC20MetaData.GetAbi()

func getName(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := ERC20ABI.Pack("name")
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress,
			Abi:      ERC20ABI,
			Method:   `name`,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ERC20ABI,
		Method:   `name`,
		CallData: parsedData,
		Name:     name,
	}
}
func getSymbol(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := ERC20ABI.Pack("symbol")
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress,
			Abi:      ERC20ABI,
			Method:   `symbol`,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ERC20ABI,
		Method:   `symbol`,
		CallData: parsedData,
		Name:     name,
	}
}
func getDecimals(name string, contractAddress common.Address) ethereum.Call {
	parsedData, err := ERC20ABI.Pack("decimals")
	if err != nil {
		return ethereum.Call{
			Target:   contractAddress,
			Abi:      ERC20ABI,
			Method:   `decimals`,
			CallData: nil,
			Name:     name,
		}
	}
	return ethereum.Call{
		Target:   contractAddress,
		Abi:      ERC20ABI,
		Method:   `decimals`,
		CallData: parsedData,
		Name:     name,
	}
}

/**************************************************************************************************
** fetchBasicInformations will, for a list of addresses, fetch all the relevant basic information
** for the related token. This includes the name, the symbol, the decimals and the details about
** the "underlying tokens" (ex: dai for aDAI). Supported underlying are Curve, AAVE, Compound.
** Yearn's underlying are already in the initial token list.
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

func fetchBasicInformations(chainID uint64, tokens []common.Address) map[string]*TERC20 {
	/**********************************************************************************************
	** The first step is to prepare the multicall, connecting to the multicall instance and
	** preparing the array of calls to send. All calls for all tokens will be send in a single
	** multicall and will later be accessible via a concatened string `tokenAddress + methodName`.
	**********************************************************************************************/
	caller := ethereum.MulticallClientForChainID[chainID]
	calls := []ethereum.Call{}
	for _, token := range tokens {
		calls = append(calls, getName(token.String(), token))
		calls = append(calls, getSymbol(token.String(), token))
		calls = append(calls, getDecimals(token.String(), token))
	}

	/**********************************************************************************************
	** Regular fix for Fantom's RPC, which limit the number of calls in a multicall to a very low
	** number. We split the multicall in multiple calls of 3 calls each.
	** Otherwise, we just send the multicall as is.
	**********************************************************************************************/
	// maxBatch := math.MaxInt64
	// if chainID == 250 {
	maxBatch := 100
	// }

	/**********************************************************************************************
	** Then we can proceed the responses. We will create a new relatedTokensList to be able to know
	** which token to fetch then (ex: for aDAI, we also need to fetch the DAI token).
	** Nb: A special case is for Ethereum coin, which is defaulted as address 0xEeeee....EEeE.
	**********************************************************************************************/
	tokenList := make(map[string]*TERC20)
	response := caller.ExecuteByBatch(calls, maxBatch, nil)
	for _, token := range tokens {
		rawName := response[token.String()+`name`]
		rawSymbol := response[token.String()+`symbol`]
		rawDecimals := response[token.String()+`decimals`]

		newToken := &TERC20{
			Address:  token,
			Name:     helpers.DecodeString(rawName),
			Symbol:   helpers.DecodeString(rawSymbol),
			Decimals: helpers.DecodeUint64(rawDecimals),
		}
		tokenList[token.Hex()] = newToken
	}

	return tokenList
}
