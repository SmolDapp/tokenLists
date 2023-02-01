package helpers

var IGNORED_CHAINIDS = map[uint64]bool{
	3:           true, // Ropsten
	4:           true, // Rinkeby
	5:           true, // Goerli
	40:          true, // Telos
	42:          true, // Kovan
	65:          true, // OKExChain
	66:          true, // OKxChain
	97:          true, // Binance Smart Chain Testnet
	122:         true, // Fuse
	128:         true, // Heco
	199:         true, // Bitorrent
	256:         true, // Heco Testnet
	288:         true, // Boba
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

func IsIgnoredChain(chainId uint64) bool {
	return IGNORED_CHAINIDS[chainId]
}
