package helpers

import (
	"strconv"
	"strings"

	"github.com/migratooor/tokenLists/generators/common/chains"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/logs"
	"github.com/migratooor/tokenLists/generators/common/models"
	"github.com/migratooor/tokenLists/generators/common/solana"
	"github.com/migratooor/tokenLists/generators/common/utils"
)

type TSolanaMetadata struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

/**************************************************************************************************
* The RetrieveBasicInformations function reads the token list and returns a list of tokens with
* their basic informations (name, symbol, logoURI, decimals, chainID). These informations are
* retrieved from an on-chain reader.
*************************************************************************************************/
func RetrieveBasicInformations(chainID uint64, addresses []string) map[string]ethereum.TERC20 {
	erc20Map := make(map[string]ethereum.TERC20)
	missingAddresses := []string{}

	if !chains.IsChainIDSupported(chainID) {
		return erc20Map
	}

	if _, ok := ALL_EXISTING_TOKENS[chainID]; !ok {
		ALL_EXISTING_TOKENS[chainID] = make(map[string]models.TokenListToken)
	}

	for _, v := range addresses {
		if token, ok := ALL_EXISTING_TOKENS[chainID][utils.ToAddress(v)]; ok {
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
			erc20Map[utils.ToAddress(v)] = ethereum.TERC20{
				Address:  utils.ToAddress(v),
				Name:     token.Name,
				Symbol:   token.Symbol,
				Decimals: uint64(token.Decimals),
				ChainID:  chainID,
			}
		} else {
			missingAddresses = append(missingAddresses, v)
		}
	}

	/**********************************************************************************************
	** If the chain is an EVM chain, we will fetch the basic informations from the on-chain reader.
	**********************************************************************************************/
	if chains.CHAINS[chainID].Type == `EVM` {
		erc20FromChain := ethereum.FetchBasicInformations(chainID, missingAddresses)
		for k, v := range erc20FromChain {
			erc20Map[utils.ToAddress(k)] = *v
			if v.Name == `` && v.Symbol == `` {
				logs.Warning(`[FETCHED_TOKEN] - Missing name and symbol for token:`, v.Address, `on chain:`, chainID)
			} else if v.Name == `` {
				logs.Warning(`[FETCHED_TOKEN] - Missing name for token:`, v.Address, `on chain:`, chainID)
			} else if v.Symbol == `` {
				logs.Warning(`[FETCHED_TOKEN] - Missing symbol for token:`, v.Address, `on chain:`, chainID)
			}
			ALL_EXISTING_TOKENS[chainID][utils.ToAddress(v.Address)] = models.TokenListToken{
				Address:    utils.ToAddress(v.Address),
				Name:       v.Name,
				Symbol:     v.Symbol,
				LogoURI:    ``,
				Decimals:   int(v.Decimals),
				ChainID:    chainID,
				Occurrence: 1,
			}
		}
	}
	/**********************************************************************************************
	** If the chain is a SVM chain, we need to find another way to retrieve the basic informations.
	**********************************************************************************************/
	if chains.CHAINS[chainID].Type == `SVM` {
		erc20FromChain := solana.FetchBasicInformations(chainID, missingAddresses)

		for address, token := range erc20FromChain {
			erc20Map[utils.ToAddress(address)] = *token

			/**************************************************************************************
			** Add the token to the map
			**************************************************************************************/
			if token.Name == `` && token.Symbol == `` {
				logs.Warning(`[FETCHED_TOKEN] - Missing name and symbol for token:`, address, `on chain:`, chainID)
			} else if token.Name == `` {
				logs.Warning(`[FETCHED_TOKEN] - Missing name for token:`, address, `on chain:`, chainID)
			} else if token.Symbol == `` {
				logs.Warning(`[FETCHED_TOKEN] - Missing symbol for token:`, address, `on chain:`, chainID)
			}
			solanaToken := models.TokenListToken{
				Address:    utils.ToAddress(address),
				Name:       token.Name,
				Symbol:     token.Symbol,
				LogoURI:    token.LogoURI,
				Decimals:   int(token.Decimals),
				ChainID:    chainID,
				Occurrence: 1,
			}
			ALL_EXISTING_TOKENS[chainID][utils.ToAddress(address)] = solanaToken
		}
	}

	return erc20Map
}

/**************************************************************************************************
 * The groupByChainID function is a small helper function to group tokens by chainID
 *************************************************************************************************/
func GroupByChainID(tokens []models.TokenListToken) map[uint64][]string {
	tokensPerChainID := make(map[uint64][]string)
	for _, token := range tokens {
		if !chains.IsChainIDSupported(token.ChainID) {
			continue
		}
		tokensPerChainID[token.ChainID] = append(tokensPerChainID[token.ChainID], utils.ToAddress(token.Address))
	}
	return tokensPerChainID
}

/**************************************************************************************************
 * The getExistingLogo function is a small helper function to get the existing logo of a token
 *************************************************************************************************/
func getExistingLogo(chainID uint64, lookingFor string, slice []models.TokenListToken) string {
	if token, ok := ALL_EXISTING_TOKENS[chainID][utils.ToAddress(lookingFor)]; ok {
		if token.LogoURI != `` {
			return token.LogoURI
		}
	}

	for _, token := range slice {
		if !chains.IsChainIDSupported(token.ChainID) {
			continue
		}
		if strings.EqualFold(token.Address, lookingFor) && chainID == token.ChainID {
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
			if token, ok := tokensInfo[utils.ToAddress(existingToken)]; ok {
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
func GetTokensFromAddresses(chainID uint64, tokenAddresses []string) []models.TokenListToken {
	tokenList := []models.TokenListToken{}
	tokenAddresses = append(tokenAddresses, chains.CHAINS[chainID].ExtraTokens...)
	tokensInfo := RetrieveBasicInformations(chainID, tokenAddresses)
	for _, address := range tokenAddresses {
		if token, ok := tokensInfo[utils.ToAddress(address)]; ok {
			logoURI := getExistingLogo(chainID, address, tokenList)
			if logoURI == `` {
				logoURI = `https://assets.smold.app/api/token/` + strconv.FormatUint(chainID, 10) + `/` + token.Address + `/logo-128.png`
			}
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				logoURI,
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
	tokenAddresses []string,
	tokenIcons map[string]string,
) []models.TokenListToken {
	tokenList := []models.TokenListToken{}
	tokenAddresses = append(tokenAddresses, chains.CHAINS[chainID].ExtraTokens...)
	tokensInfo := RetrieveBasicInformations(chainID, tokenAddresses)

	for _, address := range tokenAddresses {
		if token, ok := tokensInfo[utils.ToAddress(address)]; ok {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				tokenIcons[utils.ToAddress(address)],
				chainID,
				int(token.Decimals),
			); err == nil {
				tokenList = append(tokenList, newToken)
			}
		}
	}
	return tokenList
}

func GetTokensFromAddressesWithIcons_NOREPLACEMENT(
	chainID uint64,
	tokenAddresses []string,
	tokenIcons map[string]string,
) []models.TokenListToken {
	tokenList := []models.TokenListToken{}
	tokenAddresses = append(tokenAddresses, chains.CHAINS[chainID].ExtraTokens...)
	tokensInfo := RetrieveBasicInformations(chainID, tokenAddresses)

	for _, address := range tokenAddresses {
		if token, ok := tokensInfo[utils.ToAddress(address)]; ok {
			if newToken, err := SetToken_NOREPLACEMENT(
				token.Address,
				token.Name,
				token.Symbol,
				tokenIcons[utils.ToAddress(address)],
				chainID,
				int(token.Decimals),
			); err == nil {
				tokenList = append(tokenList, newToken)
			}
		}
	}
	return tokenList
}
