// @Since 2024-03-24.
// @Author Fury, All rights Reserved.

package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	assert.True(t, IsValidEmail("", true))
	assert.False(t, IsValidEmail("", false))

	validEmails := []string{
		"abc@example.org",
		"abc.123@example.org",
		"abc#we_rfc@example.com",
		"abc:def@-example.cc",
	}
	for _, email := range validEmails {
		assert.True(t, IsValidEmail(email, true))
		assert.True(t, IsValidEmail(email, false))
	}

	invalidEmails := []string{
		"abc.123#exaple.org",
		"ab#werfc@exaple",
		"abc_asdf90@abcd",
		"ab$example.org",
		"abc:def@-exa#mple.cc",
	}
	for _, email := range invalidEmails {
		assert.False(t, IsValidEmail(email, true))
		assert.False(t, IsValidEmail(email, false))
	}
}
