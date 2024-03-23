// @Since 2024-03-21.
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
		"http://10.34.56.78/web/#/core+/x-rule_",
	}
	for _, v := range valid {
		assert.True(t, IsValidHttpURL(v, false))
		assert.True(t, IsValidHttpURL(v, true))
	}

	inValid := []string{
		"   ",
		"\n",
		"\r",
		"=yy",
		"https://099idA",
		"http$//_helloA97C0",
		"ftp://#elloC0",
		"http:@lloC0",
		"https::--abc",
	}
	for _, v := range inValid {
		assert.False(t, IsValidHttpURL(v, false))
		assert.False(t, IsValidHttpURL(v, true))
	}
}
