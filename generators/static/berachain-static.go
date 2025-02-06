package static

// This list has been manually curated by the Smol team and last updated on 06/02/2025
var BERACHAIN_NOTICE = `
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

var BERACHAIN_STATIC_TOKENLIST = map[uint64][]TStaticElement{
	80094: {
		{`0x0555e30da8f98308edb960aa94c0db47230d2b9c`, `https://assets.smold.app/api/token/80094/0x0555e30da8f98308edb960aa94c0db47230d2b9c/logo-128.png`},
		{`0x2f6f07cdcf3588944bf4c42ac74ff24bf56e7590`, `https://assets.smold.app/api/token/80094/0x2f6f07cdcf3588944bf4c42ac74ff24bf56e7590/logo-128.png`},
		{`0x549943e04f40284185054145c6e4e9568c1d3241`, `https://assets.smold.app/api/token/80094/0x549943e04f40284185054145c6e4e9568c1d3241/logo-128.png`},
		{`0x688e72142674041f8f6af4c808a4045ca1d6ac82`, `https://assets.smold.app/api/token/80094/0x688e72142674041f8f6af4c808a4045ca1d6ac82/logo-128.png`},
		{`0x6969696969696969696969696969696969696969`, `https://assets.smold.app/api/token/80094/0x6969696969696969696969696969696969696969/logo-128.png`},
		{`0xfcbd14dc51f0a4d49d5e53c2e0950e0bc26d0dce`, `https://assets.smold.app/api/token/80094/0xfcbd14dc51f0a4d49d5e53c2e0950e0bc26d0dce/logo-128.png`},
	},
}
