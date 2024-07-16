package static

import "github.com/ethereum/go-ethereum/common"

// This list has been manually curated by the Filecoin team and last updated on 05/06/2024

var FILECOIN_NOTICE = `
------------------------------------------------------------------------------------------------------------------------
- PLEASE UPDATE TIMESTAMP ABOVE
------------------------------------------------------------------------------------------------------------------------
- ALWAYS ADD NEW TOKENS TO THE END OF THE LIST
- THE FIRST ELEMENT IN THE OBJECT IS THE TOKEN ADDRESS. PLEASE MAKE SURE TO KEEP THE common.hexToAddress PART
- THE SECOND ELEMENT IN THE OBJECT IS THE LOGO URI. IF THE LOGO URI IS NOT AVAILABLE, PLEASE LEAVE IT EMPTY WITH AN EMPTY STRING
- THE NAME, SYMBOL, DECIMALS WILL BE FETCHED FROM THE BLOCKCHAIN
- TO ADD A NEW NETWORK, PLEASE ADD A NEW MAP WITH THE NETWORK ID AS THE KEY AND THE LIST OF TOKENS AS THE VALUE
------------------------------------------------------------------------------------------------------------------------
`

var FILECOIN_STATIC = map[uint64][]TStaticElement{
	314: {
		{common.HexToAddress(`0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`), `https://assets.smold.app/api/token/314/0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE/logo-128.png`},
		{common.HexToAddress(`0x60E1773636CF5E4A227d9AC24F20fEca034ee25A`), ``},
		{common.HexToAddress(`0x690908f7fa93afC040CFbD9fE1dDd2C2668Aa0e0`), ``},
		{common.HexToAddress(`0xAaa93ac72bECfbBc9149f293466bbdAa4b5Ef68C`), ``},
		{common.HexToAddress(`0x57E3BB9F790185Cfe70Cc2C15Ed5d6B84dCf4aDb`), ``},
		{common.HexToAddress(`0xd0437765D1Dc0e2fA14E97d290F135eFdf1a8a9A`), ``},
		{common.HexToAddress(`0x5C7e299CF531eb66f2A1dF637d37AbB78e6200C7`), ``},
		{common.HexToAddress(`0xb829b68f57CC546dA7E5806A929e53bE32a4625D`), ``},
		{common.HexToAddress(`0x2421db204968A367CC2C866CD057fA754Cb84EdF`), ``},
		{common.HexToAddress(`0x522b61755b5ff8176b2931da7bf1a5f9414eb710`), ``},
		{common.HexToAddress(`0xEB466342C4d449BC9f53A865D5Cb90586f405215`), ``},
	},
}
