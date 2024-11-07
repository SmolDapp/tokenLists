package solana

import (
	"context"
	"strings"

	bin "github.com/gagliardetto/binary"
	solTokenMetadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	solana "github.com/gagliardetto/solana-go"
	solToken "github.com/gagliardetto/solana-go/programs/token"
	solRPC "github.com/gagliardetto/solana-go/rpc"
	"github.com/migratooor/tokenLists/generators/common/ethereum"
	"github.com/migratooor/tokenLists/generators/common/logs"
	"github.com/migratooor/tokenLists/generators/common/utils"
)

type TSolanaMetadata struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

func FetchBasicInformations(chainID uint64, tokens []string) map[string]*ethereum.TERC20 {
	tokenList := map[string]*ethereum.TERC20{}
	client := ethereum.GetRPC(chainID).(*solRPC.Client)
	solPublicKeys := []solana.PublicKey{}
	for _, address := range tokens {
		solanaPublicKey, err := solana.PublicKeyFromBase58(address)
		if err != nil {
			logs.Error(`[SOLANA] - Error fetching token:`, err)
			continue
		}
		solPublicKeys = append(solPublicKeys, solanaPublicKey)
	}

	/**********************************************************************************************
	** Fetch token mints in batches of 100 accounts at a time to avoid hitting RPC limits.
	** For each batch:
	** 1. Get multiple accounts data from RPC
	** 2. Decode the binary data into Mint structs
	** 3. Store in mintMap keyed by public key string
	**********************************************************************************************/
	mintMap := map[string]solToken.Mint{}
	for i := 0; i < len(solPublicKeys); i += 100 {
		end := i + 100
		if end > len(solPublicKeys) {
			end = len(solPublicKeys)
		}
		out, err := client.GetMultipleAccounts(context.TODO(), solPublicKeys[i:end]...)
		if err != nil {
			logs.Error(`[SOLANA] - Error fetching token mint:`, err)
			continue
		}
		for index, account := range out.Value {
			key := solPublicKeys[i+index].String()
			var mint solToken.Mint
			if err = bin.NewBinDecoder(account.Data.GetBinary()).Decode(&mint); err != nil {
				logs.Error(`[SOLANA] - Error decoding token mint:`, err)
				continue
			}
			mintMap[key] = mint
		}
	}

	for address, mint := range mintMap {
		var meta solTokenMetadata.Metadata

		/**************************************************************************************
		** Fetch the onChain token informations
		**************************************************************************************/
		tokenAddress := solana.MustPublicKeyFromBase58(address)
		metaAddress, _, err := solana.FindTokenMetadataAddress(tokenAddress)
		if err != nil {
			continue
		}
		resp, err := client.GetAccountInfo(context.TODO(), metaAddress)
		if err != nil {
			continue
		}
		d := bin.NewBorshDecoder(resp.Value.Data.GetBinary())
		if err = d.Decode(&meta); err != nil {
			continue
		}
		name := strings.TrimRight(meta.Data.Name, "\x00")
		symbol := strings.TrimRight(meta.Data.Symbol, "\x00")
		tokenJsonMetadata := strings.TrimRight(meta.Data.Uri, "\x00")
		logoURI := ``
		if tokenJsonMetadata != `` {
			validImageExtensions := []string{".png", ".jpg", ".jpeg", ".gif", ".svg"}
			isDirectImageURL := false
			for _, ext := range validImageExtensions {
				if strings.HasSuffix(tokenJsonMetadata, ext) {
					logoURI = tokenJsonMetadata
					isDirectImageURL = true
					break
				}
			}
			if !isDirectImageURL {
				jsonMetadata := fetchJSON[TSolanaMetadata](tokenJsonMetadata)
				if jsonMetadata.Image != `` {
					logoURI = jsonMetadata.Image
				}
			}
		}

		erc20FromChain := ethereum.TERC20{
			Address:  address,
			Name:     name,
			Symbol:   symbol,
			Decimals: uint64(mint.Decimals),
			ChainID:  chainID,
			LogoURI:  logoURI,
		}
		tokenList[utils.ToAddress(address)] = &erc20FromChain
	}

	return tokenList
}
