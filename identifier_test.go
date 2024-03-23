// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidIdentifier(t *testing.T) {
	assert.True(t, IsValidIdentifier("", true))
	assert.False(t, IsValidIdentifier("", false))

	valid := []string{
		"idA099",
		"_helloA97C0",
		"A97_helloC0",
	}
	for _, v := range valid {
		assert.True(t, IsValidIdentifier(v, false))
		assert.True(t, IsValidIdentifier(v, true))
	}

	inValid := []string{
		"   ",
		"\n",
		"\t ",
		"\r",
		"099idA",
		"$_helloA97C0",
		"#elloC0",
		"@lloC0",
		"--abc",
		"=yy",
		"+id",
	}
	for _, v := range inValid {
		assert.False(t, IsValidIdentifier(v, false))
		assert.False(t, IsValidIdentifier(v, true))
	}
}
