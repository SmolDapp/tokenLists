package helpers

import (
	"sort"

	"github.com/migratooor/tokenLists/generators/common/models"
)

type tokenSorter struct {
	tokens []models.TokenListToken
	by     func(p1, p2 *models.TokenListToken) bool
}

type By func(p1, p2 *models.TokenListToken) bool

func (by By) Sort(tokens []models.TokenListToken) {
	ps := &tokenSorter{
		tokens: tokens,
		by:     by,
	}
	sort.Sort(ps)
}

// Len is part of sort.Interface.
func (s *tokenSorter) Len() int {
	return len(s.tokens)
}

// Swap is part of sort.Interface.
func (s *tokenSorter) Swap(i, j int) {
	s.tokens[i], s.tokens[j] = s.tokens[j], s.tokens[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *tokenSorter) Less(i, j int) bool {
	return s.by(&s.tokens[i], &s.tokens[j])
}
