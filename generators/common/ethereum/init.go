package ethereum

import (
	"context"
	"math"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/logs"
)

// RPC contains the ethclient.Client for a specific chainID
var RPC = map[uint64]*ethclient.Client{}

// MulticallClientForChainID holds the multicall client for a specific chainID
var MulticallClientForChainID = make(map[uint64]TEthMultiCaller)

// RPC_ENDPOINTS contains the node endpoints to connect the blockchains
var RPC_ENDPOINTS = map[uint64]string{}

func Init() {
	godotenv.Load(`.env`)

	bearer := os.Getenv(`BEARER_FOR_1INCH`)
	if bearer == "" {
		logs.Warning(`Missing environment variable BEARER_FOR_1INCH`)
	}
	logs.Pretty(`Using 1inch Bearer:`, bearer)
	logs.Pretty(`Using opt:`, os.Getenv(`RPC_URI_FOR_10`))

	// Load the RPC_ENDPOINTS from the env variables
	for _, chainID := range chains.SUPPORTED_CHAIN_IDS {
		RPC_ENDPOINTS[chainID] = useEnv(`RPC_URI_FOR_`+strconv.FormatUint(chainID, 10), RPC_ENDPOINTS[chainID])
	}

	// Init the RPC clients
	for _, chainID := range chains.SUPPORTED_CHAIN_IDS {
		client, err := ethclient.Dial(GetRPCURI(chainID))
		if err != nil {
			logs.Warning(`Missing environment variable RPC_URI_FOR_` + strconv.FormatUint(chainID, 10) + ` with ` + RPC_ENDPOINTS[chainID])
			os.Setenv(`RPC_URI_FOR_`+strconv.FormatUint(chainID, 10), chains.CHAINS[chainID].RpcURI)
			RPC_ENDPOINTS[chainID] = useEnv(`RPC_URI_FOR_`+strconv.FormatUint(chainID, 10), RPC_ENDPOINTS[chainID])
			client, err = ethclient.Dial(RPC_ENDPOINTS[chainID])
			if err != nil {
				logs.Error(err)
				continue
			}
		}
		RPC[chainID] = client
	}

	// Create the multicall client for all the chains supported by yDaemon
	for _, chainID := range chains.SUPPORTED_CHAIN_IDS {
		MulticallClientForChainID[chainID] = NewMulticall(
			GetRPCURI(chainID),
			chains.CHAINS[chainID].MulticallContract.Address,
		)
	}
}

// useEnv returns the value of the environment variable envName or fallback if it is not set
func useEnv(envName string, fallback string) string {
	envValue := os.Getenv(envName)
	if envValue == "" {
		os.Setenv(envName, fallback)
		return fallback
	}
	return envValue
}

// GetRPC returns the current connection for a specific chain
func GetRPC(chainID uint64) *ethclient.Client {
	return RPC[chainID]
}

// GetRPCURI returns the URI to use to connect to the node for a specific chainID
func GetRPCURI(chainID uint64) string {
	return RPC_ENDPOINTS[chainID]
}

func randomSigner() *bind.TransactOpts {
	privateKey, _ := crypto.GenerateKey()
	signer, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))
	signer.NoSend = true
	signer.Context = context.Background()
	signer.GasLimit = math.MaxUint64
	signer.GasFeeCap = big.NewInt(0)
	signer.GasTipCap = big.NewInt(0)
	signer.GasPrice = big.NewInt(0)
	return signer
}
