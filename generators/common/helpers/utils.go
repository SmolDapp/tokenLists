package helpers

import (
	"github.com/ethereum/go-ethereum/common"
)

// SafeString returns the provided variable or a fallback if it is empty
func SafeString(value string, fallback string) string {
	if value == `` {
		return fallback
	}
	return value
}

// SafeInt returns the provided variable or a fallback if it is empty
func SafeInt(value int, fallback int) int {
	if value == 0 {
		return fallback
	}
	return value
}

// ToAddress convert a string to a checksummed address
func ToAddress(value string) string {
	return common.HexToAddress(value).Hex()
}
