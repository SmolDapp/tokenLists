package main

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func handleZkSyncTokenList(chainID uint64, tokenAddresses []common.Address, imageURI []string) []TokenListToken {
	tokenList := []TokenListToken{}

	tokensInfo := retrieveBasicInformations(chainID, tokenAddresses)
	for index, address := range tokenAddresses {
		if token, ok := tokensInfo[address.Hex()]; ok {
			if newToken, err := SetToken(
				token.Address,
				token.Name,
				token.Symbol,
				imageURI[index],
				chainID,
				int(token.Decimals),
			); err == nil {
				tokenList = append(tokenList, newToken)
			}
		}
	}
	chainCoin := BASE_EXPLORERS_URI[chainID].Coin
	if chainCoin.Address == "" {
		chainCoin = ETHER
	}
	chainCoin.ChainID = chainID
	chainCoin.LogoURI = `https://assets.smold.app/api/token/` + strconv.FormatUint(chainID, 10) + `/` + chainCoin.Address + `/logo-128.png`
	tokenList = append(tokenList, chainCoin)
	return tokenList
}

func fetchZkSyncTokenList() []TokenListToken {
	imageURI := []string{
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/eth.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/usdc.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/mute.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/combo.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/perp.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/lusd.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/dvf.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/woo.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/xcrmrk.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/deri.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/dextf.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/govi.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/1inch.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/sis.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/lqty.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/pepe.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/reth.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/rpl.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/ufi.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/cbeth.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/ibex.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/raise.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/lsd.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/ethx.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/wbtc.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/knc.png?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/bel.png?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/zz.png?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/bitcoin.png?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/nbx.svg?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/wagmi.png?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/pool.png?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/wbtc.png?alt=media`,
		`https://firebasestorage.googleapis.com/v0/b/token-library.appspot.com/o/usdt.png?alt=media`,
	}
	tokens := []common.Address{
		common.HexToAddress(`0x000000000000000000000000000000000000800A`),
		common.HexToAddress(`0x3355df6D4c9C3035724Fd0e3914dE96A5a83aaf4`),
		common.HexToAddress(`0x0e97C7a0F8B2C9885C8ac9fC6136e829CbC21d42`),
		common.HexToAddress(`0xc2B13Bb90E33F1E191b8aA8F44Ce11534D5698E3`),
		common.HexToAddress(`0x42c1c56be243c250AB24D2ecdcC77F9cCAa59601`),
		common.HexToAddress(`0x503234F203fC7Eb888EEC8513210612a43Cf6115`),
		common.HexToAddress(`0xBbD1bA24d589C319C86519646817F2F153c9B716`),
		common.HexToAddress(`0x9E22D758629761FC5708c171d06c2faBB60B5159`),
		common.HexToAddress(`0x6F1A89C16a49549508a2b6D2ac6F34523AA2A545`),
		common.HexToAddress(`0x140D5bc5b62d6cB492B1A475127F50d531023803`),
		common.HexToAddress(`0x9929bCAC4417A21d7e6FC86F6Dae1Cc7f27A2e41`),
		common.HexToAddress(`0xD63eF5e9C628c8a0E8984CDfb7444AEE44B09044`),
		common.HexToAddress(`0x3f0B8B206A7FBdB3ecFc08c9407CA83F5aB1Ce59`),
		common.HexToAddress(`0xdd9f72afED3631a6C85b5369D84875e6c42f1827`),
		common.HexToAddress(`0xf755cF4f0887279a8BCBE5E39eE062a5B7188401`),
		common.HexToAddress(`0xFD282F16a64c6D304aC05d1A58Da15bed0467c71`),
		common.HexToAddress(`0x32Fd44bB869620C0EF993754c8a00Be67C464806`),
		common.HexToAddress(`0x1CF8553Da5a75C20cdC33532cb19Ef7E3bFFf5BC`),
		common.HexToAddress(`0xa0C1BC64364d39c7239bd0118b70039dBe5BbdAE`),
		common.HexToAddress(`0x75Af292c1c9a37b3EA2E6041168B4E48875b9ED5`),
		common.HexToAddress(`0xbe9f8C0d6f0Fd7e46CDaCCA340747EA2f247991D`),
		common.HexToAddress(`0x3D79F1e3f6AFd3F30EA450aFffb8632AED59B46f`),
		common.HexToAddress(`0x458A2E32eAbc7626187E6b75f29D7030a5202bD4`),
		common.HexToAddress(`0x668cc2668Eeeaf8075d38E72EF54fa546BF3C39c`),
		common.HexToAddress(`0xBBeB516fb02a01611cBBE0453Fe3c580D7281011`),
		common.HexToAddress(`0x6ee46Cb7cD2f15Ee1ec9534cf29a5b51C83283e6`),
		common.HexToAddress(`0xB83CFB285fc8D936E8647FA9b1cC641dBAae92D9`),
		common.HexToAddress(`0x1ab721f531Cab4c87d536bE8B985EAFCE17f0184`),
		common.HexToAddress(`0x26b7F317C440E57db2fb4b377A3f1b3BBF5463C7`),
		common.HexToAddress(`0x2d850F34E957BA3dcbEe47fc2c79ff78044fB12e`),
		common.HexToAddress(`0xD7C6210f3d6011D6B1BdDfA60440fe763340Df4c`),
		common.HexToAddress(`0x97003aC71CC4a096E06C73e753d9b84f0039A064`),
		common.HexToAddress(`0xBBeB516fb02a01611cBBE0453Fe3c580D7281011`),
		common.HexToAddress(`0x493257fD37EDB34451f62EDf8D2a0C418852bA4C`),
	}
	return handleZkSyncTokenList(324, tokens, imageURI)
}

func buildZkSyncTokenList() {
	tokenList := loadTokenListFromJsonFile(`zksync.json`)
	tokenList.Name = `zkSync`
	tokenList.LogoURI = ``
	tokenList.Keywords = []string{`zksync`}
	tokens := []TokenListToken{}
	tokens = append(tokens, fetchZkSyncTokenList()...)
	saveTokenListInJsonFile(tokenList, tokens, `zksync.json`, Standard)
}
