

// This list has been manually curated by the Smol team and last updated on 29/01/2025
var SOPHON_NOTICE = `
------------------------------------------------------------------------------------------------------------------------
- PLEASE UPDATE TIMESTAMP ABOVE
------------------------------------------------------------------------------------------------------------------------
- ALWAYS ADD NEW TOKENS TO THE END OF THE LIST
- THE FIRST ELEMENT IN THE OBJECT IS THE TOKEN ADDRESS.
- THE SECOND ELEMENT IN THE OBJECT IS THE LOGO URI. IF THE LOGO URI IS NOT AVAILABLE, PLEASE LEAVE IT EMPTY WITH AN EMPTY STRING
- THE NAME, SYMBOL, DECIMALS WILL BE FETCHED FROM THE BLOCKCHAIN
- TO ADD A NEW NETWORK, PLEASE ADD A NEW MAP WITH THE NETWORK ID AS THE KEY AND THE LIST OF TOKENS AS THE VALUE
- /!!!!\ THE TOKEN MUST BE NON-REBASING /!!!!\ 
------------------------------------------------------------------------------------------------------------------------
`

var SOPHON_STATIC_TOKENLIST = map[uint64][]TStaticElement{
	50104: {
		{`0x72af9F169B619D85A47Dfa8fefbCD39dE55c567D`, `https://assets.smold.app/api/token/50104/0x72af9F169B619D85A47Dfa8fefbCD39dE55c567D/logo-128.png`} //ETH
		{`0x7A94dFb534ec790277Ad332EFB9B11687A130880`, `https://assets.smold.app/api/token/50104/0x7A94dFb534ec790277Ad332EFB9B11687A130880/logo-128.png`} //BEAM
		{`0x351ca00C5B64C03CA6d1eB9ABe5355ed449ac19C`, `https://assets.smold.app/api/token/50104/0x351ca00C5B64C03CA6d1eB9ABe5355ed449ac19C/logo-128.png`} //stZENT
		{`0x42E5bFdC7032D1F8e4B4D3fea3Eb5dB9C01Ef997`, `https://assets.smold.app/api/token/50104/0x42E5bFdC7032D1F8e4B4D3fea3Eb5dB9C01Ef997/logo-128.png`} //stATH
		{`0x06432e41131B70708c1C46814c2fcE9e720edd8A`, `https://assets.smold.app/api/token/50104/0x06432e41131B70708c1C46814c2fcE9e720edd8A/logo-128.png`} //PEPE
		{`0x6386dA73545ae4E2B2E0393688fA8B65Bb9a7169`, `https://assets.smold.app/api/token/50104/0x6386dA73545ae4E2B2E0393688fA8B65Bb9a7169/logo-128.png`} //USDT
		{`0x9Aa0F72392B5784Ad86c6f3E899bCc053D00Db4F`, `https://assets.smold.app/api/token/50104/0x9Aa0F72392B5784Ad86c6f3E899bCc053D00Db4F/logo-128.png`} //USDC
		{`0x60D02F185F80644e1A5ae35497736dd31d1b078B`, `https://assets.smold.app/api/token/50104/0x60D02F185F80644e1A5ae35497736dd31d1b078B/logo-128.png`} //wstETH
		{`0x5E9Fc50b44988B66BA84500f8Bc32C0493fe8F8D`, `https://assets.smold.app/api/token/50104/0x5E9Fc50b44988B66BA84500f8Bc32C0493fe8F8D/logo-128.png`} //weETH
		{`0xeCcbb9360d235548473Cb8C752735F68e652439B`, `https://assets.smold.app/api/token/50104/0xeCcbb9360d235548473Cb8C752735F68e652439B/logo-128.png`} //sDAI
		{`0x88171a5BbAcd92ca5e25575c5904581C80B025Dd`, `https://assets.smold.app/api/token/50104/0x88171a5BbAcd92ca5e25575c5904581C80B025Dd/logo-128.png`} //DAI
		{`0xF1f9E08a0818594FDe4713AE0Db1E46672Ca960E`, `https://assets.smold.app/api/token/50104/0xF1f9E08a0818594FDe4713AE0Db1E46672Ca960E/logo-128.png`} //wBTC
		{`0x841bE23f3872309f50E0c955c8Cb4b660f733811`, `https://assets.smold.app/api/token/50104/0x841bE23f3872309f50E0c955c8Cb4b660f733811/logo-128.png`} //stAZUR
		{`0xB2253e3D9A00e1471Ae7A3F77cBe7d81ec542A3e`, `https://assets.smold.app/api/token/50104/0xB2253e3D9A00e1471Ae7A3F77cBe7d81ec542A3e/logo-128.png`} //stAVAIL
	},
}