package ethereum

import (
	"encoding/hex"

	"github.com/migratooor/tokenLists/generators/common/logs"
)

// DecodeString decodes a string from a slice of interfaces
func DecodeString(something []interface{}, fallback string) string {
	if len(something) == 0 {
		return fallback
	}
	asString, ok := something[0].(string)
	if !ok {
		logs.Error(`DecodeString: unexpected type for decoded value`)
		return fallback
	}
	return asString
}

// DecodeHex decodes a hax from a slice of interfaces and try to convert it to a string
func DecodeHex(something []interface{}, fallback string) string {
	if len(something) == 0 {
		return fallback
	}
	asBytes32, ok := something[0].([32]uint8)
	if !ok {
		logs.Error(`DecodeHex: unexpected type for decoded value`)
		return fallback
	}
	if len(asBytes32) == 0 {
		return fallback
	}
	asBytes := make([]byte, 32)
	for i := 0; i < 32; i++ {
		if asBytes32[i] == 0 {
			asBytes = asBytes[:i]
			break
		}
		asBytes[i] = asBytes32[i]
	}
	asHex := hex.EncodeToString(asBytes[:])
	asString, err := hex.DecodeString(asHex)
	if err != nil {
		return fallback
	}
	return string(asString)
}

// DecodeUint64 decodes a uint64 from a slice of interfaces
func DecodeUint64(something []interface{}, fallback uint64) uint64 {
	if len(something) == 0 {
		return fallback
	}
	asUint8, ok := something[0].(uint8)
	if !ok {
		logs.Error(`DecodeUint64: unexpected type for decoded value`)
		return fallback
	}
	return uint64(asUint8)
}
