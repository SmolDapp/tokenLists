# Supported Chains

This document outlines the chains supported by our project, how to check for supported chains, and the process for adding new chains.

## Currently Supported Chains

Our project supports the following chains:

- Ethereum (ChainID: 1)
- Goerli (ChainID: 5)
- Optimism (ChainID: 10)
- Binance Smart Chain (ChainID: 56)
- Gnosis (ChainID: 100)
- Polygon (ChainID: 137)
- Fantom (ChainID: 250)
- Filecoin (ChainID: 314)
- zkSync (ChainID: 324)
- Metis (ChainID: 1088)
- Polygon zkEVM (ChainID: 1101)
- Mantle (ChainID: 5000)
- Base (ChainID: 8453)
- Arbitrum (ChainID: 42161)
- Celo (ChainID: 42220)
- Avalanche (ChainID: 43114)
- Linea (ChainID: 59144)
- Blast (ChainID: 81457)
- Scroll (ChainID: 534352)
- Zora (ChainID: 7777777)
- Rari (ChainID: 1380012617)
- Solana (ChainID: 1151111081099710)
- Ink (ChainID: 57073)
- Katana (ChainID: 747474)

## Adding a New Supported Chain

To add a new supported chain:

1. Define a new `TChain` struct for the chain in the appropriate file (e.g., `ethereum.go` for Ethereum-based chains).
2. Add the new chain to the `CHAINS` map in `config.go`.

Example of adding a new chain:

```go
var NEW_CHAIN = TChain{
    ID:            12345,
    Name:          "New Chain",
    LogoURI:       "https://example.com/new_chain_logo.png",
    RpcURI:        "https://rpc.newchain.example.com",
    MaxBlockRange: 1000,
    MaxBatchSize:  100,
    WeightRatio:   1.0,
    IsTestNet:     false,
    MulticallContract: TContractData{
        Address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
        Block:   1000000,
    },
    Coin: models.TokenListToken{
        Address:  "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
        ChainID:  12345,
        Name:     "New Chain Coin",
        Symbol:   "NCC",
        Decimals: 18,
        LogoURI:  "https://example.com/new_chain_coin_logo.png",
    },
    BlacklistedVaults: []string{},
    ExtraTokens:       []string{},
    IgnoredTokens:     []string{},
}

// In config.go
var CHAINS = map[uint64]TChain{
    // ... existing chains ...
    12345: NEW_CHAIN,
}
```

3. Some generators might need to be updated/added to support the new chain.

## TChain Struct Explanation

The `TChain` struct contains essential information about each supported chain:

```go
type TChain struct {
 ID uint64 // The unique identifier for the chain (ChainID)
 Name string // The human-readable name of the chain
 LogoURI string // URL to the chain's logo image
 RpcURI string // The RPC endpoint for interacting with the chain
 MaxBlockRange uint64 // Maximum number of blocks to query in a single request
 MaxBatchSize uint64 // Maximum number of items to include in a batch request
 WeightRatio float64 // A float value used for weighting or prioritizing the chain
 IsTestNet bool // Boolean indicating whether the chain is a testnet
 MulticallContract TContractData // Information about the Multicall contract on this chain
 Coin models.TokenListToken // Details about the native coin of the chain
 BlacklistedVaults []string // List of addresses for blacklisted vaults
 ExtraTokens []string // Additional token addresses to include
 IgnoredTokens []string // Token addresses to ignore during processing
}
```
