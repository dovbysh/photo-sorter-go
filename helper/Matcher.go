package helper

import (
	"errors"
	"regexp"
)

// Matcher is regexp wraper
type Matcher interface {
	Match(s string) bool
	GetMatched() (string, error)
}

// MatcherImpl is regexp wraper
type MatcherImpl struct {
	rx      []*regexp.Regexp
	matched string
}

// NewMatcher is fabric method for MatcherImpl
func NewMatcher(rxStrings ...string) (*MatcherImpl, error) {
	m := new(MatcherImpl)

	for _, s := range rxStrings {
		r, err := regexp.Compile(s)

		if err != nil {
			return nil, err
		}

		m.rx = append(m.rx, r)
	}

	return m, nil
}

// Match returns true if any Regular Expression matched s
func (m *MatcherImpl) Match(s string) bool {
	for _, rx := range m.rx {
		m.matched = rx.FindString(s)

		if len(m.matched) != 0 {
			return true
		}
	}

	return false
}

// GetMatched returns last found string from Match or error
func (m *MatcherImpl) GetMatched() (string, error) {
	if len(m.matched) == 0 {
		return "", errors.New("Not found")
	}

	return m.matched, nil
}
