package helpers

import "github.com/ethereum/go-ethereum/common"

var IGNORED_TOKENS = map[uint64][]common.Address{
	1: {
		common.HexToAddress("0xdF5e0e81Dff6FAF3A7e52BA697820c5e32D806A8"),
		common.HexToAddress("0x84119cb33E8F590D75c2D6Ea4e6B0741a7494EDA"),
		common.HexToAddress("0x74C1E4b8caE59269ec1D85D3D4F324396048F4ac"),
		common.HexToAddress("0xdd94De9cFE063577051A5eb7465D08317d8808B6"),
		common.HexToAddress("0xb5A5F22694352C15B00323844aD545ABb2B11028"),
		common.HexToAddress("0xe3831c5A982B279A198456D577cfb90424cb6340"),
		common.HexToAddress("0x97deC872013f6B5fB443861090ad931542878126"),
		common.HexToAddress("0x4Dfd148B532e934a2a26eA65689cf6268753e130"),
		common.HexToAddress("0x009c80EfF4F5d8fcA2B961ee607B00B9C64eF9F2"),
		common.HexToAddress("0xEF68e7C694F40c8202821eDF525dE3782458639f"),
		common.HexToAddress("0x1Ca43a170BaD619322e6f54d46b57e504dB663aA"),
		common.HexToAddress("0xF03f8D65BaFA598611C3495124093c56e8F638f0"),
		common.HexToAddress("0x0DFCb45EAE071B3b846E220560Bbcdd958414d78"),
		common.HexToAddress("0xE4c94d45f7Aef7018a5D66f44aF780ec6023378e"),
		common.HexToAddress("0xbD168CbF9d3a375B38dC51A202B5E8a4E52069Ed"),
		common.HexToAddress("0x7606267A4bfff2c5010c92924348C3e4221955f2"),
		common.HexToAddress("0x28c8d01FF633eA9Cd8fc6a451D7457889E698de6"),
		common.HexToAddress("0x90162f41886c0946D09999736f1C15c8a105A421"),
		common.HexToAddress("0xC66eA802717bFb9833400264Dd12c2bCeAa34a6d"),
		common.HexToAddress("0x55b9a11c2e8351b4Ffc7b11561148bfaC9977855"),
		common.HexToAddress("0xe9Cf7887b93150D4F2Da7dFc6D502B216438F244"),
		common.HexToAddress("0xa78775bba7a542F291e5ef7f13C6204E704A90Ba"),
		common.HexToAddress("0x6aEDbF8dFF31437220dF351950Ba2a3362168d1b"),
		common.HexToAddress("0x2C4Bd064b998838076fa341A83d007FC2FA50957"),
		common.HexToAddress("0x66eb65D7Ab8e9567ba0fa6E37c305956c5341574"),
		common.HexToAddress("0x3045d1A840364c3657b8Df6c6F86a4359c23472B"),
		common.HexToAddress("0x1C5b760F133220855340003B43cC9113EC494823"),
		common.HexToAddress("0xeDBaF3c5100302dCddA53269322f3730b1F0416d"),
		common.HexToAddress("0xE5bA47fD94CB645ba4119222e34fB33F59C7CD90"),
		common.HexToAddress("0x585d3dFfD59b5201979363cC5432020D910DC40C"),
		common.HexToAddress("0x2a1530C4C41db0B0b2bB646CB5Eb1A67b7158667"),
		common.HexToAddress("0xF173214C720f58E03e194085B1DB28B50aCDeeaD"),
		common.HexToAddress("0x0Ba45A8b5d5575935B8158a88C631E9F9C95a2e5"),
		common.HexToAddress("0xBA96731324dE188ebC1eD87ca74544dDEbC07D7f"),
		common.HexToAddress("0x2859021eE7F2Cb10162E67F33Af2D22764B31aFf"),
		common.HexToAddress("0xEf6B4cE8C9Bc83744fbcdE2657b32eC18790458A"),
		common.HexToAddress("0xd6fDDe76B8C1C45B33790cc8751D5b88984c44ec"),
		common.HexToAddress("0x38c6A68304cdEfb9BEc48BbFaABA5C5B47818bb2"),
		common.HexToAddress("0xCc13Fc627EFfd6E35D2D2706Ea3C4D7396c610ea"),
		common.HexToAddress("0xF2c96E402c9199682d5dED26D3771c6B192c01af"),
		common.HexToAddress("0x3019BF2a2eF8040C242C9a4c5c4BD4C81678b2A1"),
		common.HexToAddress("0x12513335ffD5DAfc2334e98625d27c1CA84bff86"),
		common.HexToAddress("0x089A502032166e07Ae83eb434c16790cA2FA4661"),
		common.HexToAddress("0xbdEB4b83251Fb146687fa19D1C660F99411eefe3"),
		common.HexToAddress("0x5e3845A1d78DB544613EdbE43Dc1Ea497266d3b8"),
		common.HexToAddress("0x5c543e7AE0A1104f78406C340E9C64FD9fCE5170"),
		common.HexToAddress("0xE0B7927c4aF23765Cb51314A0E0521A9645F0E2A"),
		common.HexToAddress("0x1F7e8fe01AEbA6fDAEA85161746f4D53DC9bdA4F"),
		common.HexToAddress("0xB31C219959E06f9aFBeB36b388a4BaD13E802725"),
		common.HexToAddress("0xfD0fD32A20532ad690731c2685d77c351015ebBa"),
		common.HexToAddress("0xcaA7e4656f6A2B59f5f99c745F91AB26D1210DCe"),
		common.HexToAddress("0xd51e852630DeBC24E9e1041a03d80A0107F8Ef0C"),
		common.HexToAddress("0x86Fa049857E0209aa7D9e616F7eb3b3B78ECfdb0"),
		common.HexToAddress("0xec6432B90e7fD4d9f872cc5C781f05B617DB861E"),
		common.HexToAddress(`0x0E69D0A2bbB30aBcB7e5CfEA0E4FDe19C00A8d47`),
	},
	56: {
		common.HexToAddress("0xc00e94Cb662C3520282E6f5717214004A7f26888"),
	},
	137: {
		common.HexToAddress("0xec6432B90e7fD4d9f872cc5C781f05B617DB861E"),
	},
}

func IsIgnoredToken(chainId uint64, address common.Address) bool {
	if address.Hex() == common.HexToAddress("0x0000000000000000000000000000000000000000").Hex() {
		return true
	}
	for _, ignoredToken := range IGNORED_TOKENS[chainId] {
		if ignoredToken.Hex() == address.Hex() {
			return true
		}
	}
	return false
}
