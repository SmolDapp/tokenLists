package helpers

import (
	"sync"

	"github.com/migratooor/tokenLists/generators/common/models"
)

func InitSyncMap[T any](chainIDs map[uint64]T) *sync.Map {
	tokensForChainIDSyncMap := sync.Map{}
	for chainID := range chainIDs {
		tokensForChainIDSyncMap.Store(chainID, []models.TokenListToken{})
	}
	return &tokensForChainIDSyncMap
}

func ExtractSyncMap(mapper *sync.Map) []models.TokenListToken {
	tokenList := []models.TokenListToken{}
	mapper.Range(func(chainID, syncMapRaw interface{}) bool {
		syncMap, _ := syncMapRaw.([]models.TokenListToken)
		tokenList = append(tokenList, syncMap...)
		return true
	})
	return tokenList
}
