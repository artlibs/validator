// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestBlank(t *testing.T) {
	allBlanks := []string{
		"\t", "     ", "\n \r",
	}
	for _, value := range allBlanks {
		tv, ok := testBlank(value)
		assert.True(t, ok)
		assert.Equal(t, 0, len(tv))
	}

	nonBlanks := map[string]int{
		"\n go \r":   2,
		"it is":      5,
		"not\t":      3,
		"  blank   ": 5,
	}
	for k, n := range nonBlanks {
		nk, ok := testBlank(k)
		assert.False(t, ok)
		assert.Equal(t, n, len(nk))
	}
}
