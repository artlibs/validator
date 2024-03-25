// @Since 2024-03-24.
// @Author Leo, All rights Reserved.

package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagMap(t *testing.T) {
	t.Helper()

	assert.True(t, TagMap["email"](
		"abc@example.org",
		true,
	))

	assert.True(t, TagMap["http_url"](
		"https://example.org",
		true,
	))

	assert.True(t, TagMap["id_card"](
		"510212197002066715",
		true,
	))

	assert.True(t, TagMap["identifier"](
		"idA099",
		true,
	))

	assert.True(t, TagMap["money"](
		"123,233.1",
		true,
	))

	assert.True(t, TagMap["phone_number"](
		"18157610011",
		true,
	))
}
