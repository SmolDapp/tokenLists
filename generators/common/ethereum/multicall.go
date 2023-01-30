package ethereum

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/migratooor/tokenLists/generators/common/contracts"
)

type Call struct {
	Name     string         `json:"name"`
	Method   string         `json:"method"`
	Version  string         `json:"version"`
	Abi      *abi.ABI       `json:"abi"`
	Target   common.Address `json:"target"`
	CallData []byte         `json:"call_data"`
}

type CallResponse struct {
	Success    bool   `json:"success"`
	ReturnData []byte `json:"returnData"`
}

type TEthMultiCaller struct {
	Signer          *bind.TransactOpts
	Client          *ethclient.Client
	Abi             *abi.ABI
	ContractAddress common.Address
}

func (call Call) GetMultiCall() contracts.Multicall3Call {
	return contracts.Multicall3Call{
		Target:   call.Target,
		CallData: call.CallData,
	}
}

// NewMulticall creates a new instance of a TEthMultiCaller. This is the instance we
// will later use to perform multiple ethereum calls batched in the same transaction.
// For performance reason, this should be initialized once and then reused.
func NewMulticall(rpcURI string, multicallAddress common.Address) TEthMultiCaller {
	if rpcURI == "" {
		return TEthMultiCaller{}
	}
	client, err := ethclient.Dial(rpcURI)
	if err != nil {
		time.Sleep(time.Second)
		return NewMulticall(rpcURI, multicallAddress)
	}

	// Load Multicall abi for later use
	mcAbi, err := contracts.Multicall3MetaData.GetAbi()
	if err != nil {
		time.Sleep(time.Second)
		return NewMulticall(rpcURI, multicallAddress)
	}

	return TEthMultiCaller{
		Signer:          randomSigner(),
		Client:          client,
		Abi:             mcAbi,
		ContractAddress: multicallAddress,
	}
}

func (caller *TEthMultiCaller) execute(
	multiCallGroup []contracts.Multicall3Call,
	blockNumber *big.Int,
) ([]byte, error) {
	// Prepare calldata for multicall
	abi, _ := contracts.Multicall3MetaData.GetAbi()
	callData, err := abi.Pack("tryAggregate", false, multiCallGroup)
	if err != nil {
		return []byte{}, err
	}

	// Perform multicall
	resp, err := caller.Client.CallContract(
		context.Background(),
		ethereum.CallMsg{
			To:   &caller.ContractAddress,
			Data: callData,
			Gas:  0,
			From: caller.Signer.From,
		},
		blockNumber,
	)
	if err != nil {
		return []byte{}, err
	}
	return resp, nil
}

// ExecuteByBatch will take a group of calls, split them in fixed-size group to
// avoid the gasLimit error, and execute as many transactions as required to get
// the results.
func (caller *TEthMultiCaller) ExecuteByBatch(
	calls []Call,
	batchSize int,
	blockNumber *big.Int,
) map[string][]interface{} {
	var responses []CallResponse
	// Create mapping for results. Be aware that we sometimes get two empty results initially, unsure why
	results := make(map[string][]interface{})

	// Add calls to multicall structure for the contract
	var multiCalls = make([]contracts.Multicall3Call, 0, len(calls))
	for _, call := range calls {
		multiCalls = append(multiCalls, call.GetMultiCall())
	}

	for i := 0; i < len(multiCalls); i += batchSize {
		var group []contracts.Multicall3Call

		if i+batchSize >= len(multiCalls) {
			group = multiCalls[i:]
		} else {
			group = multiCalls[i : i+batchSize]
		}

		tempPackedResp, err := caller.execute(group, blockNumber)
		if err != nil {
			continue
		}

		// Unpack results
		unpackedResp, err := caller.Abi.Unpack("tryAggregate", tempPackedResp)
		if err != nil {
			continue
		}

		a, err := json.Marshal(unpackedResp[0])
		if err != nil {
			continue
		}

		// Unpack results
		var tempResp []CallResponse
		if err := json.Unmarshal(a, &tempResp); err != nil {
			continue
		}

		responses = append(responses, tempResp...)
	}

	for i, response := range responses {
		if response.ReturnData != nil {
			unpacked, err := calls[i].Abi.Unpack(calls[i].Method, response.ReturnData)
			if err != nil {
				// logs.Warning("Failed to unpack method " + calls[i].Method + " for " + calls[i].Name + " : " + err.Error())
				results[calls[i].Name+calls[i].Method] = nil
			} else {
				results[calls[i].Name+calls[i].Method] = unpacked
			}
		} else {
			results[calls[i].Name+calls[i].Method] = nil
		}
	}

	return results
}
