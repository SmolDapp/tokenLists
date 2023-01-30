package main

// TokenListToken is the token struct used in the default token list
type TokenListToken struct {
	ChainID  int    `json:"chainId"`
	Decimals int    `json:"decimals"`
	Count    int    `json:"count,omitempty"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	LogoURI  string `json:"logoURI"`
}

// TokenListData is the token list struct used in the default token list
type TokenListData struct {
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
	Version   struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	} `json:"version"`
	Tags     map[string]interface{} `json:"tags"`
	LogoURI  string                 `json:"logoURI"`
	Keywords []string               `json:"keywords"`
	Tokens   []TokenListToken       `json:"tokens"`
}
