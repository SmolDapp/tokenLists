package helpers

import "github.com/ethereum/go-ethereum/common"

var IGNORED_TOKENS = map[uint64][]common.Address{
	1: {
		common.HexToAddress("0x84119cb33E8F590D75c2D6Ea4e6B0741a7494EDA"),
	},
}

func IsIgnoredToken(chainId uint64, address common.Address) bool {
	if address.Hex() == common.HexToAddress("0x0000000000000000000000000000000000000000").Hex() {
		return true
	}
	for _, ignoredToken := range IGNORED_TOKENS[chainId] {
		if ignoredToken.Hex() == address.Hex() {
			return true
		}
	}
	return false
}
