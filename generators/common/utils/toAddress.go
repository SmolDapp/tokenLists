package utils

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// ToAddress convert a string to a checksummed address
func ToAddress(value string) string {
	// If it's an Ethereum address, convert it to checksummed
	if strings.HasPrefix(value, `0x`) {
		return common.HexToAddress(value).Hex()
	}

	// Otherwise, we hope it's already a checksummed address
	return value
}
