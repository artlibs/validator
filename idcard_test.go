// @Since 2024-03-24.
// @Author Fury, All rights Reserved.

package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidIdCard(t *testing.T) {
	assert.True(t, IsValidIdCard("", true))
	assert.False(t, IsValidIdCard("", false))

	validIdCards := []string{
		"510212197002066715", // zkh
		"612321198301063313", // zkk
	}
	for _, id := range validIdCards {
		assert.True(t, IsValidIdCard(id, true))
		assert.True(t, IsValidIdCard(id, false))
	}

	invalidIdCards := []string{
		"510212197002066711",
		"612321198301063311",
	}
	for _, id := range invalidIdCards {
		assert.False(t, IsValidIdCard(id, true))
		assert.False(t, IsValidIdCard(id, false))
	}
}
