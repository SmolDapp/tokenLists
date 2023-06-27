package helpers

var SUPPORTED_CHAIN_IDS = map[uint64]bool{
	1:     true, // Ethereum
	10:    true, // Optimism
	56:    true, // Binance Smart Chain
	100:   true, // xDai/Gnosis
	137:   true, // Polygon
	250:   true, // Fantom
	1101:  true, // Polygon ZKEVM
	42161: true, // Arbitrum
	43114: true, // Avalanche
}

// IsChainIDSupported returns true if the chainID is supported by our program
func IsChainIDSupported(chainID uint64) bool {
	return SUPPORTED_CHAIN_IDS[chainID]
}

var IGNORED_CHAIN_IDS = map[uint64]bool{
	3:           true, // Ropsten
	4:           true, // Rinkeby
	5:           true, // Goerli
	40:          true, // Telos
	42:          true, // Kovan
	65:          true, // OKExChain
	66:          true, // OKxChain
	69:          true, // Optimism Kovan
	97:          true, // Binance Smart Chain Testnet
	122:         true, // Fuse
	128:         true, // Heco
	199:         true, // Bitorrent
	256:         true, // Heco Testnet
	288:         true, // Boba
	420:         true, // Optimism Goerli
	1088:        true, // Metis Andromeda
	1284:        true, // Moonbeam
	1285:        true, // Moonriver
	1287:        true, // Moonbase
	2222:        true, // Kava
	4002:        true, // Fantom Testnet
	42170:       true, // Arbitrum Nova
	42220:       true, // Celo
	43113:       true, // Avalanche Fuji
	43288:       true, // Boba Avax
	56288:       true, // Boba BNB
	80001:       true, // Mumbai
	1666600000:  true, // Harmony Mainnet shard 0
	1666700000:  true, // Harmony Testnet shard 0
	11297108109: true, // Palm
}

// IsChainIDIgnored returns true if the chainID is ignored by our program
func IsChainIDIgnored(chainID uint64) bool {
	return IGNORED_CHAIN_IDS[chainID]
}
