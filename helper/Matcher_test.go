package helper

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMatcherSuccess(t *testing.T) {
	m, err := NewMatcher("^zzz$", "(/TIMESHIFT/\\d+/)", "xxx")

	assert.Nil(t, err)
	assert.IsType(t, new(MatcherImpl), m)
}

func TestNewMatcherWrongRegExp(t *testing.T) {
	m, err := NewMatcher("^zzz$", "~(/TIMESHIFT/\\d+/~", "~xxx~")

	assert.NotNil(t, err)
	assert.Nil(t, m)
}

func TestMatchSimpleValue(t *testing.T) {
	m, _ := NewMatcher("^zzz$", "(/TIMESHIFT/\\d+/)", "xxx")

	assert.True(t, m.Match("zzz"))
	assert.True(t, m.Match("/tmp/TIMESHIFT/584724398/zzz"))
	assert.True(t, m.Match("xxxxxxxxxxxxxxxxx"))
	assert.False(t, m.Match("asdf"))
}

func TestMatchGetMatched(t *testing.T) {
	m, _ := NewMatcher("(/TIMESHIFT/\\d+/)")
	m.Match("/tmp/TIMESHIFT/584724398/zzz")

	matched, _ := m.GetMatched()

	assert.Equal(t, "/TIMESHIFT/584724398/", matched)
}

func TestMatchGetMatchedError(t *testing.T) {
	m, _ := NewMatcher("(/TIMESHIFT/\\d+/)")
	m.Match("/tmp/TIMESHIFT/hg584724398/zzz")

	_, err := m.GetMatched()

	assert.IsType(t, errors.New(""), err)
}
