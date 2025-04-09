package helpers

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/logs"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/common/utils"
)

var ALL_EXISTING_TOKENS = map[uint64]map[string]models.TokenListToken{}

func init() {
	chainCoins := []models.TokenListToken{}

	for _, chainID := range chains.SUPPORTED_CHAIN_IDS {
		chain := chains.CHAINS[chainID]
		chainCoins = append(chainCoins, chain.Coin)
	}

	for _, coin := range chainCoins {
		if _, ok := ALL_EXISTING_TOKENS[coin.ChainID]; !ok {
			ALL_EXISTING_TOKENS[coin.ChainID] = make(map[string]models.TokenListToken)
		}
		if coin.Name == `` {
			logs.Warning(`Missing name for token:`, coin.Address, `on chain:`, coin.ChainID)
		}
		if coin.Symbol == `` {
			logs.Warning(`Missing symbol for token:`, coin.Address, `on chain:`, coin.ChainID)
		}
		ALL_EXISTING_TOKENS[coin.ChainID][coin.Address] = models.TokenListToken{
			Address:    common.HexToAddress(coin.Address).Hex(),
			Name:       coin.Name,
			Symbol:     coin.Symbol,
			LogoURI:    ``,
			Decimals:   int(coin.Decimals),
			ChainID:    coin.ChainID,
			Occurrence: 1,
		}
	}
}

func SetToken(
	address string,
	name string, symbol string, logoURI string,
	chainID uint64, decimals int,
) (models.TokenListToken, error) {
	token := models.TokenListToken{}
	if name == `` {
		return token, errors.New(`token name is empty`)
	}
	if symbol == `` {
		return token, errors.New(`token symbol is empty`)
	}
	if decimals == 0 {
		return token, errors.New(`token decimals is 0`)
	}
	if chains.IsTokenIgnored(chainID, address) {
		return token, errors.New(`token is ignored`)
	}
	if !chains.IsChainIDSupported(chainID) {
		return token, errors.New(`chainID is ignored`)
	}
	if strings.EqualFold(address, `0x2791bca1f2de4661ed88a30c99a7a9449aa84174`) && chainID == 137 {
		name = `Bridged USD Coin (PoS)`
		symbol = `USDC.e`
	}

	if chains.CHAINS[chainID].Type == `SVM` {
		token.Address = utils.ToAddress(address)
	} else {
		token.Address = common.HexToAddress(address).Hex()
	}
	token.ChainID = chainID
	token.Decimals = decimals
	token.Name = name
	token.Symbol = symbol
	token.LogoURI = UseIcon(
		chainID,
		token.Name+` - `+token.Symbol,
		token.Address,
		logoURI)
	return token, nil
}

func SetToken_NOREPLACEMENT(
	address string,
	name string, symbol string, logoURI string,
	chainID uint64, decimals int,
) (models.TokenListToken, error) {
	token := models.TokenListToken{}
	if name == `` {
		return token, errors.New(`token name is empty`)
	}
	if symbol == `` {
		return token, errors.New(`token symbol is empty`)
	}
	if decimals == 0 {
		return token, errors.New(`token decimals is 0`)
	}
	if chains.IsTokenIgnored(chainID, address) {
		return token, errors.New(`token is ignored`)
	}
	if !chains.IsChainIDSupported(chainID) {
		return token, errors.New(`chainID is ignored`)
	}
	if strings.EqualFold(address, `0x2791bca1f2de4661ed88a30c99a7a9449aa84174`) && chainID == 137 {
		name = `Bridged USD Coin (PoS)`
		symbol = `USDC.e`
	}

	if chains.CHAINS[chainID].Type == `SVM` {
		token.Address = utils.ToAddress(address)
	} else {
		token.Address = common.HexToAddress(address).Hex()
	}
	token.ChainID = chainID
	token.Decimals = decimals
	token.Name = name
	token.Symbol = symbol
	token.LogoURI = logoURI
	return token, nil
}
