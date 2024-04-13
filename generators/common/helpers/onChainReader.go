package helpers

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/logs"
	"github.com/migratooor/tokenLists/generators/common/models"
)

/**************************************************************************************************
* The RetrieveBasicInformations function reads the token list and returns a list of tokens with
* their basic informations (name, symbol, logoURI, decimals, chainID). These informations are
* retrieved from an on-chain reader.
*************************************************************************************************/
func RetrieveBasicInformations(chainID uint64, addresses []common.Address) map[string]*ethereum.TERC20 {
	erc20Map := make(map[string]*ethereum.TERC20)
	missingAddresses := []common.Address{}

	if !chains.IsChainIDSupported(chainID) {
		return erc20Map
	}

	for _, v := range addresses {
		if token, ok := ALL_EXISTING_TOKENS[chainID][v.Hex()]; ok {
			if token.Name == `` && token.Symbol == `` && token.Decimals == 0 {
				logs.Warning(`[ALL_EXISTING_TOKENS]: Missing name, symbol and decimals for token:`, token.Address, `on chain:`, chainID)
			} else if token.Name == `` && token.Symbol == `` {
				logs.Warning(`[ALL_EXISTING_TOKENS]: Missing name and symbol for token:`, token.Address, `on chain:`, chainID)
			} else if token.Name == `` && token.Decimals == 0 {
				logs.Warning(`[ALL_EXISTING_TOKENS]: Missing name and decimals for token:`, token.Address, `on chain:`, chainID)
			} else if token.Symbol == `` && token.Decimals == 0 {
				logs.Warning(`[ALL_EXISTING_TOKENS]: Missing symbol and decimals for token:`, token.Address, `on chain:`, chainID)
			} else if token.Name == `` {
				logs.Warning(`[ALL_EXISTING_TOKENS]: Missing name for token:`, token.Address, `on chain:`, chainID)
			} else if token.Symbol == `` {
				logs.Warning(`[ALL_EXISTING_TOKENS]: Missing symbol for token:`, token.Address, `on chain:`, chainID)
			} else if token.Decimals == 0 {
				logs.Warning(`[ALL_EXISTING_TOKENS]: Missing decimals for token:`, token.Address, `on chain:`, chainID)
			}
			erc20Map[v.Hex()] = &ethereum.TERC20{
				Address:  v,
				Name:     token.Name,
				Symbol:   token.Symbol,
				Decimals: uint64(token.Decimals),
				ChainID:  chainID,
			}
		} else {
			missingAddresses = append(missingAddresses, v)
		}
	}
	erc20FromChain := ethereum.FetchBasicInformations(chainID, missingAddresses)
	for k, v := range erc20FromChain {
		erc20Map[k] = v
		if _, ok := ALL_EXISTING_TOKENS[chainID]; !ok {
			ALL_EXISTING_TOKENS[chainID] = make(map[string]models.TokenListToken)
		}
		if v.Name == `` && v.Symbol == `` {
			logs.Warning(`[FETCHED_TOKEN] - Missing name and symbol for token:`, v.Address, `on chain:`, chainID)
		} else if v.Name == `` {
			logs.Warning(`[FETCHED_TOKEN] - Missing name for token:`, v.Address, `on chain:`, chainID)
		} else if v.Symbol == `` {
			logs.Warning(`[FETCHED_TOKEN] - Missing symbol for token:`, v.Address, `on chain:`, chainID)
		}
		ALL_EXISTING_TOKENS[chainID][v.Address.Hex()] = models.TokenListToken{
			Address:    v.Address.Hex(),
			Name:       v.Name,
			Symbol:     v.Symbol,
			LogoURI:    ``,
			Decimals:   int(v.Decimals),
			ChainID:    chainID,
			Occurrence: 1,
		}
	}
	return erc20Map
}

/**************************************************************************************************
 * The groupByChainID function is a small helper function to group tokens by chainID
 *************************************************************************************************/
func GroupByChainID(tokens []models.TokenListToken) map[uint64][]common.Address {
	tokensPerChainID := make(map[uint64][]common.Address)
	for _, token := range tokens {
		if !chains.IsChainIDSupported(token.ChainID) {
			continue
		}
		tokensPerChainID[token.ChainID] = append(tokensPerChainID[token.ChainID], common.HexToAddress(token.Address))
	}
	return tokensPerChainID
}

/**************************************************************************************************
 * The getExistingLogo function is a small helper function to get the existing logo of a token
 *************************************************************************************************/
func getExistingLogo(chainID uint64, lookingFor common.Address, slice []models.TokenListToken) string {
	for _, token := range slice {
		if !chains.IsChainIDSupported(token.ChainID) {
			continue
		}
		if common.HexToAddress(token.Address).Hex() == lookingFor.Hex() && chainID == token.ChainID {
			return token.LogoURI
		}
	}
	return ``
}

/**************************************************************************************************
 * The GetTokensFromList function reads the token list and returns a list of tokens with their
 * basic informations (name, symbol, logoURI, decimals, chainID). These informations are retrieved
 * from an on-chain reader.
 *************************************************************************************************/
func GetTokensFromList(tokensFromList []models.TokenListToken) []models.TokenListToken {
	tokens := []models.TokenListToken{}
	grouped := GroupByChainID(tokensFromList)

	for chainID, tokensForChain := range grouped {
		if !chains.IsChainIDSupported(chainID) {
			continue
		}

		tokensForChain = append(tokensForChain, chains.CHAINS[chainID].ExtraTokens...)
		tokensInfo := RetrieveBasicInformations(chainID, tokensForChain)
		for _, existingToken := range tokensForChain {
			if token, ok := tokensInfo[existingToken.Hex()]; ok {
				if newToken, err := SetToken(
					token.Address,
					token.Name,
					token.Symbol,
					getExistingLogo(chainID, existingToken, tokensFromList),
					chainID,
					int(token.Decimals),
				); err == nil {
					tokens = append(tokens, newToken)
				}
			}
		}
	}

	return tokens
}

/**************************************************************************************************
** The GetTokensFromAddresses function reads the addresses of the tokens and returns a list of
** tokens with their basic informations (name, symbol, logoURI, decimals, chainID). These
** informations are retrieved from an on-chain reader.
*************************************************************************************************/
func GetTokensFromAddresses(chainID uint64, tokenAddresses []common.Address) []models.TokenListToken {
	tokenList := []models.TokenListToken{}
	tokenAddresses = append(tokenAddresses, chains.CHAINS[chainID].ExtraTokens...)
	tokensInfo := RetrieveBasicInformations(chainID, tokenAddresses)
	for _, address := range tokenAddresses {
		if token, ok := tokensInfo[address.Hex()]; ok {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				`https://assets.smold.app/api/token/`+strconv.FormatUint(chainID, 10)+`/`+token.Address.Hex()+`/logo-128.png`,
				chainID,
				int(token.Decimals),
			); err == nil {
				tokenList = append(tokenList, newToken)
			}
		}
	}
	return tokenList
}

func GetTokensFromAddressesWithIcons(
	chainID uint64,
	tokenAddresses []common.Address,
	tokenIcons map[string]string,
) []models.TokenListToken {
	tokenList := []models.TokenListToken{}
	tokenAddresses = append(tokenAddresses, chains.CHAINS[chainID].ExtraTokens...)
	tokensInfo := RetrieveBasicInformations(chainID, tokenAddresses)

	for _, address := range tokenAddresses {
		if token, ok := tokensInfo[address.Hex()]; ok {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				tokenIcons[address.Hex()],
				chainID,
				int(token.Decimals),
			); err == nil {
				tokenList = append(tokenList, newToken)
			}
		}
	}
	return tokenList
}
