// @Since 2024-03-23.
// @Author Fury, All rights Reserved.

package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidHttpURL(t *testing.T) {
	assert.True(t, IsValidHttpURL("", true))
	assert.False(t, IsValidHttpURL("", false))

	valid := []string{
		"http://192.168.202.123",
		"httpS://192.168.202.123:8080",
		"http://example.org",
		"http://example.org/",
		"htTp://examPLE.org?k=v",
		"http://example.org/?k=v",
		"http://example.org/sub/?k=v",
		"hTtp://example.org:80?k=v1&k=v2",
		"http://example.org/?k=v1&k=v2",
		"http://example.org:80/sub/?k=v1&k=v2",
		"httP://example.ORG/#/?k=v1&k=v2",
		"https://example.org/~sub/?k=v1&k=v2",
		"https://example.org/#/?k=v1&k=v2  ",
		"\t https://example.org  ",
		" https://user:234@example.org/?k=9  ",
		"http://10.21.32.43/web/#/core+/x-rule_",
	}
	for _, v := range valid {
		assert.True(t, IsValidHttpURL(v, false))
		assert.True(t, IsValidHttpURL(v, true))
	}

	inValid := []string{
		"  \n ",
		"=yy\r",
		"https://099idA",
		"http$//_helloA97C0",
		"ftp://#elloC0",
		"http//:23@lloC0.com",
		"https:/:--abc.net",
	}
	for _, v := range inValid {
		assert.False(t, IsValidHttpURL(v, false))
		assert.False(t, IsValidHttpURL(v, true))
	}
}
