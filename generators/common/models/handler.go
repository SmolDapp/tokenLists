package models

// InitTokenList initializes the token list
func InitTokenList() TokenListData[TokenListToken] {
	newTokenList := TokenListData[TokenListToken]{
		Name:      ``,
		Timestamp: ``,
		Keywords:  []string{},
		Tokens:    []TokenListToken{},
	}
	newTokenList.Version.Major = 0
	newTokenList.Version.Minor = 0
	newTokenList.Version.Patch = 0
	newTokenList.PreviousTokensMap = make(map[string]TokenListToken)
	newTokenList.NextTokensMap = make(map[string]TokenListToken)

	return newTokenList
}
