package models

// TokenListToken is the token struct used in the default token list
type TokenListToken struct {
	Address  string                 `json:"address"`
	Name     string                 `json:"name"`
	Symbol   string                 `json:"symbol"`
	LogoURI  string                 `json:"logoURI,omitempty"`
	ChainID  uint64                 `json:"chainId"`
	Decimals int                    `json:"decimals"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The following fields are optional and not exported
	Occurrence int `json:"-"` // Use for aggregation: number of time this token was found
}

// TokenListData is the token list struct used in the default token list
// [T any](uri string) (data T) {
type TokenListData[T any] struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
	Version     struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	} `json:"version"`
	LogoURI           string                    `json:"logoURI"`
	Keywords          []string                  `json:"keywords"`
	Tokens            []T                       `json:"tokens"`
	PreviousTokensMap map[string]TokenListToken `json:"-"`
	NextTokensMap     map[string]TokenListToken `json:"-"`
	Metadata          map[string]interface{}    `json:"metadata,omitempty"`
}
